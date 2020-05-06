package scraper

type Story struct {
	Canvas           Canvas             `json:"canvas"`
	Label            string             `json:"label"`
	ID               string             `json:"id"`
	CardIds          []string           `json:"cardIds"`
	RelatedContent   map[string]Result  `json:"relatedContent"`
	EditorialArtwork map[string]Artwork `json:"editorialArtwork"`
	Kind             string             `json:"kind"`
	Link             Link               `json:"link"`
	DisplayStyle     string             `json:"displayStyle"`
	EditorialNotes   EditorialNotes     `json:"editorialNotes"`
	CardDisplayStyle string             `json:"cardDisplayStyle"`
	DisplaySubStyle  string             `json:"displaySubStyle"`
}

// EditorialArtwork:
// SubscriptionHero      Artwork `json:"subscriptionHero"`
// BrandLogo             Artwork `json:"brandLogo"`
// OriginalFlowcaseBrick Artwork `json:"originalFlowcaseBrick"`
// StoreFlowcase         Artwork `json:"storeFlowcase"`
// BannerUber            Artwork `json:"bannerUber"`
// DayCard               Artwork `json:"dayCard"`
// MediaCard             Artwork `json:"mediaCard"`
// GeneralCard           Artwork `json:"generalCard"`
