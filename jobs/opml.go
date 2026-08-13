package jobs

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

var doc OpmlDocument

func LoadSources(path string) ([]Source, error) {	
	var sources []Source

	body, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read opml file: %w", err)
	}

	if err := xml.Unmarshal(body, &doc); err != nil {
		return nil, fmt.Errorf("parse opml file: %w", err)
	}

	seen := make(map[string]struct{})

	CollectOutline(doc.Body.Outlines, seen, &sources)
	return sources, nil
}

func CollectOutline(outlines []OpmlOutline, seen map[string]struct{}, sources *[]Source) {
	for _, outline := range outlines {
		url := strings.TrimSpace(outline.XMLURL)
		if url != "" {
			if _, exists := seen[url]; !exists {
				name := strings.TrimSpace(outline.Title)
				if name == "" {
					name = strings.TrimSpace(outline.Text)
				}

				seen[url] = struct{}{}
				*sources = append(*sources, Source{
					Name:      name,
					URL:       url,
					Processor: ProcessFeed,
				})
			}
		}

		if len(outline.Outlines) > 0 {
			CollectOutline(outline.Outlines, seen, sources)
		}
	}
}
