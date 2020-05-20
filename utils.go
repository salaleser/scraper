package scraper

func buildStoreFront(location string, language string) string {
	if storeFront, ok := StoreFronts[location]; ok {
		return storeFront
	}

	return ""
}
