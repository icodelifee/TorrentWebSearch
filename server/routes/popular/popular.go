package popular

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"icodelife.com/itorrent-gofiber/internal"
	"icodelife.com/itorrent-gofiber/models"
)

func GetPopularStreaming() []models.TmdbData {

	c := colly.NewCollector()
	list := make([]models.TmdbData, 0)
	c.OnHTML("h2", func(e *colly.HTMLElement) {
		href, _ := e.DOM.Find("a").Attr("href")
		list = append(list, models.TmdbData{
			Title: e.Text,
			Link:  href,
		})
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.Visit("https://www.themoviedb.org/remote/panel?panel=popular_scroller&group=streaming")
	c.Wait()
	return list
}

func GetPopularDetails(res *fiber.Ctx) {

	list := GetPopularStreaming()
	popularList := make([]models.TMDBData, 0)
	for _, v := range list {
		imdb := getImdbDetails(v.Link, !strings.Contains(v.Link, "movie"))
		popularList = append(popularList, imdb)
	}
	res.JSON(popularList)
}

func getImdbDetails(link string, isSeries bool) models.TMDBData {

	imdbId := fmt.Sprintf(`https://api.themoviedb.org/3/%s?api_key=%s&language=en-US&append_to_response=external_ids`, link, internal.MyApiToken)
	res, err := http.Get(imdbId)

	if err != nil {
		fmt.Print(err.Error())
	}

	defer res.Body.Close()

	if !isSeries {
		var movie models.TMDBMovie
		json.NewDecoder(res.Body).Decode(&movie)
		genres := make([]string, 0)
		for _, v := range movie.Genres {
			genres = append(genres, v.Name)
		}
		return models.TMDBData{
			BackdropPath: "https://image.tmdb.org/t/p/w500" + movie.BackdropPath,
			Title:        movie.Title,
			Overview:     movie.Overview,
			PosterPath:   "https://image.tmdb.org/t/p/w500" + movie.PosterPath,
			ReleaseDate:  movie.ReleaseDate,
			VoteAverage:  movie.VoteAverage,
			Genres:       genres,
			ID:           movie.ID,
			Popularity:   movie.Popularity,
			Runtime:      movie.Runtime,
			Tagline:      movie.Tagline,
			IMDBID:       movie.ExternalIDS.ImdbID,
		}
	} else {
		var series models.TMDBSeries
		json.NewDecoder(res.Body).Decode(&series)
		genres := make([]string, 0)
		for _, v := range series.Genres {
			genres = append(genres, v.Name)
		}
		var totalRuntime int64 = 0
		if len(series.EpisodeRunTime) > 0 {
			totalRuntime = series.EpisodeRunTime[len(series.EpisodeRunTime)-1]
		}
		return models.TMDBData{
			BackdropPath: "https://image.tmdb.org/t/p/w500" + series.BackdropPath,
			Title:        series.Name,
			Overview:     series.Overview,
			PosterPath:   "https://image.tmdb.org/t/p/w500" + series.PosterPath,
			ReleaseDate:  series.FirstAirDate,
			VoteAverage:  series.VoteAverage,
			Genres:       genres,
			ID:           series.ID,
			Popularity:   series.Popularity,
			Runtime:      totalRuntime,
			Tagline:      series.Tagline,
			IMDBID:       series.ExternalIDS.ImdbID,
		}
	}
}
