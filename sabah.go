package main

import (
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func init() {
	s := Sabah{"http://www.sabah.com.tr/rss/gundem.xml"}
	register(s)
}

type Sabah struct {
	feedUrl string
}

func (s Sabah) url() string {
	return s.feedUrl
}

func (s Sabah) contentHandler(root *html.Node) (string, bool) {

	newsContent, check := scrape.Find(root, scrape.ByClass("txt"))

	if check {
		newsParagraphs := scrape.FindAll(newsContent, scrape.ByTag(atom.P))

		newsText := ""
		for _, paragragh := range newsParagraphs {
			newsText = newsText + " " + scrape.Text(paragragh)
		}

		return newsText, true

	} else {
		return "", false
	}

}
