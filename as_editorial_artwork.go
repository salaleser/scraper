package scraper

type EditorialArtwork struct {
	SubscriptionHero      Artwork `json:"subscriptionHero"`
	BrandLogo             Artwork `json:"brandLogo"`
	OriginalFlowcaseBrick Artwork `json:"originalFlowcaseBrick"`
	StoreFlowcase         Artwork `json:"storeFlowcase"`
	BannerUber            Artwork `json:"bannerUber"`
	DayCard               Artwork `json:"dayCard"`
	MediaCard             Artwork `json:"mediaCard"`
	GeneralCard           Artwork `json:"generalCard"`
}
