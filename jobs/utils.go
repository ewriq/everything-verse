package jobs

import (
	"bytes"
	"everything-verse/database"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func fetch(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json, application/xml, text/xml, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return body, nil
}

func existsOrInsert(item Item) (bool, error) {
	dbMutex.RLock()
	exists := database.Exists(item.Key)
	dbMutex.RUnlock()

	if exists {
		return false, nil
	}

	dbMutex.Lock()
	defer dbMutex.Unlock()

	if database.Exists(item.Key) {
		return false, nil
	}

	if err := database.Insert(item.Key, item.Content, item.Title); err != nil {
		return false, fmt.Errorf("database insert failed: %w", err)
	}

	fmt.Printf("ADDED: %s (Key: %s)\n", item.Title, item.Key)
	return true, nil
}

func stripHTML(input string) string {
	if input == "" {
		return ""
	}

	var b bytes.Buffer
	z := html.NewTokenizer(strings.NewReader(input))

	for {
		switch z.Next() {
		case html.ErrorToken:
			return strings.TrimSpace(b.String())
		case html.TextToken:
			b.Write(z.Text())
		}
	}
}
