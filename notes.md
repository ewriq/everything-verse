Tabii! History ve order tabloları ile sıkıştırma kısmını kaldırarak güncellenmiş döküman şöyle oldu:

---

# Gündemdeki Verileri Topla

```shell
curl https://wikimedia.org/api/rest_v1/metrics/pageviews/top/tr.wikipedia/all-access/2025/08/01
```

# Gündemdeki Verileri Tek Tek Analiz Et

```shell
curl "https://tr.wikipedia.org/w/api.php?action=query&prop=extracts&explaintext=true&titles=Mattia_Ahmet_Minguzzi_cinayeti&format=json"
```

# Detaylar

* Eğer veritabanında ilgili içerik yoksa,

  # Gündemdeki Verileri Tek Tek Analiz Et kısmına gidip o makale hakkında özet (extract) çekilecek ve

  veriler data tablosuna kaydedilecek.

---

### Cron İşleyişi

* Her gün belirli bir saatte, dünkü gündem verisi çekilecek.
* İşlem tamamlandıktan sonra, eksik olan günlerin verileri sırasıyla çekilip kaydedilecek.
* İşlem sonrası 10 dakika beklenip sonraki işlemler devam edecek.

---

# Veri Ekleme Mantığı

* Veritabanında, aranan içerik öncelikle kontrol edilecek.
* İçerik varsa tekrar eklenmeyecek, doğrudan mevcut veri döndürülecek.
* Yoksa içerik çekilip veritabanına eklenecek.

# tablo şeması 
extract : data
title : title 
normalized.from


extarct title query id 