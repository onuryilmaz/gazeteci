package main

import (
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"strings"
)

func init() {
	p := Posta{"http://www.posta.com.tr/xml/rss/rss_3_0.xml"}
	register(p)
}

type Posta struct {
	feedUrl string
}

func (p Posta) url() string {
	return p.feedUrl
}

func (p Posta) contentHandler(root *html.Node) (string, bool) {
	newsContentMatcherStar := func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "itemprop"), "articleBody")
	}

	newsContents := scrape.FindAll(root, newsContentMatcherStar)
	if len(newsContents) < 1 {
		return "", false
	}
	newsText := ""
	for _, node := range newsContents {
		newsText = newsText + scrape.Text(node)

	}
	return newsText, true
}
