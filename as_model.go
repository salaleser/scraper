package scraper

type Model struct {
	FcKind      string `json:"fcKind"`
	AdamID      string `json:"adamId"`
	DoNotFilter bool   `json:"doNotFilter"`
	Children    []struct {
		FcKind      string `json:"fcKind"`
		AdamID      string `json:"adamId"`
		DoNotFilter bool   `json:"doNotFilter"`
		Children    []struct {
			FcKind      string `json:"fcKind"`
			AdamID      string `json:"adamId"`
			DoNotFilter bool   `json:"doNotFilter"`
			Children    []struct {
				FcKind      string `json:"fcKind"`
				AdamID      string `json:"adamId"`
				DoNotFilter bool   `json:"doNotFilter"`
				Link        Link   `json:"link"`
				DesignBadge string `json:"designBadge"`
			} `json:"children,omitempty"`
			Name            string `json:"name,omitempty"`
			DisplayStyle    string `json:"displayStyle,omitempty"`
			SeeAllURL       string `json:"seeAllUrl,omitempty"`
			SuppressTagline string `json:"suppressTagline,omitempty"`
			Content         []struct {
				Type      string `json:"type"`
				KindIds   []int  `json:"kindIds"`
				ContentID string `json:"contentId"`
			} `json:"content,omitempty"`
			Tagline string `json:"tagline,omitempty"`
			Links   []struct {
				Type   string `json:"type"`
				Label  string `json:"label"`
				URL    string `json:"url"`
				Target string `json:"target"`
			} `json:"links,omitempty"`
		} `json:"children"`
	} `json:"children"`
	ShouldUseGradients bool `json:"shouldUseGradients"`
}
