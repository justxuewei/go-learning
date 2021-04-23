package ch07_interface

import (
	"golang.org/x/net/html"
	"testing"
)

func TestHtmlParser(t *testing.T) {
	node, _ := html.Parse(NewReader("<h1>Hello</h1>"))
	t.Log(node.Data)
}
