package jobs

import (
	"everything-verse/database"
	"fmt"
	"io"
	"net/http"
	urls "net/url"
	"strings"
	//"time"

	"golang.org/x/net/html"
)



func DeepSearch() {
	titles, err := database.GetTitles()
	if err != nil {
		fmt.Println("ERROR: Could not get titles:", err)
		return
	}

	totalInserted := 0

	for _, t := range titles {
		src := SearchSource{Title: t}
		fmt.Printf("INFO: Searching DuckDuckGo for '%s'...\n", src.Title)
		if inserted := processSearchSource(src); inserted {
			fmt.Printf("INFO: New data added for '%s'\n", src.Title)
			totalInserted++
		}
		//time.Sleep(2 * time.Second)
	}

	if totalInserted == 0 {
		fmt.Println("INFO: No new data added from any search")
	} else {
		fmt.Printf("INFO: Successfully added data for %d titles\n", totalInserted)
	}
}

func processSearchSource(s SearchSource) bool {
	links, err := duckduckGoSearch(s.Title)
	if err != nil {
		return false
	}

	insertedAny := false
	for i, link := range links {
		if i >= 3 {
			break
		}
		text, err := scrapeText(link)
		if err != nil || text == "" {
			continue
		}
		if err := database.Insert(s.Title, text, s.Title); err == nil {
			insertedAny = true
		}
		//time.Sleep(1 * time.Second)
	}
	return insertedAny
}

func duckduckGoSearch(query string) ([]string, error) {
	searchURL := fmt.Sprintf("https://duckduckgo.com/html/?q=%s", urls.QueryEscape(query))
	resp, err := http.Get(searchURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	htmlStr := string(body)

	var results []string
	parts := strings.Split(htmlStr, `<a rel="nofollow" class="result__a" href="`)
	for i := 1; i < len(parts); i++ {
		end := strings.Index(parts[i], `"`)
		if end > 0 {
			results = append(results, parts[i][:end])
		}
	}
	return results, nil
}

func scrapeText(pageURL string) (string, error) {
	resp, err := http.Get(pageURL)
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
