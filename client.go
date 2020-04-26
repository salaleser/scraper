package scraper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func AsStory(location string, language string) StoryResponse {
	const baseURL = "http://itunes.apple.com/WebObjects/MZStore.woa/wa/viewToday"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as today url: %v\n", err)
	}

	query, _ := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as today query: %v\n", err)
	}
	query.Add("cc", location)
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", "143469-16,29 t:apps3")                                                // TODO учесть другие страны
	req.Header.Add("user-agent", "AppStore/3.0 iOS/11.1.1 model/iPhone6,2 hwp/s5l8960x build/15B150 (6; dt:90)") // TODO
	req.Header.Add("x-apple-i-timezone", "GMT+3")                                                                // TODO
	req.Header.Add("Host", "itunes.apple.com")                                                                   // TODO

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "app store today request: %v\n", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading as today response body: %v\n", err)
	}

	return parseAsStory(body[:])
}

// AsSuggestions returns suggestions by keyword
func AsSuggestions(keyword string, location string, language string) []byte {
	const baseURL = "https://search.itunes.apple.com/WebObjects/MZSearchHints.woa/wa/hints"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as url: %v\n", err)
	}

	query, _ := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as query: %v\n", err)
	}
	query.Add("clientApplication", "Software")
	query.Add("caller", "com.apple.AppStore")
	query.Add("version", "1")
	query.Add("term", keyword)
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", "143469-16,29 t:apps3")                                                // TODO учесть другие страны
	req.Header.Add("user-agent", "AppStore/3.0 iOS/11.1.1 model/iPhone6,2 hwp/s5l8960x build/15B150 (6; dt:90)") // TODO
	req.Header.Add("x-apple-i-timezone", "GMT+3")                                                                // TODO
	req.Header.Add("Host", "search.itunes.apple.com")                                                            // TODO

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "app store request: %v\n", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading as response body: %v\n", err)
	}

	return body[:]
}

// AsAppIDs returns application IDs by keyword
func AsAppIDs(keyword string, location string, language string) []Metadata {
	const baseURL = "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as url: %v\n", err)
	}

	query, _ := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as query: %v\n", err)
	}
	query.Add("clientApplication", "Software")
	query.Add("caller", "com.apple.AppStore")
	query.Add("version", "1")
	query.Add("term", keyword)
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", "143469-16,29 t:apps3")                                                // TODO учесть другие страны
	req.Header.Add("user-agent", "AppStore/3.0 iOS/11.1.1 model/iPhone6,2 hwp/s5l8960x build/15B150 (6; dt:90)") // TODO
	req.Header.Add("x-apple-i-timezone", "GMT+3")                                                                // TODO
	req.Header.Add("Host", "search.itunes.apple.com")                                                            // TODO

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "app store request: %v\n", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading as response body: %v\n", err)
	}

	return parseAsIDsBody(body[:])
}

// AsMetadataBody returns body
func AsMetadataBody(appID string, location string, language string) Metadata {
	const baseURLpart = "https://apps.apple.com/ru/app/id"
	uri, err := url.Parse(baseURLpart + appID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as url: %v\n", err)
	}

	query, _ := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as query: %v\n", err)
	}
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", "143469-16,29 t:apps3")                                                // TODO учесть другие страны
	req.Header.Add("user-agent", "AppStore/3.0 iOS/11.1.1 model/iPhone6,2 hwp/s5l8960x build/15B150 (6; dt:90)") // TODO
	req.Header.Add("x-apple-i-timezone", "GMT+3")                                                                // TODO
	req.Header.Add("Host", "apps.apple.com")                                                                     // TODO

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "app store request: %v\n", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading as response body: %v\n", err)
	}

	return parseAsMetadataBody(body[:])
}

// GpAppIDs returns application IDs by keyword
func GpAppIDs(keyword string, location string, language string) []Metadata {
	const baseURL = "https://play.google.com/_/PlayStoreUi/data/batchexecute"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing gp url: %v\n", err)
	}

	query, _ := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing gp query: %v\n", err)
	}
	query.Add("gl", location)
	query.Add("hl", language)
	// TODO add queries (soc-app, ...)
	uri.RawQuery = query.Encode()

	value := fmt.Sprintf("[[[lGYRle,'[[null,[[10,[10,%d]],true,null,[96,27,4,8,57,30,110,79,11,16,49,1,3,9,12,104,55,56,51,10,34,77]],[%s],4,null,null,null,[]]]',null,%s]]]", 5, keyword, keyword)

	data := url.Values{}
	data.Add("f.req", value)

	resp, err := http.PostForm(baseURL, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gp request: %v\n", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading gp resopnse body: %v\n", err)
	}

	return parseGpIDsBody(body[5:])
}

// GpMetadataBody returns body
func GpMetadataBody(appID string, location string, language string) Metadata {
	const baseURL = "https://play.google.com/_/PlayStoreUi/data/batchexecute"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing gp url: %v\n", err)
	}

	query, _ := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing gp query: %v\n", err)
	}
	query.Add("gl", location)
	query.Add("hl", language)
	// TODO add queries (soc-app, ...)
	uri.RawQuery = query.Encode()

	// v1 := fmt.Sprintf("[d5UeYe,'[[%s,7]]',null,%s]", appID, appID)        // price
	// v3 := fmt.Sprintf("[MLWfjd,'[[%s,7]]',null,%s]", appID, appID)        // rating
	// v4 := fmt.Sprintf("[IoIWBc,'[[null,[%s,7]]]',null,%s]", appID, appID) // version
	// v5 := fmt.Sprintf("[k8610b,'[[null,[%s,7]]]',null,%s]", appID, appID) // short rating
	// v6 := fmt.Sprintf("[BQ0Ly,'[[null,[%s,7]]]',null,%s]", appID, appID)  // downloads

	value := fmt.Sprintf("[[[jLZZ2e,'[[%s,7],2]',null,%s]]]", appID, appID) // ASO

	data := url.Values{}
	data.Add("f.req", value)

	resp, err := http.PostForm(baseURL, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "google play request: %v\n", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading gp response body: %v\n", err)
	}

	return parseGpMetadataBody(body[5:])
}
