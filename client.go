package scraper

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const asUserAgent = "AppStore/3.0 iOS/11.1.1 model/iPhone6,2 hwp/s5l8960x build/15B150 (6; dt:90)"
const asStoreFront = "143469-16,29 t:apps3"
const asProxyURL = "http://176.9.112.168:5005"

var debug = false

// Story returns a Story by its ID.
func Story(storyID int, cc string, l string) (Page, error) {
	const errMsg = "[ERR] scraper.AsStory(%d,%s,%s): %v\n"
	const baseURL = "https://apps.apple.com/%s/story/id%d"
	uri, err := url.Parse(fmt.Sprintf(baseURL, cc, storyID))
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, storyID, cc, l, err)
		return Page{}, nil
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, storyID, cc, l, err)
		return Page{}, nil
	}
	query.Add("cc", cc)
	uri.RawQuery = query.Encode()

	storeFront := buildStoreFront(cc, l)

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", storeFront)
	req.Header.Add("user-agent", asUserAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, storyID, cc, l, err)
		return Page{}, nil
	}

	if resp.StatusCode != 200 {
		return Page{}, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, storyID, cc, l, err)
		return Page{}, nil
	}

	page, err := ParsePage(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, storyID, cc, l, err)
		return Page{}, err
	}

	return page, nil
}

// Suggestions returns suggestions by a keyword.
func Suggestions(keyword string, cc string, l string) []byte {
	const errMsg = "[ERR] scraper.AsSuggestions(%s,%s,%s): %v\n"
	const baseURL = "https://search.itunes.apple.com/WebObjects/MZSearchHints.woa/wa/hints"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return nil
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return nil
	}
	query.Add("clientApplication", "Software")
	query.Add("caller", "com.apple.AppStore")
	query.Add("version", "1")
	query.Add("term", keyword)
	uri.RawQuery = query.Encode()

	storeFront := buildStoreFront(cc, l)

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", storeFront)
	req.Header.Add("user-agent", asUserAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return nil
	}

	return body
}

// Genre returns App Store root page for Genre structure
func Genre(id int, cc string) (Page, error) {
	const errMsg = "[ERR] scraper.AsGenre(%d,%s): %v\n"
	const baseURL = "https://itunes.apple.com/%s/genre"
	uri, err := url.Parse(fmt.Sprintf(baseURL, cc))
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, err)
		return Page{}, err
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, err)
		return Page{}, err
	}
	query.Add("id", strconv.Itoa(id))
	uri.RawQuery = query.Encode()

	storeFront := buildStoreFront(cc, "")

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", storeFront)
	req.Header.Add("user-agent", asUserAgent)

	proxyURL, err := url.Parse(asProxyURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, err)
		return Page{}, err
	}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	client = &http.Client{} // использовать ли прокси TODO
	if debug {
		log.Printf("[DBG] %s (%s): %s", cc, storeFront, req.URL.String())
	}

	resp, err := client.Do(req)
	if err != nil {
		if debug {
			fmt.Fprintf(os.Stderr, errMsg, id, cc, err)
		}
		return Page{}, err
	}

	if resp.StatusCode != 200 {
		return Page{}, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, err)
		return Page{}, err
	}

	page, err := ParsePage(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, err)
		return Page{}, err
	}

	return page, nil
}

// Grouping returns App Store root page for Grouping structure
func Grouping(id int, cc string, l string) (Page, error) {
	const errMsg = "[ERR] scraper.AsGrouping(%d,%s,%s): %v\n"
	const baseURL = "http://itunes.apple.com/WebObjects/MZStore.woa/wa/viewGrouping"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}
	query.Add("cc", cc)
	query.Add("id", strconv.Itoa(id))
	uri.RawQuery = query.Encode()

	storeFront := buildStoreFront(cc, l)

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", storeFront)
	req.Header.Add("user-agent", asUserAgent)

	proxyURL, err := url.Parse(asProxyURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	client = &http.Client{} // TODO Использовать ли прокси
	if debug {
		log.Printf("[DBG] %s (%s): %s", cc, storeFront, req.URL.String())
	}

	resp, err := client.Do(req)
	if err != nil {
		if debug {
			fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		}
		return Page{}, err
	}

	if resp.StatusCode != 200 {
		return Page{}, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	page, err := ParsePage(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, id, cc, l, err)
		return Page{}, err
	}

	return page, nil
}

