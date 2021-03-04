package engine

import (
	"gostudy/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, request := range seeds {
		requests = append(requests, request)
	}
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", request.Url)
		body, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", request.Url, err)
			continue
		}
		parseResult := request.ParserFunc(body)
		requests = append(requests, parseResult.Request...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %s", item)
		}
	}
}
