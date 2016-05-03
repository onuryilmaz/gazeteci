package main

import "flag"

var gazeteci *Gazeteci

func initialize() {
	mongoPtr := flag.String("mongo", "mongodb://127.0.0.1:27017", "MongoDB URL")
	dbPtr := flag.String("db", "gazete", "MongoDB Database")
	collPtr := flag.String("coll", "collection", "MongoDB Collection")
	flag.Parse()

	var mongoConfig MongoConfig = MongoConfig{*mongoPtr, *dbPtr, *collPtr}
	gazeteci = New(mongoConfig)
}

func register(feed NewsFeed) {
	if gazeteci == nil {
		initialize()
	}
	gazeteci.Register(feed)
}

func main() {

	gazeteci.Start()
}
