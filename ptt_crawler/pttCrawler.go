package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
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

func picSave(url string) {
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
			defer res.Body.Close()
			_, err = io.Copy(file, imgRes.Body)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
}

func pagePic(pageURLprefix string, pNum int) {
	url := pageURLprefix + strconv.Itoa(pNum) + ".html"
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
		matched, _ := regexp.MatchString("[正妹]", title)
		if ok && matched {
			ref := "https://www.ptt.cc" + ref
			fmt.Printf("Found %s, link : %s\n", title, ref)
			picSave(ref)
		}
	})
}

func pttCrawler(url string, targetPage int) {
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

	// get page number
	doc.Find("a.btn").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		matched, _ := regexp.MatchString("上頁", text)
		ref, ok := s.Attr("href")
		if matched && ok {
			numPages := strings.Split(ref, "/")[3]
			numPages = strings.Split(numPages, ".")[0]
			numPages = strings.TrimPrefix(numPages, "index")

			pageURLPrefix := "https://www.ptt.cc/bbs/Beauty/index"
			s, err := strconv.Atoi(numPages)
			if err != err {
				log.Fatal(err)
			}
			for i := s; i >= s-targetPage; i-- {
				pagePic(pageURLPrefix, i)
			}
		}
	})
}

func main() {
	board := "https://www.ptt.cc/bbs/Beauty/index.html"
	targetPage := 1
	pttCrawler(board, targetPage)
}
