package scraper

type Offers []struct {
	ActionText     ActionText `json:"actionText"`
	Type           string     `json:"type"`
	PriceFormatted string     `json:"priceFormatted"`
	Price          int        `json:"price"`
	BuyParams      string     `json:"buyParams"`
	Version        Version    `json:"version"`
	Assets         Assets     `json:"assets"`
}
