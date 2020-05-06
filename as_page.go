package scraper

type Page struct {
	StorePlatformData map[string]StorePlatformData `json:"storePlatformData"`
	PageData          PageData                     `json:"pageData"`
	Properties        Properties                   `json:"properties"`
}
