package Webscraping

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func Scrap() ([]string, []string) {

	titles := []string{}
	stars := []string{}

	res, _ := http.Get("https://www.imdb.com/chart/top")

	if res.StatusCode != 200 {
		fmt.Println("Some Error...", res.StatusCode)
		return nil, nil
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	go doc.Find(".titleColumn").Each(func(i int, selection *goquery.Selection) {
		title := selection.Find("a").Text()
		titles = append(titles, title)

	})

	doc.Find(".ratingColumn.imdbRating").Each(func(i int, selection *goquery.Selection) {
		star := selection.Find("strong").Text()
		stars = append(stars, star)
	})

	return titles, stars
}
