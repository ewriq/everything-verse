package jobs


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