// Room returns a Room by its ID.
func Room(fcID int, cc string, l string) (Page, error) {
	const errMsg = "[ERR] scraper.AsRoom(%d,%s,%s): %v\n"
	const baseURL = "https://itunes.apple.com/WebObjects/MZStore.woa/wa/viewRoom"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, fcID, cc, l, err)
		return Page{}, err
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, fcID, cc, l, err)
		return Page{}, err
	}
	query.Add("fcId", strconv.Itoa(fcID))
	// query.Add("genreIdString", "6014")                           // TODO изучить
	// query.Add("mediaTypeString", "Mobile+Software+Applications") // TODO изучить
	uri.RawQuery = query.Encode()

	storeFront := buildStoreFront(cc, l)

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", storeFront)
	req.Header.Add("user-agent", asUserAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, fcID, cc, l, err)
		return Page{}, err
	}

	if resp.StatusCode != 200 {
		return Page{}, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, fcID, cc, l, err)
		return Page{}, err
	}

	page, err := ParsePage(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, fcID, cc, l, err)
		return Page{}, err
	}

	return page, nil
}

// AppIDs returns application IDs by a keyword.
func AppIDs(keyword string, cc string, l string) []MetadataResponse {
	const errMsg = "[ERR] scraper.AsAppIDs(%s,%s,%s): %v\n"
	const baseURL = "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return []MetadataResponse{}
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return []MetadataResponse{}
	}
	query.Add("clientApplication", "Software")
	query.Add("caller", "com.apple.AppStore")
	query.Add("version", "1")
	query.Add("term", keyword)
	uri.RawQuery = query.Encode()

	storeFront := buildStoreFront(cc, l)

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", storeFront)
	req.Header.Add("user-agent", asUserAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return []MetadataResponse{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, cc, l, err)
		return []MetadataResponse{}
	}

	return parseAsIDs(body)
}

// Metadata returns an Application's metadata by its ID.
func Metadata(appID string, cc string, l string) MetadataResponse {
	const errMsg = "[ERR] scraper.AsMetadata(%s,%s,%s): %v\n"
	const baseURLpart = "https://apps.apple.com/%s/app/id%s"
	uri, err := url.Parse(fmt.Sprintf(baseURLpart, cc, appID))
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, appID, cc, l, err)
		return MetadataResponse{}
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, appID, cc, l, err)
		return MetadataResponse{}
	}
	uri.RawQuery = query.Encode()

	storeFront := buildStoreFront(cc, l)

	req, err := http.NewRequest("GET", uri.String(), nil)
	req.Header.Add("x-apple-store-front", storeFront)
	req.Header.Add("user-agent", asUserAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, appID, cc, l, err)
		return MetadataResponse{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, appID, cc, l, err)
		return MetadataResponse{}
	}

	return parseAsMetadata(body)
}

// GpAppIDs returns application IDs by a keyword.
func GpAppIDs(keyword string, gl string, hl string) []MetadataResponse {
	const errMsg = "[ERR] scraper.GpAppIDs(%s,%s,%s): %v\n"
	const baseURL = "https://play.google.com/_/PlayStoreUi/data/batchexecute"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, gl, hl, err)
		return []MetadataResponse{}
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, gl, hl, err)
		return []MetadataResponse{}
	}
	query.Add("gl", gl)
	query.Add("hl", hl)
	uri.RawQuery = query.Encode()

	value := fmt.Sprintf("[[[lGYRle,'[[null,[[10,[10,%d]],true,null,[96,27,4,8,57,30,110,79,11,16,49,1,3,9,12,104,55,56,51,10,34,77]],[%s],4,null,null,null,[]]]',null,%s]]]", 5, keyword, keyword)

	data := url.Values{}
	data.Add("f.req", value)

	resp, err := http.PostForm(baseURL, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, gl, hl, err)
		return []MetadataResponse{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, keyword, gl, hl, err)
		return []MetadataResponse{}
	}

	return parseGpIDs(body[5:])
}

// GpMetadata returns an Application's metadata by its ID.
func GpMetadata(appID string, gl string, hl string) MetadataResponse {
	const errMsg = "[ERR] scraper.GpMetadata(%s,%s,%s): %v\n"
	const baseURL = "https://play.google.com/_/PlayStoreUi/data/batchexecute"
	uri, err := url.Parse(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, appID, gl, hl, err)
		return MetadataResponse{}
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, appID, gl, hl, err)
		return MetadataResponse{}
	}
	query.Add("gl", gl)
	query.Add("hl", hl)
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
		fmt.Fprintf(os.Stderr, errMsg, appID, gl, hl, err)
		return MetadataResponse{}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, errMsg, appID, gl, hl, err)
		return MetadataResponse{}
	}

	return parseGpMetadata(body[5:])
}
