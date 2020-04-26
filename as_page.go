package scraper

type Page struct {
	StorePlatformData StorePlatformData `json:"storePlatformData"`
	PageData          PageData          `json:"pageData"`
	Properties        Properties        `json:"properties"`
}
