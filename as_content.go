package scraper

type Content struct {
	Type      string `json:"type"`
	KindIds   []int  `json:"kindIds"`
	ContentID string `json:"contentId"`
}
