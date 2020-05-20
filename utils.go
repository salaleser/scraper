package scraper

import (
	"fmt"
	"strings"
)

func buildStoreFront(location string, language string) string {
	p, ok := PlatformIDs["iphone"]
	if !ok {
		p = "29"
	}

	sf, ok := StoreFronts[strings.ToUpper(location)] // FIXME case
	if !ok {
		return ""
	}

	l, ok := Languages[language]
	if !ok {
		return fmt.Sprintf("%s,%s t:apps3", sf, p)
	}

	return fmt.Sprintf("%s-%s,%s t:apps3", sf, l, p)
}
