package scraper

type Model struct {
	AdamID             string    `json:"adamId"`
	Children           []Model   `json:"children,omitempty"`
	Content            []Content `json:"content,omitempty"`
	DesignBadge        string    `json:"designBadge,omitempty"`
	DisplayStyle       string    `json:"displayStyle,omitempty"`
	DoNotFilter        bool      `json:"doNotFilter"`
	FcKind             string    `json:"fcKind"`
	Link               Link      `json:"link,omitempty"`
	Links              []Link    `json:"links,omitempty"`
	Name               string    `json:"name,omitempty"`
	SeeAllURL          string    `json:"seeAllUrl,omitempty"`
	ShouldUseGradients bool      `json:"shouldUseGradients"`
	SuppressTagline    string    `json:"suppressTagline,omitempty"`
	Tagline            string    `json:"tagline,omitempty"`
}
