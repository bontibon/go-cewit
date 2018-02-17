package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var (
		doc *goquery.Document
		err error
	)

	// Offline cache in case of network issues
	f, err := os.Open("09-scraping/cache.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err = goquery.NewDocumentFromReader(f)

	// START OMIT
	// doc, err = goquery.NewDocument("http://hack.cewit.org/")
	if err != nil {
		log.Fatal(err)
	}
	s := doc.Find(`h1`).
		FilterFunction(func(i int, s *goquery.Selection) bool {
			return s.Text() == "Prizes"
		}).
		First().
		Closest("section").
		Find("div.slick-track > div > h3")

	prizes := make(map[string]bool, s.Length())
	s.Each(func(i int, s *goquery.Selection) {
		prizes[s.Text()] = true
	})
	for prize := range prizes {
		fmt.Println(prize)
	}
	// END OMIT
}
