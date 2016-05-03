package main

import (
	"fmt"
	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/jteeuwen/go-pkg-xmlx"
	"golang.org/x/net/html"
	"gopkg.in/mgo.v2"
	"net/http"
	"os"
	"sync"
	"time"
)

type Gazeteci struct {
	dbConfig   MongoConfig
	feeds      map[string]NewsFeed
	wg         sync.WaitGroup
	session    *mgo.Session
	collection *mgo.Collection
}

type MongoConfig struct {
	Host       string
	DB         string
	Collection string
}

type NewsFeed interface {
	url() string
	contentHandler(*html.Node) (string, bool)
}

type News struct {
	Title string `bson:"title"`
	URL   string `bson:"url"`
	Text  string `bson:"text"`
}

func New(cfg MongoConfig) *Gazeteci {
	obj := new(Gazeteci)
	obj.feeds = make(map[string]NewsFeed, 0)
	obj.dbConfig = cfg
	return obj
}

func (g Gazeteci) Register(feed NewsFeed) {
	g.feeds[feed.url()] = feed
}

func (g Gazeteci) Start() {

	var err error
	g.session, err = mgo.Dial(g.dbConfig.Host)
	if err != nil {
		panic(err)
	}
	g.collection = g.session.DB(g.dbConfig.DB).C(g.dbConfig.Collection)
	defer g.session.Close()

	for _, feed := range g.feeds {
		g.wg.Add(1)
		go g.pollFeed(feed.url(), 5, nil)
	}
	g.wg.Wait()
}

func (g Gazeteci) pollFeed(uri string, timeout int, cr xmlx.CharsetFunc) {

	defer g.wg.Done()
	feed := rss.New(timeout, true, g.chanHandler, g.itemHandler)

	if feed == nil {
		panic("Failed to start feed")
	}
	for {
		if err := feed.Fetch(uri, cr); err != nil {
			fmt.Fprintf(os.Stderr, "[e] %s: %s\n", uri, err)
			return
		}

		<-time.After(time.Duration(feed.SecondsTillUpdate() * 1e9))
	}
}

func (g Gazeteci) chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	//fmt.Printf("%d new channel(s) in %s\n", len(newchannels), feed.Url)
}

func (g Gazeteci) itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	for _, item := range newitems {
		resp, err := http.Get(item.Links[0].Href)
		if err != nil {
			fmt.Println("Error in getting news URL!")
		}
		root, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Println("Error in getting news page!")
		}
		newsContent, check := g.feeds[feed.Url].contentHandler(root)

		if check {
			news := News{item.Title, item.Links[0].Href, newsContent}
			if err := g.collection.Insert(news); err != nil {
				panic(err)
			}
			fmt.Println("Added: " + item.Title)

		}
	}
}
