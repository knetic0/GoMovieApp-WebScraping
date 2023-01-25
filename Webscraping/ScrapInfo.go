package Webscraping

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func ScrapInfo(url string) []string {
	infos := []string{}

	res, _ := http.Get(fmt.Sprintf("https://www.rottentomatoes.com/m/%s", url))

	if res.StatusCode != 200 {
		fmt.Println("Some Error...", res.StatusCode)
		return nil
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	doc.Find(".meta-row.clearfix").Each(func(i int, selection *goquery.Selection) {
		info := selection.Find(".meta-value").Text()
		infos = append(infos, info)
	})

	return infos
}
