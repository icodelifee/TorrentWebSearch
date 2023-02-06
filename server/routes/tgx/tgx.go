package tgx

import (
	"fmt"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"icodelife.com/itorrent-gofiber/models"
)

func GetTGX(res *fiber.Ctx) error {
	queryStr := strings.ReplaceAll(res.Query("q"), " ", "+")
	url := baseUrl + "torrents.php?search=" + queryStr
	c := colly.NewCollector()
	torrents := make([]models.Torrent, 0)

	var wg sync.WaitGroup

	c.OnHTML("body", func(element *colly.HTMLElement) {
		element.DOM.Find(".tgxtablerow.txlight").Each(func(i int, selection *goquery.Selection) {
			t := models.Torrent{}

			t.Title, _ = selection.Find("#click").Find("a").First().Attr("title")

			magnet := selection.Find("a").Nodes[4].Attr[0].Val
			if strings.Contains(magnet, "magnet") {
				t.Magnet = magnet
			} else {
				t.Magnet = selection.Find("a").Nodes[3].Attr[0].Val
			}
			selection.Find(".tgxtablecell").Each(func(i int, selection *goquery.Selection) {
				t.Size = selection.Find(".badge-secondary").Text()
				t.Seeds = selection.Find("[title=\"Seeders/Leechers\"]").Find("[color=\"green\"]").Text()
				t.Leechs = selection.Find("[title=\"Seeders/Leechers\"]").Find("[color=\"#ff0000\"]").Text()
			})
			torrents = append(torrents, t)
		})
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Printf("{ response : %d, error: %s}", response.StatusCode, err)
	})

	c.OnScraped(func(response *colly.Response) {
		wg.Wait()
	})
	_ = c.Visit(url)

	return res.Status(200).JSON(torrents)

}
