# Everything-Verse 🌐

A comprehensive content aggregation system that collects and processes data from **200+ tech sources** including news, blogs, developer resources, and social platforms. Built with Go and Fiber framework.

## 🚀 Features

### 📊 **Massive Data Collection**
- **200+ Sources**: Tech news, developer blogs, AI/ML, security, cloud platforms, and more
- **Real-time Processing**: Concurrent data fetching and processing
- **Smart Deduplication**: Prevents duplicate content storage
- **Multi-format Support**: RSS feeds, JSON APIs, and custom processors

### 🏗️ **Architecture**
- **High Performance**: Concurrent processing with semaphore limiting
- **Scalable Design**: Easy to add new sources and processors
- **Database Optimization**: SQLite with efficient indexing
- **RESTful API**: Clean API endpoints for data access

### 📈 **Source Categories**
- **Social & Community**: HackerNews, Reddit, Mastodon, ProductHunt
- **Development & Tech**: StackExchange, Dev.to, GitHub Trending
- **Programming Languages**: Go, Python, Rust, JavaScript, Java, C#, PHP, Ruby, Swift, Kotlin
- **Frontend & Web**: React, Vue, Angular, CSS, JavaScript frameworks
- **Mobile Development**: Android, iOS, React Native, Flutter
- **Game Development**: Unity, Unreal Engine, Godot
- **AI & Machine Learning**: OpenAI, Google AI, DeepMind, Microsoft AI
- **Security & Privacy**: Schneier, Krebs, Security Weekly
- **Cloud & DevOps**: AWS, Google Cloud, Azure, Docker, Kubernetes
- **Database & Data**: MongoDB, PostgreSQL, Redis, Elasticsearch
- **Hardware & IoT**: Arduino, Raspberry Pi, ESP32
- **Linux & Open Source**: Linux Foundation, Red Hat, Ubuntu
- **Blockchain & Crypto**: Ethereum, Bitcoin, CoinDesk
- **Tech Companies**: Netflix, Uber, Airbnb, Spotify, GitHub, Stripe

## 🛠️ Technology Stack

- **Backend**: Go 1.24.5
- **Web Framework**: Fiber v2.52.9
- **Database**: SQLite with GORM
- **HTTP Client**: Custom optimized client with connection pooling
- **Concurrency**: Goroutines with WaitGroups and semaphores
- **Data Processing**: RSS parsing, JSON processing, HTML parsing

## 📦 Installation

### Prerequisites
- Go 1.24.5 or higher
- Git

### Quick Start

1. **Clone the repository**
```bash
git clone https://github.com/your-username/everything-verse.git
cd everything-verse
```

2. **Install dependencies**
```bash
go mod download
```

3. **Run the application**
```bash
go run main.go
```

The server will start on `http://localhost:3000`

## 🔧 Configuration

### Environment Variables
```bash
# Server Configuration
PORT=3000

# Database Configuration
DB_PATH=./database/data.db

# Concurrency Settings
MAX_CONCURRENT_FETCH=20
MAX_CONCURRENT_DB=10
MAX_ITEMS_TO_FETCH=100
```

### Constants (in `jobs/model.go`)
```go
const (
    maxLookbackDays      = 7
    maxItemsToFetch      = 100
    maxConcurrentDB      = 10
    maxConcurrentFetch   = 20
    httpTimeout          = 30 * time.Second
    dbTimeout            = 5 * time.Second
)
```

## 📡 API Endpoints

### Data Retrieval
```bash
# Get all data
GET /api/data

# Search data by keyword
GET /api/search?q=keyword

# Get data by source
GET /api/source/{source_name}

# Get recent data
GET /api/recent?limit=50
```

### Health Check
```bash
GET /health
```

## 🔄 Data Collection Process

### Automatic Collection
- **Frequency**: Daily at scheduled intervals
- **Workers**: 20 concurrent workers
- **Processing**: All 200+ sources processed in parallel
- **Storage**: SQLite database with optimized queries

