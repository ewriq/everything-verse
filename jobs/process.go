package jobs

import (
	"encoding/xml"
	"errors"
	"fmt"
	"sort"
	"strings"
)

func ProcessFeed(body []byte) ([]Item, error) {
	items, err := ParseItems(body)
	if err == nil && len(items) > 0 {
		return items, nil
	}

	items, atomErr := ParseAtomItems(body)

	if atomErr == nil && len(items) > 0 {
		return items, nil
	}
	if err != nil {
		return nil, err
	}
	return nil, atomErr
}

func ParseItems(body []byte) ([]Item, error) {
	var feed RssEnvelope
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, err
	}

	if len(feed.Channel.Items) == 0 {
		return nil, errors.New("rss feed has no items")
	}

	items := make([]Item, 0, min(maxRSSItemsPerFeed, len(feed.Channel.Items)))
	for _, entry := range feed.Channel.Items {
		title := strings.TrimSpace(StripHTML(entry.Title))
		content := strings.TrimSpace(StripHTML(entry.Description))
		keySource := FirstEmpty(strings.TrimSpace(entry.GUID), strings.TrimSpace(entry.Link), title)
		if keySource == "" {
			continue
		}
		if content == "" {
			content = title
		}

		items = append(items, Item{
			Key:     BuildFeedKey("rss", keySource),
			Title:   title,
			Content: content,
		})
		if len(items) >= maxRSSItemsPerFeed {
			break
		}
	}

	if len(items) == 0 {
		return nil, errors.New("rss feed produced no valid items")
	}
	return items, nil
}

func ParseAtomItems(body []byte) ([]Item, error) {
	var feed AtomEnvelope
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, err
	}

	if len(feed.Entries) == 0 {
		return nil, errors.New("atom feed entries")
	}

	items := make([]Item, 0, min(maxRSSItemsPerFeed, len(feed.Entries)))
	for _, entry := range feed.Entries {
		title := strings.TrimSpace(StripHTML(entry.Title))
		content := strings.TrimSpace(StripHTML(FirstEmpty(entry.Content, entry.Summary)))
		link := strings.TrimSpace(AtomEntryLink(entry.Links))
		keySource := FirstEmpty(strings.TrimSpace(entry.ID), link, title)
		if keySource == "" {
			continue
		}

		if content == "" {
			content = title
		}

		items = append(items, Item{
			Key:     BuildFeedKey("atom", keySource),
			Title:   title,
			Content: content,
		})

		if len(items) >= maxRSSItemsPerFeed {
			break
		}
	}

	if len(items) == 0 {
		return nil, errors.New(" no valid items")
	}
	return items, nil
}

func AtomEntryLink(links []AtomLink) string {
	if len(links) == 0 {
		return ""
	}
	if len(links) == 1 {
		return links[0].Href
	}

	sorted := make([]AtomLink, len(links))
	copy(sorted, links)
	sort.SliceStable(sorted, func(i, j int) bool {
		if sorted[i].Rel == "alternate" {
			return true
		}
		if sorted[j].Rel == "alternate" {
			return false
		}
		return i < j
	})
	return sorted[0].Href
}

func ProcessSource(s Source) (bool, error) {
	body, err := Fetch(s.URL)
	if err != nil {
		return false, fmt.Errorf("failed to Fetch: %w", err)
	}

	items, err := s.Processor(body)
	if err != nil {
		return false, fmt.Errorf("failed to process: %w", err)
	}

	inserted := false

	for _, item := range items {
		if len(item.Title) == 0 || len(item.Key) == 0 || len(item.Content) < 300 {
			continue
		}

		ok, err := Exists(item, s.URL)
		if err != nil {
			return false, fmt.Errorf("insert %q: %w", item.Key, err)
		}

		if ok {
			inserted = true
		}
	}

	return inserted, nil
}