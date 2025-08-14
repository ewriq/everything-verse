package jobs

import (
	"everything-verse/database"
	"fmt"
	"io"
	"log"
	"net/http"
	urls "net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type WebCrawler struct {
	visited         map[string]bool
	discoveredSites map[string]bool
	queue           chan string
	wg              *sync.WaitGroup
	mutex           *sync.RWMutex
	client          *http.Client
	linkRegex       *regexp.Regexp
	maxWorkers      int
}

func NewWebCrawler(maxWorkers int) *WebCrawler {
	return &WebCrawler{
		visited:         make(map[string]bool),
		discoveredSites: make(map[string]bool),
		queue:           make(chan string, 1000),
		wg:              &sync.WaitGroup{},
		mutex:           &sync.RWMutex{},
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		linkRegex:  regexp.MustCompile(`https?://[^\s<>"']+`),
		maxWorkers: maxWorkers,
	}
}

func (wc *WebCrawler) isValidURL(rawURL string) bool {
	u, err := urls.Parse(rawURL)
	if err != nil {
		return false
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	path := strings.ToLower(u.Path)
	blockedExt := []string{".jpg", ".jpeg", ".png", ".gif", ".pdf", ".zip", ".mp4", ".mp3"}
	for _, ext := range blockedExt {
		if strings.HasSuffix(path, ext) {
			return false
		}
	}
	return true
}

func (wc *WebCrawler) scrapeText(pageURL string) (string, error) {
	resp, err := wc.client.Get(pageURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	var textParts []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			t := strings.TrimSpace(n.Data)
			if t != "" {
				textParts = append(textParts, t)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return strings.Join(textParts, " "), nil
}

func (wc *WebCrawler) fetchLinks(targetURL string) []string {
	wc.mutex.Lock()
	if wc.visited[targetURL] {
		wc.mutex.Unlock()
		return nil
	}
	wc.visited[targetURL] = true
	wc.mutex.Unlock()

	fmt.Println("INFO: Crawling", targetURL)

	resp, err := wc.client.Get(targetURL)
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	links := wc.linkRegex.FindAllString(string(body), -1)
	var valid []string
	for _, link := range links {
		link = strings.Trim(link, `"'<>()`)
		if wc.isValidURL(link) {
			valid = append(valid, link)
		}
	}
	return valid
}

func (wc *WebCrawler) worker() {
	defer wc.wg.Done()
	for target := range wc.queue {
		links := wc.fetchLinks(target)

		if text, err := wc.scrapeText(target); err == nil && text != "" {
			if !database.Exists(target) {
				database.Insert(target, text, target)
			}
		}
		
		for _, link := range links {
			wc.mutex.RLock()
			if !wc.visited[link] && len(wc.queue) < cap(wc.queue)-1 {
				select {
				case wc.queue <- link:
				default:
				}
			}
			wc.mutex.RUnlock()
		}
	}
}

func CrawlWeb() {
	crawler := NewWebCrawler(5)

	seedURLs := []string{
		// Teknoloji & Programlama
		"https://github.com/trending",
		"https://news.ycombinator.com",
		"https://reddit.com/r/programming",
		"https://stackoverflow.com/questions",
		"https://medium.com/topic/technology",
		"https://dev.to",
		"https://lobste.rs",
		"https://slashdot.org",
		"https://techcrunch.com",
		"https://arstechnica.com",
		"https://wired.com",
		"https://theverge.com",

		// Sosyal & Haber
		"https://reddit.com/r/all",
		"https://reddit.com/r/worldnews",
		"https://twitter.com/explore",
		"https://digg.com",
		"https://buzzfeed.com",
		"https://mashable.com",
		"https://engadget.com",

		// Eğitim & Bilim
		"https://wikipedia.org",
		"https://archive.org",
		"https://coursera.org",
		"https://khanacademy.org",
		"https://ted.com/talks",
		"https://nature.com",
		"https://sciencedaily.com",

		// İş & Finans
		"https://bloomberg.com",
		"https://reuters.com",
		"https://wsj.com",
		"https://forbes.com",
		"https://entrepreneur.com",
		"https://inc.com",

		// Medya & Eğlence
		"https://youtube.com/trending",
		"https://netflix.com",
		"https://imdb.com",
		"https://spotify.com",
		"https://twitch.tv",

		// Türkiye Siteleri
		"https://eksisozluk.com",
		"https://donanimhaber.com",
		"https://shiftdelete.net",
		"https://webtekno.com",
		"https://tamindir.com",
		"https://hurriyet.com.tr",
		"https://sozcu.com.tr",
		"https://sabah.com.tr",
	}

	for _, url := range seedURLs {
		if crawler.isValidURL(url) {
			crawler.queue <- url
		}
	}

	for i := 0; i < crawler.maxWorkers; i++ {
		crawler.wg.Add(1)
		go crawler.worker()
	}

	crawler.wg.Wait()
}
