package popular

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"icodelife.com/itorrent-gofiber/internal"
	"icodelife.com/itorrent-gofiber/models"
)

func GetPopularTorrent(res *fiber.Ctx) {
	queryStr := strings.ReplaceAll(res.Query("t"), " ", "+")
	c := colly.NewCollector()
	torrents := make([]models.Torrent, 0)
	c.OnHTML("body", func(element *colly.HTMLElement) {

		element.DOM.Find(".tgxtablerow.txlight").Each(func(i int, selection *goquery.Selection) {

			t := models.Torrent{}

			t.Title, _ = selection.Find("#click").Find("a").First().Attr("title")

			selection.Find("a").Each(func(i int, selection *goquery.Selection) {
				if i == 3 {
					t.Link, _ = selection.Attr("href")
				}
				if i == 4 {
					t.Magnet, _ = selection.Attr("href")
				}
			})

			selection.Find(".tgxtablecell").Each(func(i int, selection *goquery.Selection) {
				t.Size = selection.Find(".badge-secondary").Text()
				t.Seeds = selection.Find("[title=\"Seeders/Leechers\"]").Find("[color=\"green\"]").Text()
				t.Leechs = selection.Find("[title=\"Seeders/Leechers\"]").Find("[color=\"#ff0000\"]").Text()
			})

			t.Date = selection.Find("small").Last().Text()

			torrents = append(torrents, t)
		})
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.Body, r.StatusCode, "\nError:", err.Error())
	})

	err := c.Visit(internal.TGXUrl + "torrents.php?search=" + queryStr)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.Wait()
	res.JSON(torrents)
}
