package markdown

import (
	"golang.org/x/net/html"
)

type NodeKind int

const (
	NodeUnknown NodeKind = iota
	NodeHTMLInternal
	NodeParagraph
	NodePlainText
	NodeStrongText
	NodeEmphasizedText
	NodeStrikeText
	NodeMonoText
	NodeLink
	NodeHeader
	NodeImage
	NodePreformatted
	NodeUnorderedList
	NodeOrderedList
	NodeListItem
)

const (
	NodeAttrHref    = "href"
	NodeHeaderOrder = "head-order"
	NodeImageSrc    = "img-src"
	NodeImageAlt    = "img-alt"
)

type Node struct {
	Kind        NodeKind
	Attrs       map[string]string
	Data        string
	FirstChild  *Node
	NextSibling *Node
	PrevSibling *Node
}

func FromHTMLNode(n *html.Node) *Node {
	if n == nil {
		return nil
	}

	root := &Node{
		Attrs: make(map[string]string),
	}

	switch n.Type {
	case html.TextNode:
		root.Kind = NodePlainText
		root.Data = n.Data
	case html.ElementNode:
		switch n.Data {
		case "a":
			root.Kind = NodeLink
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					root.Attrs[NodeAttrHref] = attr.Val
				}
			}
		case "b":
			root.Kind = NodeStrongText
		case "strong":
			root.Kind = NodeStrongText
		case "i":
			root.Kind = NodeEmphasizedText
		case "em":
			root.Kind = NodeEmphasizedText
		case "s":
			root.Kind = NodeStrikeText
		case "pre":
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "wp-block-code" {
					return &Node{
						Kind: NodePreformatted,
						FirstChild: &Node{
							Kind: NodePlainText,
							Data: n.FirstChild.FirstChild.Data,
						},
					}
				}
			}
		case "code":
			root.Kind = NodeMonoText
		case "h1":
			root.Kind = NodeHeader
			root.Attrs[NodeHeaderOrder] = "1"
		case "h2":
			root.Kind = NodeHeader
			root.Attrs[NodeHeaderOrder] = "2"
		case "h3":
			root.Kind = NodeHeader
			root.Attrs[NodeHeaderOrder] = "3"
		case "h4":
			root.Kind = NodeHeader
			root.Attrs[NodeHeaderOrder] = "4"
		case "h5":
			root.Kind = NodeHeader
			root.Attrs[NodeHeaderOrder] = "5"
		case "h6":
			root.Kind = NodeHeader
			root.Attrs[NodeHeaderOrder] = "6"
		case "ol":
			root.Kind = NodeOrderedList
		case "ul":
			root.Kind = NodeUnorderedList
		case "li":
			root.Kind = NodeListItem
		case "img":
			root.Kind = NodeImage
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					root.Attrs[NodeImageSrc] = attr.Val
				}
			}
		}
	default:
		root.Kind = NodeHTMLInternal
	}

	var last *Node

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		r := FromHTMLNode(c)
		if r == nil {
			continue
		}

		if root.FirstChild == nil {
			root.FirstChild = r
			last = r
		} else {
			r.PrevSibling = last
			last.NextSibling = r
			last = r
		}
	}

	return root
}
