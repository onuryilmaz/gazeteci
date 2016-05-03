package main

import (
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"strings"
)

func init() {
	s := Star{"http://www.star.com.tr/rss/rss.asp?cid=13"}
	register(s)
}

type Star struct {
	feedUrl string
}

func (s Star) url() string {
	return s.feedUrl
}

func (s Star) contentHandler(root *html.Node) (string, bool) {

	newsContentMatcherStar := func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "id"), "detaytext")
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
