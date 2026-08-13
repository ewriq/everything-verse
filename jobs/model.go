package jobs

import (
	"net/http"
	"sync"
	"time"
)

var (
	maxWorker = 5
	maxRSSItemsPerFeed = 30
	httpTimeout        = 30 * time.Second
	SourcesPath = "sources.opml"
	httpClient = &http.Client{
		Timeout: httpTimeout,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
	}
	dbMutex sync.Mutex
	DataCollectionMu   sync.Mutex
	DataCollectionBusy bool
	InsertedCount int32
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

type OpmlDocument struct {
	Body OpmlBody `xml:"body"`
}

type OpmlBody struct {
	Outlines []OpmlOutline `xml:"outline"`
}

type OpmlOutline struct {
	Title    string        `xml:"title,attr"`
	Text     string        `xml:"text,attr"`
	Type     string        `xml:"type,attr"`
	XMLURL   string        `xml:"xmlUrl,attr"`
	Outlines []OpmlOutline `xml:"outline"`
}
type RssEnvelope struct {
	Channel struct {
		Items []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	GUID        string `xml:"guid"`
}

type AtomEnvelope struct {
	Entries []AtomEntry `xml:"entry"`
}

type AtomEntry struct {
	ID      string     `xml:"id"`
	Title   string     `xml:"title"`
	Summary string     `xml:"summary"`
	Content string     `xml:"content"`
	Links   []AtomLink `xml:"link"`
}

type AtomLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
}
