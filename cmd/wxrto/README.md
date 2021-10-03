# wxrto

A tool for converting a WordPress E(x)tended RSS file into a static
site.

**Warning: this tool is incomplete and is under development.** It
implements _just enough_ of the Hugo-style content output that I
needed to port my Wordpress blog to Hugo.

```txt
$ ./wxrto --help
Usage of ./wxrto:
  -generator string
    	static site generator output format (default "hugo")
  -input string
    	the WordPress WXR file to convert (if not provided, stdin will be used)
  -outdir string
    	directory to save converted files and assets (default "output")
```
