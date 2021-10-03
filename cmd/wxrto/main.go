// Command wxrto converts a WordPress E(x)tended RSS file into a static
// site.
package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"text/template"
	"time"

	"golang.org/x/net/html"

	"github.com/connorkuehl/wxr"
	"github.com/connorkuehl/wxr/cmd/wxrto/internal/markdown"
)

var (
	// numTokens limits the number of concurrent conversions.
	numTokens = 16

	// tokens is used as a lock that each conversion goroutine must
	// acquire before running (it must also release it when it's done).
	// See acquire() and release().
	tokens = make(chan struct{}, numTokens)

	// acquired tracks running converter goroutines. It is updated
	// as part of acquire() and release().
	acquired sync.WaitGroup

	// inputFile is the path to the WordPress E(x)tended RSS file that
	// will be decomposed into Markdown files for a static site generator.
	inputFile *string

	// generator is the name of a static site generator to convert to.
	generator *string

	// outputDir is the root directory to where markdown files and assets
	// will be saved to.
	outputDir *string

	// contentDir is the directory relative to outputDir where posts, pages,
	// and assets will be saved to.
	contentDir = func(out string) string { return fmt.Sprintf("%s/content", out) }

	// postsDir is the directory relative to the contentDir where blog post Markdown
	// files will be written to.
	postsDir = func(out string) string { return fmt.Sprintf("%s/posts", contentDir(out)) }

	// pagesDir is the directory relative to the contentDir where standalone static
	// Markdown pages will be written to.
	pagesDir = contentDir
)

func init() {
	inputFile = flag.String("input", "", "the WordPress WXR file to convert (if not provided, stdin will be used)")
	generator = flag.String("generator", "hugo", "static site generator output format")
	outputDir = flag.String("outdir", "output", "directory to save converted files and assets")
	flag.Parse()

	// Populate the pool of tokens outright to avoid deadlock
	for i := 0; i < numTokens; i++ {
		tokens <- struct{}{}
	}
}

// acquire must be called before starting a converter goroutine.
func acquire() {
	<-tokens
	acquired.Add(1)
}

// release must be called by a goroutine before it exits.
func release() {
	acquired.Done()
	tokens <- struct{}{}
}

func main() {
	var in io.Reader

	// Input filepath wasn't provided, fall back to stdin
	if *inputFile == "" {
		in = os.Stdin
	} else {
		f, err := os.Open(*inputFile)
		if err != nil {
			log.Fatal(err)
		}
		in = f
	}

	// xml.Unmarshal wants to operate on a []byte
	buf := bytes.Buffer{}
	_, err := io.Copy(&buf, in)
	if err != nil {
		log.Fatal(err)
	}

	// Deserialize it into XML
	var rss wxr.RSS
	err = xml.Unmarshal(buf.Bytes(), &rss)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range rss.Channel.Items {
		// The processItem goroutine will release the lock before returning
		acquire()
		go func(r *wxr.RSS, i wxr.Item) {
			processItem(r, i)
		}(&rss, item)
	}

	// Wait for any dispatched goroutines to finish up before exiting
	acquired.Wait()
}

// processItem converts a WordPress blog post or static page into a Markdown
// file that is compatible with the selected generator.
func processItem(rss *wxr.RSS, item wxr.Item) {
	// Make sure we release this
	defer release()

	postType := stripCharData(item.PostType)
	if postType != "post" && postType != "page" {
		return
	}

	status := stripCharData(item.Status)
	if status == "trash" {
		log.Printf("%q marked as trash, ignoring", item.Title)
		return
	}

	var path string
	if postType == "post" {
		path = postsDir(*outputDir)
	} else {
		path = pagesDir(*outputDir)
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		log.Printf("failed to make output directory %q: %v", path, err)
		return
	}

	// Parse the date so we can prefix the post with YYYY-MM-DD.
	//
	// TODO: make this more configurable, not everyone wants the date
	// in the filename.
	posted, err := time.Parse("2006-01-02 15:04:05", stripCharData(item.PostDate))
	if err != nil {
		log.Printf("failed to parse post date: %v", err)
		return
	}

	// TODO: probably write a function to kebab-case the title, as I'm
	// not sure WP guarantees this will be the way I think it is
	name := stripCharData(item.PostName)

	var filename string
	if postType == "post" {
		filename = fmt.Sprintf("%s/%s-%s.md", path, posted.Format("2006-01-02"), name)
	} else {
		filename = fmt.Sprintf("%s/%s.md", path, name)
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Printf("unable to create file %q: %v", filename, err)
		return
	}

	htmlDoc, err := html.Parse(strings.NewReader(item.Content.Data))
	if err != nil {
		log.Printf("parsing %q failed: %v", item.Title, err)
		return
	}

	mdDoc := markdown.New(htmlDoc)

	// TODO: Either visit the HTML tree or the markdown tree and download
	// assets to outputDir/contentDir/static. The links will probably have
	// to be fixed up.

	// Set up the front matter template
	tmpl, ok := Posts[*generator]
	if !ok {
		log.Fatal(fmt.Errorf("generator %q not installed", *generator))
		return
	}

	// TODO: This will have to be configurable to support more than just
	// Hugo
	var frontmatter struct {
		Title string
		Date  string
		Draft bool
	}

	frontmatter.Title = `"` + item.Title + `"`
	frontmatter.Date = posted.Format("2006-01-02 15:04:05")
	if item.Status == "publish" {
		frontmatter.Draft = false
	} else {
		frontmatter.Draft = true
	}

	t := template.Must(template.New(*generator + "-post").Parse(tmpl))

	// Write the frontmatter, then the Markdown
	t.Execute(file, frontmatter)
	io.Copy(file, strings.NewReader(mdDoc.Markdown()))
	log.Printf("%q => %q", item.Title, filename)
}

// TODO: Fixup the wxr/xml package to parse out the chardata
func stripCharData(s string) string {
	return strings.TrimSuffix(strings.TrimPrefix(s, "<![CDATA["), "]]>")
}
