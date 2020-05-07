package scraper

// Result contains an Application structure or a Story structure.
// EditorialArtwork can be: subscriptionHero, brandLogo, originalFlowcaseBrick,
// storeFlowcase, bannerUber, dayCard, mediaCard, generalCard.
type Result struct {
	ArtistID                               string                   `json:"artistId"`
	ArtistName                             string                   `json:"artistName"`
	ArtistURL                              string                   `json:"artistUrl"`
	Artwork                                Artwork                  `json:"artwork"`
	BundleID                               string                   `json:"bundleId"`
	Canvas                                 Canvas                   `json:"canvas"`
	CardDisplayStyle                       string                   `json:"cardDisplayStyle"`
	CardIds                                []string                 `json:"cardIds"`
	ChartPositionForStore                  map[string]ChartPosition `json:"chartPositionForStore"`
	СircularArtwork                        Artwork                  `json:"circularArtwork"`
	ContentRatingsBySystem                 ContentRatingsBySystem   `json:"contentRatingsBySystem"`
	Copyright                              string                   `json:"copyright"`
	Description                            Description              `json:"description"`
	DeviceFamilies                         []string                 `json:"deviceFamilies"`
	DisplayStyle                           string                   `json:"displayStyle"`
	DisplaySubStyle                        string                   `json:"displaySubStyle"`
	EditorialArtwork                       map[string]Artwork       `json:"editorialArtwork"`
	EditorialNotes                         EditorialNotes           `json:"editorialNotes"`
	FirstVersionSupportingInAppPurchaseAPI string                   `json:"firstVersionSupportingInAppPurchaseApi"`
	Genres                                 Genres                   `json:"genres"`
	GenreNames                             []string                 `json:"genreNames"`
	HasInAppPurchases                      bool                     `json:"hasInAppPurchases"`
	ID                                     string                   `json:"id"`
	ItunesNotes                            ItunesNotes              `json:"itunesNotes"`
	IsAppleWatchSupported                  bool                     `json:"isAppleWatchSupported"`
	Kind                                   string                   `json:"kind"`
	Label                                  string                   `json:"label"`
	Link                                   Link                     `json:"link"`
	MinimumOSVersion                       string                   `json:"minimumOSVersion"`
	Name                                   string                   `json:"name"`
	NameRaw                                string                   `json:"nameRaw"`
	NameSortValue                          string                   `json:"nameSortValue"`
	Offers                                 Offers                   `json:"offers"`
	OvalArtwork                            Artwork                  `json:"ovalArtwork"`
	ReleaseDate                            string                   `json:"releaseDate"`
	RelatedContent                         map[string]Result        `json:"relatedContent"`
	RequiredCapabilities                   string                   `json:"requiredCapabilities"`
	ScreenshotsByType                      map[string][]Artwork     `json:"screenshotsByType"`
	ShortURL                               string                   `json:"shortUrl"`
	Subtitle                               string                   `json:"subtitle"`
	URL                                    string                   `json:"url"`
	UserRating                             UserRating               `json:"userRating"`
}