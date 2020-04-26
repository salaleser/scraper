package scraper

type StorePlatformData struct {
	ProductDv            StorePlatformDataType `json:"product-dv"`
	NativeSearchLockup   StorePlatformDataType `json:"native-search-lockup"`
	Lockup               StorePlatformDataType `json:"lockup"`
	EditorialItemProduct StorePlatformDataType `json:"editorial-item-product"`
}

type StorePlatformDataType struct {
	Results         interface{} `json:"results"` // Application ID, Story ID
	Version         int         `json:"version"`
	IsAuthenticated bool        `json:"isAuthenticated"`
	Meta            Meta        `json:"meta"`
}
