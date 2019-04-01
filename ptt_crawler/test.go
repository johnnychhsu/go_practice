package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func connect(url string) *http.Response {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return res
}

func main() {
	res, err := http.Get("https://www.ptt.cc/bbs/Beauty/M.1554027077.A.69E.html")
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
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		imageURL, ok := s.Attr("href")
		matched, _ := regexp.MatchString("https://i.imgur.com", imageURL)
		if ok && matched {
			imageURLList := strings.Split(imageURL, "/")
			imageID := imageURLList[len(imageURLList)-1]

			fileName := "img/" + imageID
			file, err := os.Create(fileName)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			imgRes := connect(imageURL)
			fmt.Println(fileName)
			_, err = io.Copy(file, imgRes.Body)
			if err != nil {
				log.Fatal(err)
			}
			defer imgRes.Body.Close()
		}
	})
}
