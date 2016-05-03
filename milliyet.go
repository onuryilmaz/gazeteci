package main

import (
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"strings"
)

func init() {
	m := Milliyet{"http://www.milliyet.com.tr/rss/rssNew/dunyaRss.xml"}
	register(m)
}

type Milliyet struct {
	feedUrl string
}

func (m Milliyet) url() string {
	return m.feedUrl
}

func (m Milliyet) contentHandler(root *html.Node) (string, bool) {

	newsContentMatcherMilliyet := func(n *html.Node) bool {
		return strings.Contains(scrape.Attr(n, "itemprop"), "articleBody")
	}

	newsContents := scrape.FindAll(root, newsContentMatcherMilliyet)
	if len(newsContents) < 1 {
		return "", false
	}
	newsText := ""
	for _, node := range newsContents {
		newsText = newsText + scrape.Text(node)

	}
	return newsText, true

}
