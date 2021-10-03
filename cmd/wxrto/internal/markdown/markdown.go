package markdown

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type Node interface {
	Markdown() string
}

func markdownNodes(ns []Node) string {
	var s strings.Builder
	for _, n := range ns {
		s.WriteString(n.Markdown())
	}
	return s.String()
}

type TextPlain struct {
	Value string
}

func (t *TextPlain) Markdown() string {
	return t.Value
}

type TextStrong struct {
	Inner []Node
}

func (t *TextStrong) Markdown() string {
	return fmt.Sprintf("*%s*", markdownNodes(t.Inner))
}

type TextEmphasized struct {
	Inner []Node
}

func (t *TextEmphasized) Markdown() string {
	return fmt.Sprintf("_%s_", markdownNodes(t.Inner))
}

type TextMonospaced struct {
	Inner []Node
}

func (t *TextMonospaced) Markdown() string {
	return fmt.Sprintf("`%s`", markdownNodes(t.Inner))
}

type TextStrike struct {
	Inner []Node
}

func (t *TextStrike) Markdown() string {
	return fmt.Sprintf("~~%s~~", markdownNodes(t.Inner))
}

type Link struct {
	Ref   string
	Inner []Node
}

func (l *Link) Markdown() string {
	return fmt.Sprintf("[%s](%s)", markdownNodes(l.Inner), l.Ref)
}

type Paragraph struct {
	Nodes []Node
}

func (p *Paragraph) Markdown() string {
	return markdownNodes(p.Nodes)
}

type Header struct {
	Order int
	Value string
}

func (h *Header) Markdown() string {
	head := "######"[:h.Order]
	return fmt.Sprintf("%s %s", head, h.Value)
}

type Code struct {
	Nodes []Node
}

func (c *Code) Markdown() string {
	return strings.Join([]string{"```", markdownNodes(c.Nodes), "```"}, "\n")
}

type OrderedList struct {
	Nodes []Node
}

func (o *OrderedList) Markdown() string {
	var s []string

	for i, n := range o.Nodes {
		s = append(s, fmt.Sprintf("%d %s", i, n.Markdown()))
	}
	return strings.Join(s, "\n")
}

type UnorderedList struct {
	Nodes []Node
}

func (u *UnorderedList) Markdown() string {
	var s []string

	for _, n := range u.Nodes {
		s = append(s, fmt.Sprintf("* %s", n.Markdown()))
	}
	return strings.Join(s, "\n")
}

type Image struct {
	Ref string
	Alt string
}

func (i *Image) Markdown() string {
	return fmt.Sprintf("![%s](%s)", i.Alt, i.Ref)
}

type Document struct {
	Nodes []Node
}

func (d *Document) Markdown() string {
	var s []string

	for _, n := range d.Nodes {
		s = append(s, n.Markdown())
	}
	return strings.Join(s, "\n")
}

func New(n *html.Node) Document {
	if n == nil {
		return Document{}
	}

	return Document{markdownifyChildren(n)}
}

func markdownify(n *html.Node) Node {
	if n.Type == html.TextNode {
		return &TextPlain{n.Data}
	}

	switch n.Data {
	case "a":
		return handleTagA(n)
	case "strong":
		return handleTagStrong(n)
	case "b":
		return handleTagStrong(n)
	case "em":
		return handleTagEmphasized(n)
	case "i":
		return handleTagEmphasized(n)
	case "s":
		return handleTagStrike(n)
	case "pre":
		return handleTagPre(n)
	case "code":
		return handleTagMonospaced(n)
	case "h1":
		return handleTagH(n)
	case "h2":
		return handleTagH(n)
	case "h3":
		return handleTagH(n)
	case "h4":
		return handleTagH(n)
	case "h5":
		return handleTagH(n)
	case "h6":
		return handleTagH(n)
	case "img":
		return handleTagImg(n)
	case "ol":
		return handleTagOl(n)
	case "ul":
		return handleTagUl(n)
	default:
		return handleTagP(n)
	}

	// TODO: figure out best way to communicate error for unhandled tag.
}

func markdownifyChildren(n *html.Node) []Node {
	var children []Node
	if n == nil {
		return children
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode || c.Type == html.ElementNode {
			children = append(children, markdownify(c))
		}
	}
	return children
}

func handleTagP(n *html.Node) *Paragraph {
	return &Paragraph{markdownifyChildren(n)}
}

func handleTagStrong(n *html.Node) *TextStrong {
	return &TextStrong{markdownifyChildren(n)}
}

func handleTagEmphasized(n *html.Node) *TextEmphasized {
	return &TextEmphasized{markdownifyChildren(n)}
}

func handleTagMonospaced(n *html.Node) *TextMonospaced {
	return &TextMonospaced{markdownifyChildren(n)}
}

func handleTagStrike(n *html.Node) *TextStrike {
	return &TextStrike{markdownifyChildren(n)}
}

func handleTagPre(n *html.Node) *Code {
	for _, a := range n.Attr {
		if a.Key == "class" && a.Val == "wp-block-code" {
			code := n.FirstChild
			if code != nil && code.FirstChild != nil {
				value := []Node{&TextPlain{code.FirstChild.Data}}
				return &Code{value}
			}
		}
		if a.Key == "class" && a.Val == "wp-block-preformatted" {
			contents := n.FirstChild
			if contents != nil {
				value := []Node{&TextPlain{contents.Data}}
				return &Code{value}
			}
		}
	}

	// TODO
	return &Code{[]Node{}}
}

func handleTagH(n *html.Node) *Header {
	order, err := strconv.Atoi(n.Data[1:])
	if err != nil {
		order = 1
	}

	var value string
	if n.FirstChild != nil {
		value = n.FirstChild.Data
	}
	return &Header{Order: order, Value: value}
}

func handleTagImg(n *html.Node) *Image {
	var alt string
	var ref string

	for _, a := range n.Attr {
		if a.Key == "src" {
			ref = a.Val
		}
		if a.Key == "alt" {
			alt = a.Val
		}
	}

	return &Image{Alt: alt, Ref: ref}
}

func handleTagA(n *html.Node) *Link {
	var ref string

	for _, a := range n.Attr {
		if a.Key == "href" {
			ref = a.Val
		}
	}

	return &Link{Ref: ref, Inner: markdownifyChildren(n)}
}

func handleTagOl(n *html.Node) *OrderedList {
	return &OrderedList{markdownifyChildren(n)}
}

func handleTagUl(n *html.Node) *UnorderedList {
	return &UnorderedList{markdownifyChildren(n)}
}
