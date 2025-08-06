package jobs

import (
	"encoding/json"
	"everything-verse/database"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	baseWikiURL       = "https://tr.wikipedia.org/w/api.php"
	basePageViewURL   = "https://wikimedia.org/api/rest_v1/metrics/pageviews/top/tr.wikipedia/all-access/"
	extractMinLength  = 100000000000
)


func Data() {
	dateStr := time.Now().AddDate(0, 0, -1).Format("2006/01/02")
	fullURL := basePageViewURL + dateStr

	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error fetching pageviews:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var data topPagesResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("JSON unmarshal error:", err)
		return
	}

	if len(data.Items) == 0 {
		fmt.Println("No data found for date:", dateStr)
		return
	}

	for _, article := range data.Items[0].Articles {
		analyzeArticle(article.Article)
	}
}

func analyzeArticle(title string) {
	escapedTitle := url.QueryEscape(title)

	data, _ := database.Get(escapedTitle)
	if len(data) >= 10 {
		fmt.Println("Already exists:", title)
		return
	}

	fmt.Println("Analyzing:", title)

	apiURL := fmt.Sprintf(
		"%s?action=query&prop=extracts&explaintext=true&titles=%s&format=json",
		baseWikiURL, escapedTitle,
	)

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching wiki article:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading wiki response:", err)
		return
	}

	var wikiData wikiResponse
	if err := json.Unmarshal(body, &wikiData); err != nil {
		fmt.Println("Error parsing wiki JSON:", err)
		return
	}

	for _, page := range wikiData.Query.Pages {
		if len(page.Extract) >= extractMinLength {
			fmt.Println("Inserting:", title)
			if err := database.Insert(escapedTitle, page.Extract, title); err != nil {
				fmt.Println("Database insert error:", err)
			}
		}
	}
}
