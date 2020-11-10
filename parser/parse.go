package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents href and text
type Link struct {
	href string
	text string
}

// Parse reader to links
func Parse(r io.Reader) ([]Link, error) {
	document, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(document)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var link Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.href = attr.Val
		}
	}
	link.text = text(n)
	return link
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}
	var ret string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " "
	}
	return ret
}

func linkNodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	var nodes []*html.Node
	for i := node.FirstChild; i != nil; i = i.NextSibling {
		nodes = append(nodes, linkNodes(i)...)
	}
	return nodes
}

func dfs(node *html.Node, padding string) {
	msg := node.Data
	if node.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
