package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func pttCrawler(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".title").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		ref, ok := s.Find("a").Attr("href")
		if ok {
			fmt.Printf("Found %s, link : %s\n", title, ref)
		}
	})
}

func main() {
	pttCrawler("https://www.ptt.cc/bbs/Beauty/index.html")
}
