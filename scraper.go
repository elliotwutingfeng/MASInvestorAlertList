package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/elliotwutingfeng/go-fasttld"
	"mvdan.cc/xurls/v2"
)

func unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func main() {
	urlSchemeRegex := regexp.MustCompile(`^[A-Za-z0-9+-.]+://`)
	zerowidthRegex := regexp.MustCompile("^[\u200B-\u200D\uFEFF]")
	whitespaceRegex := regexp.MustCompile("\\s+")

	resp, err := http.Get("https://www.mas.gov.sg/api/v1/ialsearch?json.nl=map&wt=json&rows=10000&q=*:*&sort=date_dt+desc&start=0")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var bodyJSON map[string]interface{}
	json.Unmarshal([]byte(body), &bodyJSON)

	response, ok := bodyJSON["response"].(map[string]interface{})
	if !ok {
		log.Fatalln(`"response" not in bodyJSON`)
	}

	docs, ok := response["docs"].([]interface{})
	if !ok {
		log.Fatalln(`"docs" not in bodyJSON`)
	}

	links := []string{}
	rxRelaxed := xurls.Relaxed()
	for _, doc := range docs {
		if d, ok := doc.(map[string]interface{}); ok {
			websiteS, ok := d["website_s"].(string)
			if !ok {
				continue
			}
			links = append(links, rxRelaxed.FindAllString(websiteS, -1)...)
		}
	}

	urls := make(map[string]bool)

	extractor, _ := fasttld.New(fasttld.SuffixListParams{})

	for _, maybeURL := range links {
		// remove zero-width spaces
		rawURL := zerowidthRegex.ReplaceAllLiteralString(maybeURL, "")
		// remove whitespaces
		rawURL = whitespaceRegex.ReplaceAllLiteralString(rawURL, "")
		// remove url scheme (e.g. http:// https:// etc.)
		rawURL = urlSchemeRegex.ReplaceAllString(rawURL, "")
		// remove whitespace on both ends
		rawURL = strings.TrimSpace(rawURL)
		// remove trailing slash
		rawURL = strings.TrimRight(rawURL, "/")

		// convert hostname to lowercase
		res, _ := extractor.Extract(fasttld.URLParams{URL: rawURL})
		subDomain := res.SubDomain
		registeredDomain := res.RegisteredDomain
		hostname := registeredDomain
		if len(subDomain) != 0 {
			hostname = subDomain + "." + registeredDomain
		}
		port := res.Port
		if len(port) != 0 {
			port = ":" + port
		}
		url := strings.ToLower(hostname) + port + res.Path
		urls[url] = true
	}

	// get map keys as slice
	sortedURLs := make([]string, len(urls))
	i := 0
	for k := range urls {
		sortedURLs[i] = k
		i++
	}

	sort.Strings(sortedURLs) // sort alphabetically

	if len(sortedURLs) > 0 {
		err = os.WriteFile("mas-investor-alert-list.txt", []byte(strings.Join(sortedURLs, "\n")), 0644)
	} else {
		log.Fatal("No URLs found")
	}
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
