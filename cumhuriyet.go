package main

import (
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"strings"
)

func init() {
	c := Cumhuriyet{"http://www.cumhuriyet.com.tr/rss/4.xml"}
	register(c)
}

type Cumhuriyet struct {
	feedUrl string
}

func (c Cumhuriyet) url() string {
	return c.feedUrl
}

func (c Cumhuriyet) contentHandler(root *html.Node) (string, bool) {

	newsContentMatcherCumhuriyet := func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "id"), "news-body")
	}

	newsContent, check := scrape.Find(root, newsContentMatcherCumhuriyet)
	if check {
		newsText := scrape.Text(newsContent)
		return newsText, true

	} else {
		return "", false
	}

}
