package scraper

type ReportConcernReasons []struct {
	ReasonID      string `json:"reasonId"`
	Name          string `json:"name"`
	UpperCaseName string `json:"upperCaseName"`
}
