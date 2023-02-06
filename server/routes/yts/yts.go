package yts

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/qopher/ytsgo"
	"icodelife.com/itorrent-gofiber/models"
)

func SearchYts(res *fiber.Ctx) error {
	queryStr := strings.ReplaceAll(res.Query("q"), " ", "+")
	sortStr := res.Query("s")
	orderStr := res.Query("o")
	yts, err := ytsgo.New()
	if err != nil {
		fmt.Print(err)
		return res.SendStatus(204)
	}
	id := ytsgo.LMSearch(queryStr)
	torrents, err := yts.ListMovies(id)
	if err != nil {
		fmt.Print(err)
		return res.SendStatus(204)
	}
	torrentData := make([]models.Torrent, 0)
	for _, torrent := range torrents.Movies {
		title := torrent.Title
		for _, tor := range torrent.Torrents {
			t := models.Torrent{}
			t.Title = title + " " + tor.Quality
			t.Size = tor.Size
			t.Magnet = tor.Magnet()
			t.Seeds = fmt.Sprint(tor.Seeds)
			t.Leechs = fmt.Sprint(tor.Peers)
			t.Date = tor.DateUploaded.String()
			torrentData = append(torrentData, t)
		}
	}

	if sortStr == "2" {
		sort.Slice(torrentData[:], func(i, j int) bool {
			size1, _ := strconv.Atoi(torrentData[i].Link)
			size2, _ := strconv.Atoi(torrentData[j].Link)
			if orderStr == "0" {
				return size1 < size2
			} else {
				return size1 > size2
			}
		})
	} else if sortStr == "1" {
		if orderStr == "0" {
			sort.Slice(torrentData[:], func(i, j int) bool {
				seed1, _ := strconv.Atoi(torrentData[i].Seeds)
				seed2, _ := strconv.Atoi(torrentData[j].Seeds)
				return seed1 < seed2
			})
		} else {
			sort.Slice(torrentData[:], func(i, j int) bool {
				seed1, _ := strconv.Atoi(torrentData[i].Seeds)
				seed2, _ := strconv.Atoi(torrentData[j].Seeds)
				return seed1 >= seed2
			})
		}
	}
	return res.Status(200).JSON(torrentData)
}
