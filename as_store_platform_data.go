package scraper

type StorePlatformData struct {
	ProductDv            StorePlatformDataObject `json:"product-dv"`
	NativeSearchLockup   StorePlatformDataObject `json:"native-search-lockup"`
	Lockup               StorePlatformDataObject `json:"lockup"`
	EditorialItemProduct StorePlatformDataObject `json:"editorial-item-product"`
}

type StorePlatformDataObject struct {
	Results         map[string]Result `json:"results"` // Application ID, Story ID
	Version         int               `json:"version"`
	IsAuthenticated bool              `json:"isAuthenticated"`
	Meta            Meta              `json:"meta"`
}
