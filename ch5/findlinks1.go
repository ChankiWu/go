// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

// ch5.2
func visit(links []string, n *html.Node) []string {
	// visit函数遍历HTML的节点树，从每一个anchor元素的href属性获得link,将这些links存入字符串数组中，并返回这个字符串数组。
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	// main函数解析HTML标准输入，通过递归函数visit获得links（链接），并打印出这些links：
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}