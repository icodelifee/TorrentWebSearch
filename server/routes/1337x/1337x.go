package _337x

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"icodelife.com/itorrent-gofiber/internal"
	"icodelife.com/itorrent-gofiber/models"
)

func Get1337x(res *fiber.Ctx) error {
	queryStr := strings.ReplaceAll(res.Query("q"), " ", "+")
	sortStr := res.Query("s")
	orderStr := res.Query("o")
	var searchUrl string
	if sortStr == "0" || sortStr == "" {
		searchUrl = fmt.Sprintf(internal.LeetxBaseUrl, queryStr)
	} else {
		searchUrl = fmt.Sprintf(internal.LeetxSortUrl, queryStr, getSort(sortStr), getOrder(orderStr))
	}
	c := colly.NewCollector()
	torrents := make([]models.Torrent, 0)

	var wg sync.WaitGroup

	c.OnHTML("body", func(element *colly.HTMLElement) {

		// if there are no torrents return function
		if element.DOM.Find("tr").Length() == 0 {
			return
		}

		element.ForEach("tr", func(i int, node *colly.HTMLElement) {
			// skip first element
			t := models.Torrent{}

			if i == 0 {
				return
			}

			t.Title = node.DOM.Find("td.coll-1.name").Children().Nodes[1].FirstChild.Data
			t.Link = node.DOM.Find("td.coll-1.name").Children().Nodes[1].Attr[0].Val
			t.Seeds = node.ChildText("td.coll-2.seeds")
			t.Leechs = node.ChildText("td.coll-3.leeches")
			t.Date = node.ChildText("td.coll-date")
			t.Size = node.DOM.Find("td.coll-4.size").Nodes[0].FirstChild.Data

			wg.Add(1)

			// go routine call
			go func(t *models.Torrent) {
				defer wg.Done()
				c := colly.NewCollector()
				c.OnHTML(".col-9.page-content", func(element *colly.HTMLElement) {
					t.Magnet = element.DOM.Find("a").Nodes[0].Attr[1].Val
				})
				c.OnScraped(func(response *colly.Response) {
					torrents = append(torrents, *t)
				})
				_ = c.Visit("https://www.1377x.to" + t.Link)
			}(&t)
		})

	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Printf("{ response : %d, error: %s}", response.StatusCode, err)
	})

	c.OnScraped(func(response *colly.Response) {
		wg.Wait()
		if len(torrents) > 0 {
			_ = res.Status(200).JSON(torrents)
		} else {
			_ = res.Status(204)
		}
	})
	_ = c.Visit(searchUrl)

	return nil
}
func getOrder(order string) string {
	if order == "0" {
		return "asc"
	} else {
		return "desc"
	}
}

func getSort(sort string) string {
	var str string
	switch sort {
	case "1":
		str = "seeders"
		break
	case "2":
		str = "time"
		break
	case "3":
		str = "size"
	default:
		str = "seeders"
		break
	}
	return str
}
