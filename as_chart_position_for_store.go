package scraper

type ChartPositionForStore struct {
	AppStore ChartPosition `json:"appStore"`
}

type ChartPosition struct {
	Position  int    `json:"position"`
	GenreName string `json:"genreName"`
	ChartURL  string `json:"chartUrl"`
}
