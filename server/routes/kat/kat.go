package kat

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"icodelife.com/itorrent-gofiber/internal"
	"icodelife.com/itorrent-gofiber/models"
)

func GetKat(res *fiber.Ctx) error {
	queryStr := strings.ReplaceAll(res.Query("q"), " ", "+")
	sortStr := res.Query("s")
	orderStr := res.Query("o")
	searchUrl := fmt.Sprintf(internal.KatBaseUrl, queryStr)
	if sortStr != "" {
		searchUrl = searchUrl + fmt.Sprintf("?sortby=%s&sort=%s", GetSort(sortStr), GetOrder(orderStr))
	}
	c := colly.NewCollector()

	torrents := make([]models.Torrent, 0)
	var wg sync.WaitGroup

	c.OnHTML("body", func(e *colly.HTMLElement) {
		if e.DOM.Find("span[itemprop=name]").Length() == 0 {
			e.ForEach("tr.odd , tr.even", func(i int, e *colly.HTMLElement) {
				t := models.Torrent{}
				t.Title = e.ChildText(".cellMainLink")
				t.Seeds = e.ChildText("td:nth-child(5)")
				t.Leechs = e.ChildText("td:nth-child(6)")
				t.Date = e.ChildText("td:nth-child(4)")
				t.Size = e.ChildText("td:nth-child(2)")
				t.Link = "https://kickasstorrents.to" + e.ChildAttr(".cellMainLink", "href")
				t.Magnet = e.ChildAttr("td:nth-child(3) a:nth-child(2)", "href")
				// go routine
				wg.Add(1)
				go func(t *models.Torrent) {
					defer wg.Done()
					c := colly.NewCollector()
					temp := t.Magnet
					c.OnHTML("body", func(e *colly.HTMLElement) {
						t.Magnet = e.ChildAttr("a.kaGiantButton", "href")
					})
					c.OnError(func(response *colly.Response, err error) {
						fmt.Printf("{ response : %d, error: %s}", response.StatusCode, err)
					})
					c.OnScraped(func(r *colly.Response) {
						if !strings.HasPrefix(t.Magnet, "magnet") {
							t.Magnet = temp
						}
						torrents = append(torrents, *t)
					})
					_ = c.Visit(t.Link)
				}(&t)
			})
		}
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Printf("{ response : %d, error: %s}", response.StatusCode, err)
	})

	c.OnScraped(func(response *colly.Response) {
		wg.Wait()
		if len(torrents) == 0 {
			res.Status(204)
		} else {
			_ = res.Status(200).JSON(torrents)
		}
	})

	_ = c.Visit(searchUrl)

	return nil
}

func GetSort(sort string) string {
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
func GetOrder(order string) string {
	if order == "0" {
		return "asc"
	} else {
		return "desc"
	}
}
