# Gündemdeki verileri topla 
```shell
curl https://wikimedia.org/api/rest_v1/metrics/pageviews/top/tr.wikipedia/all-access/2025/08/01
```

# Gündemdeki verileri tek tek analiz et 

```shell
curl https://tr.wikipedia.org/w/api.php?action=query&prop=extracts&explaintext=true&titles=Mattia_Ahmet_Minguzzi_cinayeti&format=json
```

# Detials

**Eğer** :
gündemdeki olay kayıtlı ise es geçicek gündemdeki olaylar order tablosundan bakılıcak
eğer yok ise # Gündemdeki verileri tek tek analiz et  kısmına gidecek ve oradan onun hakkında exact'ı çekecek ve 
data tablesine kayıt edecek 


### cron 
her gün belirli bir vakitte dünün datasını çekecek ve işlem bittiğinde çekmediği günün datalarını çekecek ve onları kayıt edecek 
history table ve her işlem bittikten sonra 10 dk duracak 
