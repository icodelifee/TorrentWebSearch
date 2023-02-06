package snowfl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"icodelife.com/itorrent-gofiber/models"
)

var (
	baseUrl string = "https://snowfl.com/CtEtDPzcfBHDNVxjLqXuTCFNgtXAOJLWLzOQT/"
	letters        = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func GetSnowFl(res *fiber.Ctx) {
	queryStr := res.Query("q")
	sortStr := res.Query("s")
	orderStr := res.Query("o")
	epoc := strconv.Itoa(int(time.Now().Unix()))
	resp, err := http.Get(fmt.Sprintf("%s%s/%s/0/%s/NONE/1?_=%s", baseUrl, queryStr, randSeq(), getSort(sortStr, orderStr), epoc))
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		res.Status(204)
	}
	var torrents []models.Snowfl
	apiTorrents := make([]models.Torrent, 0)
	_ = json.Unmarshal(body, &torrents)

	if sortStr == "1" && orderStr == "0" {
		for i := len(torrents)/2 - 1; i >= 0; i-- {
			opp := len(torrents) - 1 - i
			torrents[i], torrents[opp] = torrents[opp], torrents[i]
		}
	} else if sortStr == "3" && orderStr == "0" {
		for i := len(torrents)/2 - 1; i >= 0; i-- {
			opp := len(torrents) - 1 - i
			torrents[i], torrents[opp] = torrents[opp], torrents[i]
		}
	}
	for _, tor := range torrents {
		t := models.Torrent{}
		t.Title = tor.Name
		t.Size = tor.Size
		t.Seeds = strconv.Itoa(int(tor.Seeder))
		t.Leechs = strconv.Itoa(int(tor.Leecher))
		t.Magnet = tor.Magnet
		t.Date = tor.Age
		t.Link = tor.URL
		apiTorrents = append(apiTorrents, t)
	}
	if len(apiTorrents) > 0 {
		_ = res.Status(200).JSON(apiTorrents)
	} else {
		res.Status(204)
	}
}

func getSort(sort, order string) string {
	if sort == "1" {
		return "SEED"
	} else if sort == "2" {
		if order == "0" {
			return "SIZE_ASC"
		} else {
			return "SIZE"
		}
	} else if sort == "3" {
		return "DATE"
	}
	return "NONE"
}
func randSeq() string {
	b := make([]rune, 7)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
