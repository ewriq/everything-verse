* **Worker**

1. *Bounded Queue* — İş kuyruğunun kapasitesinin sınırlandırılması.
2. *Backpressure* — Yavaş bileşenlerin sistemi aşırı yüklemesinin engellenmesi.
3. *Job Lifecycle* — Collection joblarının durumlarının takip edilmesi.
4. *Dynamic Worker Count* — Worker sayısının ihtiyaca göre yönetilmesi.

* **Search**

1. *Pagination* — Büyük sonuç kümelerinin parçalara bölünmesi.
2. *Query Normalization* — Kullanıcı sorgularının FTS'e uygun hale getirilmesi.
3. *Prefix Search* — Kelime başlangıcı üzerinden arama desteği.
4. *Recency Ranking* — Güncel içeriklerin sıralamada dikkate alınması.
5. *Search Filters* — Kaynak ve içerik gibi alanlara göre filtreleme.
6. *Query Parsing* — Karmaşık arama ifadelerinin kontrollü işlenmesi.

* **FTS**

1. *External Content* — FTS içinde içerik tekrarının azaltılması.
2. *Index Synchronization* — FTS ile ana verinin kontrollü şekilde senkron tutulması.
3. *FTS Benchmark* — FTS yapılandırmalarının performans karşılaştırması.
4. *Storage Optimization* — FTS disk kullanımının azaltılması.
5. *Tokenizer Configuration* — İçeriğe uygun tokenizer yapılandırılması.
6. *Index Rebuild* — FTS indeksinin gerektiğinde yeniden oluşturulması.
7. *FTS Consistency* — Ana veri ile search indexinin tutarlılığının garanti edilmesi.

* **Domain**

1. *Source Entity* — RSS/Atom kaynaklarının DB üzerinde modellenmesi.
2. *Item Entity* — Feed içeriklerinin bağımsız entity olarak modellenmesi.
3. *Item Metadata* — İçerik metadata bilgilerinin saklanması.
4. *PublishedAt* — İçeriğin yayınlanma zamanının tutulması.
5. *GUID* — Feed itemlarının gerçek benzersiz kimliğinin korunması.
6. *Author* — İçerik yazar bilgilerinin tutulması.
7. *Categories* — İçerik kategorilerinin modellenmesi.
8. *Feed Status* — Kaynakların çalışma durumunun takip edilmesi.
9. *Canonical URL* — İçeriğin canonical URL bilgisinin belirlenmesi.

* **HTTP**

1. *Rate Limiting* — Kaynak sunuculara aşırı istek gönderilmesinin önlenmesi.
2. *Redirect Policy* — HTTP redirect davranışının kontrollü yönetilmesi.
3. *Content-Type Validation* — Beklenmeyen response içeriklerinin reddedilmesi.

* **API**

1. *Pagination* — Büyük veri sonuçlarının kontrollü döndürülmesi.
2. *Input Validation* — Kullanıcı girdilerinin kapsamlı şekilde doğrulanması.
3. *Readiness* — Uygulamanın istek kabul etmeye hazır olup olmadığının belirlenmesi.
4. *Liveness* — Uygulamanın çalışır durumda olup olmadığının belirlenmesi.
5. *API Versioning* — API değişikliklerinin geriye uyumlu yönetilmesi.

* **Observability**

1. *Structured Logging* — Logların makine tarafından işlenebilir tutulması.
2. *Metrics* — Sistem performansının sayısal olarak ölçülmesi.
3. *Error Counters* — Hata türlerinin ve sayıların takip edilmesi.
4. *Fetch Latency* — Feed isteklerinin sürelerinin ölçülmesi.
5. *Database Latency* — DB işlemlerinin sürelerinin ölçülmesi.
6. *Search Latency* — Arama performansının ölçülmesi.
7. *Worker Metrics* — Worker kullanımının takip edilmesi.
8. *Feed Metrics* — Kaynak bazlı başarı ve hata oranlarının takip edilmesi.

* **Testing**

1. *RSS Parser Tests* — RSS parser davranışlarının test edilmesi.
2. *Atom Parser Tests* — Atom parser davranışlarının test edilmesi.
3. *HTML Parser Tests* — HTML temizleme işlemlerinin test edilmesi.
4. *Deduplication Tests* — Duplicate içerik kontrolünün test edilmesi.
5. *Database Tests* — DB işlemlerinin test edilmesi.
6. *Transaction Tests* — Rollback davranışının test edilmesi.
7. *FTS Tests* — Search sonuçlarının test edilmesi.
8. *Integration Tests* — Bileşenlerin birlikte çalışmasının test edilmesi.
9. *Concurrency Tests* — Paralel işlemlerde veri bütünlüğünün test edilmesi.
10. *Benchmarks* — Kritik işlemlerin performansının ölçülmesi.

* **Optimization**

1. *Memory Usage* — Gereksiz bellek kullanımının azaltılması.
2. *CPU Usage* — Gereksiz CPU tüketiminin azaltılması.
3. *Database Throughput* — DB yazma/okuma kapasitesinin artırılması.
4. *Search Latency* — Arama süresinin azaltılması.
5. *Feed Fetch Throughput* — Feed toplama kapasitesinin artırılması.
6. *Network Usage* — Gereksiz network trafiğinin azaltılması.
7. *Allocation Reduction* — Gereksiz memory allocationlarının azaltılması.
8. *Large Dataset Testing* — Büyük veri altında sistem davranışının ölçülmesi.
9. *Load Testing* — Gerçekçi yük altında sistem sınırlarının belirlenmesi.

* **Security**

1. *Input Validation* — Zararlı veya geçersiz girdilerin kapsamlı şekilde filtrelenmesi.
2. *SSRF Protection* — Collector üzerinden iç ağlara erişimin engellenmesi.
3. *Rate Limiting* — API ve collector kullanımının sınırlandırılması.
4. *Resource Limits* — CPU, memory ve network kullanımının kontrol edilmesi.
5. *URL Validation* — İstek atılacak URLlerin güvenli şekilde doğrulanması.

* **Architecture**

1. *Repository Layer* — DB erişiminin domain mantığından ayrılması.
2. *Service Layer* — İş mantığının handler ve repositoryden ayrılması.
3. *Dependency Injection* — Bağımlılıkların kontrollü şekilde aktarılması.
4. *Configuration Management* — Runtime ayarlarının merkezi yönetilmesi.
5. *Interface Design* — Değiştirilebilir bileşenler için interface kullanılması.

* **Reliability**

1. *Failure Recovery* — Hatalardan sonra sistemin kontrollü şekilde devam etmesi.
2. *Duplicate Safety* — Duplicate kontrolünün DB seviyesinde garanti edilmesi.
3. *Transaction Recovery* — Yarım kalan DB işlemlerinin geri alınması.
4. *Graceful Degradation* — Kısmi hatalarda sistemin kullanılabilir kalması.

* **Deployment**

1. *Build Pipeline* — Build sürecinin tekrarlanabilir hale getirilmesi.
2. *Environment Separation* — Development ve production ayarlarının ayrılması.
3. *Health Checks* — Deployment sonrası uygulama durumunun kontrol edilmesi.
4. *Configuration Injection* — Ayarların deployment sırasında sağlanması.
5. *Database Backup* — SQLite verilerinin düzenli yedeklenmesi.
6. *Release Strategy* — Yeni sürümlerin kontrollü yayınlanması.
