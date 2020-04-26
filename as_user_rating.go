package scraper

type UserRating struct {
	Value               float64 `json:"value"`
	RatingCount         int     `json:"ratingCount"`
	RatingCountList     []int   `json:"ratingCountList"`
	AriaLabelForRatings string  `json:"ariaLabelForRatings"`
}
