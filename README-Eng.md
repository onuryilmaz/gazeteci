## Gazeteci (Newsstand in Turkish) [![Build Status](https://travis-ci.org/onuryilmaz/gazeteci.svg?branch=master)](https://travis-ci.org/onuryilmaz/gazeteci) [![English Readme](https://img.shields.io/badge/english-readme-blue.svg)](README-Eng.md)

* In this project, news from Turkish newspapers are gathered and dumped in to [MongoDB](https://www.mongodb.org/).
* Currently supported newspapers:
 * [Cumhuriyet](http://www.cumhuriyet.com.tr/)
 * [Milliyet](http://www.milliyet.com.tr/)
 * [Posta](http://www.posta.com.tr)
 * [Sabah](http://www.sabah.com.tr)
 * [Star](http://www.star.com.tr/)

### Usage
* Download the binary for your operating system from [releases](https://github.com/onuryilmaz/gazeteci/releases):
```
gazeteci [--mongo=MONGO_DB_URL] [--db=MONGO_DB] [--coll=MONGO_COLLECTION]
```
![](https://github.com/onuryilmaz/gazeteci/raw/master/screen-cast.gif)

### Contribute
 * For every newspaper, RSS feeds are checked for getting news URLs and then these URLs are used for scraping full texts of news.
 * In order to add a new newspaper you need to implement `NewsFeed` interface and register your new newspaper by `gazeteci.Register`.
