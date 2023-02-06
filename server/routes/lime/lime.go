package lime

import (
	"fmt"
	url2 "net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"icodelife.com/itorrent-gofiber/models"
)

func GetLimeTor(res *fiber.Ctx) error {
	queryStr := strings.ReplaceAll(res.Query("q"), " ", "+")
	sortStr := res.Query("s")
	orderStr := res.Query("o")

	url := fmt.Sprintf("https://www.limetorrents.info/search/all/%s/%s/1/", strings.TrimSpace(queryStr), getSort(sortStr))
	c := colly.NewCollector()
	torrents := make([]models.Torrent, 0)

	c.OnHTML("table.table2 tbody tr", func(element *colly.HTMLElement) {
		element.DOM.Each(func(i int, s *goquery.Selection) {
			t := models.Torrent{}
			t.Title = s.Find("td:nth-child(1)").Text()
			t.Seeds = s.Find("td:nth-child(4)").Text()
			t.Leechs = s.Find("td:nth-child(5)").Text()
			t.Date = strings.Split(s.Find("td:nth-child(2)").Text(), " - ")[0]
			t.Size = s.Find("td:nth-child(3)").Text()
			href, _ := s.Find("td.tdleft div.tt-name a:nth-child(2)").Attr("href")
			t.Link = "https://www.limetorrents.info" + href
			if len(t.Title) == 0 {
				return
			}

			// go routine go brrrrr
			str := s.Find("a.csprite_dl14").AttrOr("href", "")
			go func(t *models.Torrent, str string) {
				parsedUrl, _ := url2.Parse(str)
				infoHash := strings.Trim(parsedUrl.Path, "/torrent/.torrent")
				tracker := "&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2F9.rarbg.to%3A2740%2Fannounce&tr=udp%3A%2F%2F9.rarbg.me%3A2770%2Fannounce&tr=udp%3A%2F%2F9.rarbg.me%3A2730%2Fannounce&tr=udp%3A%2F%2F9.rarbg.me%3A2740%2Fannounce&tr=udp%3A%2F%2F9.rarbg.to%3A2720%2Fannounce&tr=udp%3A%2F%2F9.rarbg.to%3A2730%2Fannounce&tr=udp%3A%2F%2F9.rarbg.to%3A2710%2Fannounce&tr=udp%3A%2F%2F9.rarbg.to%3A2770%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.tiny-vps.com%3A6969%2Fannounce&tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&tr=udp%3A%2F%2Fretracker.lanta-net.ru%3A2710%2Fannounce&tr=udp%3A%2F%2Fipv4.tracker.harry.lu%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.cyberia.is%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce&tr=udp%3A%2F%2Fipv6.tracker.harry.lu%3A80%2Fannounce&tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.open-internet.nl%3A6969%2Fannounce&tr=udp%3A%2F%2Fopen.demonii.si%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.pirateparty.gr%3A6969%2Fannounce&tr=udp%3A%2F%2Fdenis.stalker.upeer.me%3A6969%2Fannounce&tr=udp%3A%2F%2Fp4p.arenabg.com%3A1337%2Fannounce"
				t.Magnet = "magnet:?xt=urn:btih:" + fmt.Sprintf("%s%s", infoHash, tracker)
				torrents = append(torrents, *t)
			}(&t, str)

		})
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Printf("{ response : %d, error: %s}", response.StatusCode, err)
	})
	c.OnScraped(func(response *colly.Response) {
		OrderTorrents(orderStr, sortStr, &torrents)
		if len(torrents) <= 0 {
			res.Status(204)
		} else {
			_ = res.Status(200).JSON(torrents)
		}
	})
	_ = c.Visit(url)

	return nil
}
func getSort(sort string) string {
	if sort == "1" {
		return "seeds"
	} else if sort == "2" {
		return "size"
	} else if sort == "3" {
		return "date"
	}
	return ""
}
func OrderTorrents(orderStr, sortStr string, torrents *[]models.Torrent) {
	if orderStr == "0" && sortStr != "0" {
		for i := len(*torrents)/2 - 1; i >= 0; i-- {
			opp := len(*torrents) - 1 - i
			(*torrents)[i], (*torrents)[opp] = (*torrents)[opp], (*torrents)[i]
		}
	}
}
