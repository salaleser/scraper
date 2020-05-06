package scraper

type Link struct {
	URL       string `json:"url"`
	KindID    string `json:"kindId"`
	Type      string `json:"type"`
	KindIds   []int  `json:"kindIds"`
	ContentID string `json:"contentId"`
}
