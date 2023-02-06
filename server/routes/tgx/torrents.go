package tgx

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"icodelife.com/itorrent-gofiber/models"
)

func getTorrentData(links []string) []models.HotPicks {
	c := colly.NewCollector()
	allHotpicks := make([]models.HotPicks, 0)
	c.OnHTML("body", func(element *colly.HTMLElement) {
		hotpicks := models.HotPicks{}

		torrents := make([]models.Torrent, 0)

		element.DOM.Find(".tgxtablerow.txlight").Each(func(i int, selection *goquery.Selection) {
			t := models.Torrent{}

			t.Title, _ = selection.Find("#click").Find("a").First().Attr("title")

			//edge case
			// replace empty torrent title with of imdb
			// if t.Title == "" {
			// 	t.Title = imdb.Title
			// }
			selection.Find("a").Each(func(i int, selection *goquery.Selection) {
				if i == 2 {
					t.Magnet, _ = selection.Attr("href")
				}
			})
			selection.Find(".tgxtablecell").Each(func(i int, selection *goquery.Selection) {
				t.Size = selection.Find(".badge-secondary").Text()
				t.Seeds = selection.Find("[title=\"Seeders/Leechers\"]").Find("[color=\"green\"]").Text()
				t.Leechs = selection.Find("[title=\"Seeders/Leechers\"]").Find("[color=\"#ff0000\"]").Text()
			})
			torrents = append(torrents, t)
		})
		hotpicks.Torrents = torrents
		allHotpicks = append(allHotpicks, hotpicks)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.Body, r.StatusCode, "\nError:", err.Error())
	})
	leng := len(links)
	if leng > 15 {
		leng = 15
	}
	for i := 0; i < leng; i++ {
		err := c.Visit(links[i])
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	c.Wait()
	return allHotpicks
}
