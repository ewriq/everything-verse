package jobs

import (
	"net/http"
	"sync"
	"time"
)

type topPagesResponse struct {
	Items []struct {
		Articles []struct {
			Article string `json:"article"`
			Views   int    `json:"views"`
			Rank    int    `json:"rank"`
		} `json:"articles"`
	} `json:"items"`
}

type wikiResponse struct {
	Query struct {
		Pages map[string]struct {
			Title   string `json:"title"`
			Extract string `json:"extract"`
		} `json:"pages"`
	} `json:"query"`
}

const (
	maxLookbackDays      = 7
	maxItemsToFetch      = 100
	maxConcurrentDB      = 10
	maxConcurrentFetch   = 20
	userAgent            = "everything-verse-bot/1.0 (https://github.com/your-repo)"
	wikipediaConcurrency = 10
	httpTimeout          = 30 * time.Second
	dbTimeout            = 5 * time.Second
)

var (
	httpClient = &http.Client{
		Timeout: httpTimeout,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
	}
	dbMutex sync.RWMutex
)

type Item struct {
	Key     string `json:"key"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Source struct {
	Name      string
	URL       string
	Processor func(body []byte) ([]Item, error)
}