### Manual Collection
```bash
# Trigger data collection manually
curl -X POST http://localhost:3000/api/collect
```

## 📁 Project Structure

```
everything-verse/
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
├── jobs/                   # Data collection jobs
│   ├── data.go            # Source definitions (200+ sources)
│   ├── process.go         # Data processors
│   ├── model.go           # Data models and constants
│   ├── utils.go           # Utility functions
│   ├── worker.go          # Worker implementation
│   └── index.go           # Cron job scheduler
├── database/              # Database layer
│   ├── connect.go         # Database connection
│   ├── data.go            # Data operations
│   └── model.go           # Database models
├── internal/              # Internal application logic
│   ├── handler/           # HTTP handlers
│   ├── middleware/        # HTTP middleware
│   └── routes/            # API routes
└── database/              # SQLite database files
    └── data.db           # Main database file
```

## 🚀 Performance Features

### Concurrent Processing
- **Parallel Source Processing**: All 200+ sources processed simultaneously
- **Semaphore Limiting**: Database operations limited to prevent overload
- **Connection Pooling**: Optimized HTTP client with connection reuse
- **Read-Write Mutex**: Efficient database locking strategy

### Memory Optimization
- **Streaming Processing**: Large responses processed in chunks
- **Garbage Collection**: Efficient memory management
- **Connection Limits**: Controlled concurrent connections

## 🔍 Data Sources

### Current Sources (200+)
- **Tech News**: TechCrunch, The Verge, Wired, Ars Technica
- **Developer Blogs**: Stack Overflow, Dev.to, GitHub Engineering
- **Programming Languages**: Official blogs for Go, Python, Rust, Java, C#
- **AI/ML**: OpenAI, Google AI, DeepMind, Microsoft AI
- **Security**: Schneier, Krebs, Security Weekly
- **Cloud Platforms**: AWS, Google Cloud, Azure, DigitalOcean
- **Mobile Development**: Android, iOS, React Native, Flutter
- **Game Development**: Unity, Unreal Engine, Godot
- **Blockchain**: Ethereum, Bitcoin, CoinDesk
- **Hardware**: Arduino, Raspberry Pi, ESP32

## 🛡️ Security Features

- **Rate Limiting**: Prevents API abuse
- **CORS Configuration**: Cross-origin resource sharing
- **Input Validation**: Sanitized inputs
- **Error Handling**: Graceful error responses
- **Logging**: Comprehensive request logging

## 📊 Monitoring & Logging

### Console Output
```
INFO: Starting data collection workers...
INFO: Worker 1 starting data collection
INFO: Fetching data from HackerNews...
INFO: New data added from HackerNews
INFO: Successfully added data from 45 sources
```

### Performance Metrics
- **Processing Time**: Real-time duration tracking
- **Success Rate**: Source processing statistics
- **Error Tracking**: Failed requests logging
- **Database Operations**: Insert/update statistics

## 🔧 Development

### Adding New Sources
1. Add source to `jobs/data.go`
2. Implement processor in `jobs/process.go` (if needed)
3. Test with `go run main.go`

### Example Source Addition
```go
{"New Source", "https://api.example.com/feed", processRSSFeed},
```

### Running Tests
```bash
go test ./...
```

## 📈 Scaling Considerations

### Horizontal Scaling
- **Load Balancing**: Multiple instances behind a load balancer
- **Database Sharding**: Distribute data across multiple databases
- **Caching**: Redis for frequently accessed data

### Vertical Scaling
- **Resource Limits**: Adjust concurrent worker counts
- **Memory Optimization**: Tune garbage collection
- **Database Optimization**: Index optimization and query tuning

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **Fiber Framework**: High-performance web framework
- **GORM**: Database ORM for Go
- **SQLite**: Lightweight database engine
- **All Data Sources**: For providing valuable content


---

**Everything-Verse** - Aggregating the world's tech knowledge, one source at a time. 🌐✨ 
