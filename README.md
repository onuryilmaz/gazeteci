## Gazeteci
* Bu projede Türkçe gazetelere ait haberler toplanarak MongoDB'ye kaydedilmektedir.
* Şu anda desteklenen gazeteler:
 * [Cumhuriyet](http://www.cumhuriyet.com.tr/)
 * [Milliyet](http://www.milliyet.com.tr/)
 * [Posta](http://www.posta.com.tr)
 * [Sabah](http://www.sabah.com.tr)
 * [Star](http://www.star.com.tr/)

### Kullanım
* İşletim sisteminize uygun paketi bağlantıdan indirebilirsiniz:
```
gazeteci [--mongo=MONGO_DB_URL] [--db=MONGO_DB] [--coll=MONGO_COLLECTION]
```

### Katkıda Bulunun
 * Her gazete için RSS feed'lerinden haber linkleri toplanmakta ve websiteleri scrap edilerek haber metinlerinin tamamı çekilmektedir.
 * Yeni gazete eklemek içim `NewsFeed` interface'ini implement etmeniz ve yeni gazeteyi `gazeteci.Register` metodu ile kaydetmeniz gerekmektedir.