package main

// hugoPostTmpl is the YAML frontmatter for the Hugo static site generator.
var hugoPostTmpl = `---
title: {{printf "%s" .Title}}
date: {{.Date}}
draft: {{.Draft}}
---

`

var Posts = map[string]string{
	"hugo": hugoPostTmpl,
}
