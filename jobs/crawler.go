package jobs

import (
	"everything-verse/database"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	queue   = make(chan string, 1000)
	visited = make(map[string]bool)
	mu      sync.Mutex
)

func init() {
	startURL := "https://data.gov" 
	queue <- startURL
}

func Crawling() {
	for {
		select {
		case link := <-queue:
			if !visit(link) {
				continue
			}
			processLink(link)
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func visit(link string) bool {
	mu.Lock()
	defer mu.Unlock()

	if visited[link] {
		return false
	}
	visited[link] = true
	return true
}

func processLink(link string) {
	fmt.Println("Crawling:", link)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(link)
	if err != nil {
		log.Println("HTTP error:", err)
		return
	}
	defer resp.Body.Close()

	if !strings.HasPrefix(resp.Header.Get("Content-Type"), "text/html") {
		return
	}


	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read error:", err)
		return
	}
	htmlContent := string(bodyBytes)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Println("Parse error:", err)
		return
	}

	if !database.Exists(link) {
		err := database.Insert(link, htmlContent, link)
		if err != nil {
			log.Println("DB insert error:", err)
		}
	} 

	base, _ := url.Parse(link)
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		u, err := url.Parse(strings.TrimSpace(href))
		if err != nil {
			return
		}
		abs := base.ResolveReference(u).String()
		if strings.HasPrefix(abs, "http") {
			enqueue(abs)
		}
	})
}

func enqueue(link string) {
	mu.Lock()
	defer mu.Unlock()
	if !visited[link] {
		queue <- link
	}
}
