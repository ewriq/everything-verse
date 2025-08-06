package jobs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	urls "net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/html"
)

func processReddit(body []byte) ([]Item, error) {
	var resp struct {
		Data struct {
			Children []struct {
				Data struct {
					Title    string `json:"title"`
					SelfText string `json:"selftext"`
					ID       string `json:"id"`
				} `json:"data"`
			} `json:"children"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	var items []Item
	for _, c := range resp.Data.Children {
		items = append(items, Item{
			Key:     "reddit_" + c.Data.ID,
			Title:   c.Data.Title,
			Content: c.Data.SelfText,
		})
	}
	return items, nil
}

func processHackerNews(body []byte) ([]Item, error) {
	var ids []int
	if err := json.Unmarshal(body, &ids); err != nil {
		return nil, err
	}
	if len(ids) > maxItemsToFetch {
		ids = ids[:maxItemsToFetch]
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var items []Item

	for _, id := range ids {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
			body, err := fetch(url)
			if err != nil {
				return
			}
			var data struct {
				Title string `json:"title"`
				Text  string `json:"text"`
				ID    int    `json:"id"`
			}
			if err := json.Unmarshal(body, &data); err != nil {
				return
			}
			mu.Lock()
			items = append(items, Item{
				Key:     fmt.Sprintf("hn_%d", data.ID),
				Title:   data.Title,
				Content: data.Text,
			})
			mu.Unlock()
		}(id)
	}
	wg.Wait()
	return items, nil
}

func processStackExchange(body []byte) ([]Item, error) {
	var resp struct {
		Items []struct {
			Title      string `json:"title"`
			Body       string `json:"body_markdown"`
			QuestionID int    `json:"question_id"`
		} `json:"items"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	var items []Item
	for _, q := range resp.Items {
		items = append(items, Item{
			Key:     fmt.Sprintf("se_%d", q.QuestionID),
			Title:   q.Title,
			Content: q.Body,
		})
	}
	return items, nil
}

func processMastodon(body []byte) ([]Item, error) {
	var toots []struct {
		Content string `json:"content"`
		ID      string `json:"id"`
		Account struct {
			Username string `json:"username"`
		} `json:"account"`
	}
	if err := json.Unmarshal(body, &toots); err != nil {
		return nil, err
	}
	var items []Item
	for _, toot := range toots {
		content := stripHTML(toot.Content)
		if content != "" {
			items = append(items, Item{
				Key:     "mastodon_" + toot.ID,
				Title:   "Mastodon post by @" + toot.Account.Username,
				Content: content,
			})
		}
	}
	return items, nil
}

func dataFromWikipedia() (bool, error) {
	var insertedAny atomic.Bool
	var wg sync.WaitGroup
	today := time.Now()
	limiter := make(chan struct{}, wikipediaConcurrency)

	for i := 1; i <= maxLookbackDays; i++ {
		date := today.AddDate(0, 0, -i)
		url := "https://wikimedia.org/api/rest_v1/metrics/pageviews/top/tr.wikipedia/all-access/" + date.Format("2006/01/02")

		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			body, err := fetch(url)
			if err != nil {
				return
			}
			var resp struct {
				Items []struct {
					Articles []struct {
						Article string `json:"article"`
					} `json:"articles"`
				} `json:"items"`
			}
			if err := json.Unmarshal(body, &resp); err != nil || len(resp.Items) == 0 {
				return
			}
			var articles = resp.Items[0].Articles

			var articleWG sync.WaitGroup
			for _, a := range articles {
				article := a
				articleWG.Add(1)
				limiter <- struct{}{}
				go func() {
					defer articleWG.Done()
					defer func() { <-limiter }()
					key := "wiki_" + urls.QueryEscape(article.Article)
					content, err := getWikipediaExtract(article.Article)
					if err != nil || content == "" {
						return
					}
					if ok, err := existsOrInsert(Item{Key: key, Title: article.Article, Content: content}); err == nil && ok {
						insertedAny.Store(true)
					}
				}()
			}
			articleWG.Wait()
		}(url)
	}
	wg.Wait()
	return insertedAny.Load(), nil
}

func getWikipediaExtract(title string) (string, error) {
	url := fmt.Sprintf("https://tr.wikipedia.org/w/api.php?action=query&prop=extracts&explaintext=true&titles=%s&format=json", urls.QueryEscape(title))
	body, err := fetch(url)
	if err != nil {
		return "", err
	}
	var resp struct {
		Query struct {
			Pages map[string]struct {
				Extract string `json:"extract"`
			} `json:"pages"`
		} `json:"query"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", err
	}
	for _, p := range resp.Query.Pages {
		if p.Extract != "" && !strings.HasSuffix(p.Extract, "may refer to:") {
			return p.Extract, nil
		}
	}
	return "", errors.New("makale özeti bulunamadı")
}

func processDevTo(body []byte) ([]Item, error) {
	var articles []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := json.Unmarshal(body, &articles); err != nil {
		return nil, err
	}
	var items []Item
	for _, a := range articles {
		items = append(items, Item{
			Key:     fmt.Sprintf("devto_%d", a.ID),
			Title:   a.Title,
			Content: a.Description,
		})
	}
	return items, nil
}

func processLobsters(body []byte) ([]Item, error) {
	var posts []struct {
		ShortID     string `json:"short_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := json.Unmarshal(body, &posts); err != nil {
		return nil, err
	}
	var items []Item
	for _, post := range posts {
		items = append(items, Item{
			Key:     "lobsters_" + post.ShortID,
			Title:   post.Title,
			Content: post.Description,
		})
	}
	return items, nil
}

func processProductHunt(body []byte) ([]Item, error) {
	var resp struct {
		Posts []struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Tagline string `json:"tagline"`
		} `json:"posts"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	var items []Item
	for _, p := range resp.Posts {
		items = append(items, Item{
			Key:     fmt.Sprintf("ph_%d", p.ID),
			Title:   p.Name,
			Content: p.Tagline,
		})
	}
	return items, nil
}

func processSlashdot(body []byte) ([]Item, error) {
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	var items []Item
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "item" {
			var title, desc, guid string
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode {
					switch c.Data {
					case "title":
						title = getTextContent(c)
					case "description":
						desc = getTextContent(c)
					case "guid":
						guid = getTextContent(c)
					}
				}
			}
			if guid != "" {
				items = append(items, Item{
					Key:     "slashdot_" + guid,
					Title:   title,
					Content: desc,
				})
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return items, nil
}

func getTextContent(n *html.Node) string {
	if n == nil {
		return ""
	}
	var b strings.Builder
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			b.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return b.String()
}

func processGitHubTrending(body []byte) ([]Item, error) {
	var resp struct {
		Items []struct {
			FullName    string `json:"full_name"`
			Description string `json:"description"`
			HTMLURL     string `json:"html_url"`
			Stargazers  int    `json:"stargazers_count"`
		} `json:"items"`
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	var items []Item
	for _, repo := range resp.Items {
		content := repo.Description
		if content == "" {
			content = fmt.Sprintf("GitHub repository with %d stars", repo.Stargazers)
		}
		items = append(items, Item{
			Key:     "github_" + strings.ReplaceAll(repo.FullName, "/", "_"),
			Title:   repo.FullName,
			Content: content,
		})
	}
	return items, nil
}

func processRSSFeed(body []byte) ([]Item, error) {
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	var items []Item
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "item" {
			var title, desc, guid string
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode {
					switch c.Data {
					case "title":
						title = getTextContent(c)
					case "description":
						desc = getTextContent(c)
					case "guid":
						guid = getTextContent(c)
					case "link":
						if guid == "" {
							guid = getTextContent(c)
						}
					}
				}
			}
			if guid != "" && title != "" {
				items = append(items, Item{
					Key:     "rss_" + guid,
					Title:   title,
					Content: desc,
				})
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return items, nil
}

func processSource(s Source) (bool, error) {
	body, err := fetch(s.URL)
	if err != nil {
		return false, fmt.Errorf("failed to fetch data: %w", err)
	}
	items, err := s.Processor(body)
	if err != nil {
		return false, fmt.Errorf("failed to process data: %w", err)
	}

	if len(items) == 0 {
		return false, nil
	}

	var insertedAny atomic.Bool
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10) 

	for _, itm := range items {
		if (itm.Content == "" && itm.Title == "") || itm.Key == "" {
			continue
		}
		if itm.Content == "" {
			itm.Content = itm.Title
		}

		wg.Add(1)
		go func(item Item) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			if ok, err := existsOrInsert(item); err == nil && ok {
				insertedAny.Store(true)
			}
		}(itm)
	}
	wg.Wait()

	return insertedAny.Load(), nil
}
