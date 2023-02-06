package rarbg

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/araddon/dateparse"
	"github.com/gofiber/fiber/v2"
	"github.com/qopher/go-torrentapi"
	"icodelife.com/itorrent-gofiber/models"
)

func GetRARBG(res *fiber.Ctx) error {
	queryStr := strings.ReplaceAll(res.Query("q"), " ", "+")
	sortStr := res.Query("s")
	orderStr := res.Query("o")
	api, err := torrentapi.New("itorrent")
	if err != nil {
		return nil
	}
	api.SearchString(queryStr)
	if sortStr == "1" {
		api.Ranked(true).Format("json_extended").Limit(100).Sort("seeders")
	} else if sortStr == "3" {
		api.Ranked(true).Format("json_extended").Limit(100).Sort("last")
	} else {
		api.Ranked(true).Format("json_extended").Limit(100)
	}

	results, err := api.Search()
	if err != nil {
		fmt.Print(err)
		res.Status(204)
	}
	torrents := make([]models.Torrent, 0)

	for _, res := range results {
		t := models.Torrent{}
		t.Title = res.Title
		t.Size = ByteCountDecimal(res.Size)
		t.Seeds = strconv.Itoa(res.Seeders)
		t.Leechs = strconv.Itoa(res.Leechers)
		t.Magnet = res.Download
		t.Link = strconv.Itoa(int(res.Size))
		dt, err := dateparse.ParseAny(res.PubDate)
		if err != nil {
			t.Date = res.PubDate
		}
		t.Date = dt.Format("02 Jan 06 15:04")
		torrents = append(torrents, t)
	}
	SortOrderTorrents(sortStr, orderStr, &torrents)
	if len(torrents) == 0 {
		return res.SendStatus(400)
	} else {
		return res.Status(200).JSON(torrents)

	}
}
func ByteCountDecimal(b uint64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func SortOrderTorrents(sortStr, orderStr string, torrents *[]models.Torrent) {
	if sortStr == "2" {
		sort.Slice((*torrents)[:], func(i, j int) bool {
			size1, _ := strconv.Atoi((*torrents)[i].Link)
			size2, _ := strconv.Atoi((*torrents)[j].Link)
			if orderStr == "0" {
				return size1 < size2
			} else {
				return size1 > size2
			}
		})
	} else if sortStr == "1" {
		if orderStr == "0" {
			sort.Slice((*torrents)[:], func(i, j int) bool {
				seed1, _ := strconv.Atoi((*torrents)[i].Seeds)
				seed2, _ := strconv.Atoi((*torrents)[j].Seeds)
				return seed1 < seed2
			})
		}
	} else if sortStr == "3" {
		if orderStr == "0" {
			for i := len(*torrents)/2 - 1; i >= 0; i-- {
				opp := len(*torrents) - 1 - i
				(*torrents)[i], (*torrents)[opp] = (*torrents)[opp], (*torrents)[i]
			}
		}
	}
}
