package tgx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

var baseUrl = "https://torrentgalaxy.to/"

func GetRecentTorrents() {
	log.Info("Starting Sync")
	c := colly.NewCollector()
	syncRecents(c)
	syncMovies(c)
	syncSeries(c)

	fmt.Print("Sync Completed")
}

func syncRecents(c *colly.Collector) {
	recents := getRecentLink(c)
	recentHotpicks := getTorrentData(recents)
	if len(recentHotpicks) > 0 {
		file, _ := json.Marshal(recentHotpicks)
		_ = ioutil.WriteFile("recents.json", file, 0644)
	}
}

func syncMovies(c *colly.Collector) {
	movies := getMoviesLinks(c)
	moviesHotpicks := getTorrentData(movies)
	if len(moviesHotpicks) > 0 {
		file, _ := json.Marshal(moviesHotpicks)
		_ = ioutil.WriteFile("movies.json", file, 0644)
	}
}

func syncSeries(c *colly.Collector) {
	series := getSeriesLinks(c)
	seriesHotpicks := getTorrentData(series)
	if len(seriesHotpicks) > 0 {
		file, _ := json.Marshal(seriesHotpicks)
		_ = ioutil.WriteFile("series.json", file, 0644)
	}
}

func getRecentLink(c *colly.Collector) []string {
	var links []string
	c.OnHTML(".slidingDivb-fda7642e5643f7828c51f696dfed2d419f9459ec", func(element *colly.HTMLElement) {
		element.DOM.Find("td.hotlist").Each(func(i int, selection *goquery.Selection) {
			href, _ := selection.Find("a").Attr("href")
			links = append(links, baseUrl+href)
		})
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	_ = c.Visit(baseUrl + "torrents.php")
	c.Wait()
	return links
}

func getMoviesLinks(c *colly.Collector) []string {
	var links []string
	c.OnHTML("div[class=\"panel-body slidingDivf-857ed435e3d21023cf3e1061d94565c0af77262f\"]", func(element *colly.HTMLElement) {
		element.DOM.Find("div.hotpicks").Each(func(i int, selection *goquery.Selection) {
			href, _ := selection.Find("a").Attr("href")
			links = append(links, baseUrl+href)
		})
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	_ = c.Visit(baseUrl + "torrents-hotpicks.php?cat=1")
	c.Wait()
	return links
}

func getSeriesLinks(c *colly.Collector) []string {
	var links []string
	c.OnHTML("div[class=\"panel-body slidingDivf-562c62355fb439e4daa35217d6700033e9d1abdc\"]", func(element *colly.HTMLElement) {
		element.DOM.Find("div.hotpicks").Each(func(i int, selection *goquery.Selection) {
			href, _ := selection.Find("a").Attr("href")
			links = append(links, baseUrl+href)
		})
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	_ = c.Visit(baseUrl + "torrents-hotpicks.php?cat=3")
	c.Wait()
	return links
}

func getAnimeLinks(c *colly.Collector) []string {
	var links []string
	c.OnHTML("ul", func(element *colly.HTMLElement) {
		element.DOM.Find("li").Each(func(i int, selection *goquery.Selection) {
			href, _ := selection.Find("a").Attr("href")
			links = append(links, "https://horriblesubs.info"+href)
		})
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	_ = c.Visit("https://horriblesubs.info/api.php?method=getlatest")
	c.Wait()
	return links
}
