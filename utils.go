package scraper

import (
	"fmt"
	"strings"
)

func buildStoreFront(cc string, l string) string {
	p, ok := PlatformIDs["iphone"]
	if !ok {
		p = "29"
	}

	sf, ok := StoreFronts[strings.ToUpper(cc)] // FIXME case
	if !ok {
		return ""
	}

	appleLanguageCode, ok := Languages[l]
	if !ok {
		return fmt.Sprintf("%s,%s t:apps3", sf, p)
	}

	return fmt.Sprintf("%s-%s,%s t:apps3", sf, appleLanguageCode, p)
}
