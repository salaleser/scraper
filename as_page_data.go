package scraper

type PageData struct {
	ComponentName           string                 `json:"componentName"`
	MetricsBase             MetricsBase            `json:"metricsBase"`
	Metrics                 Metrics                `json:"metrics"`
	Today                   Today                  `json:"today"`
	PageType                string                 `json:"pageType"`
	URL                     string                 `json:"url"`
	Term                    string                 `json:"term"`
	TopApps                 TopApps                `json:"topApps"`
	RatingAndAdvisories     RatingAndAdvisories    `json:"rating-and-advisories"`
	KindExtID               string                 `json:"kindExtId"`
	UserReviewsSortOptions  UserReviewsSortOptions `json:"userReviewsSortOptions"`
	CustomerReviewsURL      string                 `json:"customerReviewsUrl"`
	IsFatBinary             int                    `json:"isFatBinary"`
	UserReviewList          UserReviewList         `json:"userReviewList"`
	TotalNumberOfReviews    int                    `json:"totalNumberOfReviews"`
	VersionHistory          VersionHistory         `json:"versionHistory"`
	CustomersAlsoBoughtApps []string               `json:"customersAlsoBoughtApps"`
	KindName                string                 `json:"kindName"`
	ID                      string                 `json:"id"`
	AppRatingsLearnMoreURL  string                 `json:"appRatingsLearnMoreUrl"`
	SellerLabel             string                 `json:"sellerLabel"`
	MoreByThisDeveloper     []string               `json:"moreByThisDeveloper"`
	KindID                  int                    `json:"kindId"`
	Bubbles                 Bubbles                `json:"bubbles"`
	Sf6ResourceImagePath    string                 `json:"sf6ResourceImagePath"`
}
