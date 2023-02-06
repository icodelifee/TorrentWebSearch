package tpb

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"icodelife.com/itorrent-gofiber/models"
)

/*
	default /1/99/0
	size high = /1/5/0
	size low = 1/6/0
	latest = /1/3/0
	old = 1/4/0
	seed low 1/8/0
	seed high /1/7/0
*/

func getOrderSortStr(orderStr, sortStr string) string {
	switch sortStr {
	case "0":
		return "99"
	case "1":
		if orderStr == "0" {
			return "8"
		} else {
			return "7"
		}
	case "2":
		if orderStr == "0" {
			return "3"
		} else {
			return "4"
		}
	case "3":
		if orderStr == "0" {
			return "6"
		} else {
			return "5"

		}
	default:
		return "99"
	}
}
func GetTPB(res *fiber.Ctx) error {
	queryStr := strings.ReplaceAll(res.Query("q"), " ", "+")
	sortStr := res.Query("s")
	orderStr := res.Query("o")

	url := fmt.Sprintf("https://thepiratebay10.org/search/%s/1/%s/0", queryStr, getOrderSortStr(orderStr, sortStr))
	c := colly.NewCollector()
	var torrents = make([]models.Torrent, 0)
	c.OnHTML("#searchResult", func(e *colly.HTMLElement) {
		selector := e.DOM.Find("tbody > tr")
		if selector.Length() > 1 {
			e.ForEach("tr", func(i int, e *colly.HTMLElement) {
				if i == 0 || i == selector.Length() {
					return
				}
				t := models.Torrent{}
				e.DOM.Find("td").Each(func(i int, selection *goquery.Selection) {
					if i == 1 {
						t.Magnet = selection.Find("a").Get(1).Attr[0].Val
					} else if i == 2 {
						t.Seeds = selection.Text()
					} else if i == 3 {
						t.Leechs = selection.Text()
					}
				})
				t.Title = e.DOM.Find(".detLink").Text()
				splitString := strings.Split(e.DOM.Find(".detDesc").Text(), ",")
				if len(splitString) <= 1 {
					fmt.Print(t.Title)
				}
				t.Date = strings.ReplaceAll(strings.TrimSpace(splitString[0]), "Uploaded ", "")
				t.Size = strings.ReplaceAll(strings.TrimSpace(splitString[1]), "Size ", "")
				torrents = append(torrents, t)
			})
		}
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnScraped(func(r *colly.Response) {
		if len(torrents) > 0 {
			_ = res.Status(200).JSON(torrents)
		} else {
			res.Status(204)
		}
	})
	_ = c.Visit(url)

	return nil
}
