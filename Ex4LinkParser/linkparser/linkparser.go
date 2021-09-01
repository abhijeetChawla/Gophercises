package linkparser

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

// this will work with valid html anchor tags and extract the href and the text inside of the anchor tag
// Invalid html like anchor tag inside and anchor tag will not work
func LinkParser(fileName string) (Links []Link, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Failed to load file", err)
	}
	linksNodes := getlinkNodes(file)
	Links = makeLinks(linksNodes)
	return Links, err
}

// this will not return an error because if we find no Nodes
// then we will return an empty slice
func getlinkNodes(r io.Reader) (nodes []html.Node) {
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			nodes = append(nodes, *n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nodes
}

// Make Link converts the anchor tags to Link struct
func makeLinks(links []html.Node) (ret []Link) {
	for _, link := range links {
		var l Link
		for _, attr := range link.Attr {
			if attr.Key == "href" {
				l.Href = attr.Val
				break
			}
		}
		l.Text = getText(&link)
		ret = append(ret, l)
	}
	return ret
}

// getText gets all the text nodes from the given node
func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += getText(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}
