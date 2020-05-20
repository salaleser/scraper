package scraper

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const asUserAgent = "AppStore/3.0 iOS/11.1.1 model/iPhone6,2 hwp/s5l8960x build/15B150 (6; dt:90)"
const asStoreFront = "143469-16,29 t:apps3"
const asProxyURL = "http://176.9.112.168:5005"

var debug = false

// AsStory returns a Story by its ID.
func AsStory(storyID string, location string, language string) StoryResponse {
	const baseURL = "https://apps.apple.com/%s/story/id%s"
	uri, err := url.Parse(fmt.Sprintf(baseURL, location, storyID))
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as story url: %v\n", err)
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as story query: %v\n", err)
	}
	query.Add("cc", location)
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", asStoreFront) // TODO учесть другие страны
	req.Header.Add("user-agent", asUserAgent)           // TODO
	req.Header.Add("x-apple-i-timezone", "GMT+3")       // TODO
	req.Header.Add("Host", "itunes.apple.com")          // TODO

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "app store story request: %v\n", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading as story response body: %v\n", err)
	}

	return parseAsStory(body[:])
}

// AsSuggestions returns suggestions by a keyword.
func AsSuggestions(keyword string, location string, language string) []byte {
	const baseURL = "https://search.itunes.apple.com/WebObjects/MZSearchHints.woa/wa/hints"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as url: %v\n", err)
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as query: %v\n", err)
	}
	query.Add("clientApplication", "Software")
	query.Add("caller", "com.apple.AppStore")
	query.Add("version", "1")
	query.Add("term", keyword)
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", asStoreFront) // TODO учесть другие страны
	req.Header.Add("user-agent", asUserAgent)           // TODO
	req.Header.Add("x-apple-i-timezone", "GMT+3")       // TODO
	req.Header.Add("Host", "search.itunes.apple.com")   // TODO

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

// AsGrouping returns group TODO.
func AsGrouping(id string, location string, language string) ([]byte, error) {
	const baseURL = "http://itunes.apple.com/WebObjects/MZStore.woa/wa/viewGrouping"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as grouping url: %v\n", err)
		return nil, err
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as grouping query: %v\n", err)
		return nil, err
	}
	query.Add("cc", location)
	query.Add("id", id)
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", buildStoreFront(location, language))
	req.Header.Add("user-agent", asUserAgent) // TODO

	proxyURL, err := url.Parse(asProxyURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as grouping proxy url: %v\n", err)
		return nil, err
	}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	if debug {
		log.Printf("[DBG] request: %s", req.URL.String())
	}

	resp, err := client.Do(req)
	if err != nil {
		if debug {
			fmt.Fprintf(os.Stderr, "app store grouping request: %v\n", err)
		}
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading as grouping response body: %v\n", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		return body, errors.New(resp.Status)
	}

	return body, nil
}

// AsRoom returns a Room by its ID.
func AsRoom(adamID string, location string, language string) RoomResponse {
	const baseURL = "https://itunes.apple.com/WebObjects/MZStore.woa/wa/viewRoom"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as room url: %v\n", err)
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as room query: %v\n", err)
	}
	query.Add("fcId", adamID)
	// query.Add("genreIdString", "6014")                           // TODO изучить
	// query.Add("mediaTypeString", "Mobile+Software+Applications") // TODO изучить
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", asStoreFront) // TODO учесть другие страны
	req.Header.Add("user-agent", asUserAgent)           // TODO
	req.Header.Add("x-apple-i-timezone", "GMT+3")       // TODO
	req.Header.Add("Host", "itunes.apple.com")          // TODO

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "app store room request: %v\n", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading as room response body: %v\n", err)
	}

	return parseAsRoom(body[:])
}

// AsAppIDs returns application IDs by a keyword.
func AsAppIDs(keyword string, location string, language string) []MetadataResponse {
	const baseURL = "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as url: %v\n", err)
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as query: %v\n", err)
	}
	query.Add("clientApplication", "Software")
	query.Add("caller", "com.apple.AppStore")
	query.Add("version", "1")
	query.Add("term", keyword)
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", asStoreFront) // TODO учесть другие страны
	req.Header.Add("user-agent", asUserAgent)           // TODO
	req.Header.Add("x-apple-i-timezone", "GMT+3")       // TODO
	req.Header.Add("Host", "search.itunes.apple.com")   // TODO

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "app store request: %v\n", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading as response body: %v\n", err)
	}

	return parseAsIDs(body[:])
}

// AsMetadata returns an Application's metadata by its ID.
func AsMetadata(appID string, location string, language string) MetadataResponse {
	const baseURLpart = "https://apps.apple.com/ru/app/id"
	uri, err := url.Parse(baseURLpart + appID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as url: %v\n", err)
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing as query: %v\n", err)
	}
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", asStoreFront) // TODO учесть другие страны
	req.Header.Add("user-agent", asUserAgent)           // TODO
	req.Header.Add("x-apple-i-timezone", "GMT+3")       // TODO
	req.Header.Add("Host", "apps.apple.com")            // TODO

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "app store request: %v\n", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading as response body: %v\n", err)
	}

	return parseAsMetadata(body[:])
}

// GpAppIDs returns application IDs by a keyword.
func GpAppIDs(keyword string, location string, language string) []MetadataResponse {
	const baseURL = "https://play.google.com/_/PlayStoreUi/data/batchexecute"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing gp url: %v\n", err)
	}

	query, err := url.ParseQuery(uri.RawQuery)
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

	return parseGpIDs(body[5:])
}

// GpMetadata returns an Application's metadata by its ID.
func GpMetadata(appID string, location string, language string) MetadataResponse {
	const baseURL = "https://play.google.com/_/PlayStoreUi/data/batchexecute"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing gp url: %v\n", err)
	}

	query, err := url.ParseQuery(uri.RawQuery)
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

	return parseGpMetadata(body[5:])
}
