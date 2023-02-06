package main

import (
	"icodelife.com/itorrent-gofiber/routes/tgx"
	"icodelife.com/itorrent-gofiber/routes/yts"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	log "github.com/sirupsen/logrus"
	_337x "icodelife.com/itorrent-gofiber/routes/1337x"
	"icodelife.com/itorrent-gofiber/routes/kat"
	"icodelife.com/itorrent-gofiber/routes/lime"
	"icodelife.com/itorrent-gofiber/routes/misc"
	"icodelife.com/itorrent-gofiber/routes/rarbg"
	"icodelife.com/itorrent-gofiber/routes/tpb"
)

// func setupRoutes(app *fiber.App) {

// 	// app.Get("/popular", popular.GetPopularDetails)
// 	// app.Get("/popular-data", popular.GetPopularTorrent)

// 	// app.Get("/syncall", misc.SyncAll)
// 	// app.Static("/recents", "./recents.json")
// 	// app.Static("/movies", "./movies.json")
// 	// app.Static("/series", "./series.json")
// 	// app.Static("/anime", "./anime.json")
// }

func main() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	app := fiber.New()

	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     "*",
			AllowCredentials: true,
			AllowMethods:     "GET",
			AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept",
		}),
	)

	app.Get("/", misc.Alive)
	app.Get("/1337x", _337x.Get1337x)
	app.Get("/kat", kat.GetKat)
	app.Get("/lime", lime.GetLimeTor)
	app.Get("/rarbg", rarbg.GetRARBG)
	app.Get("/yts", yts.SearchYts)
	app.Get("/tpb", tpb.GetTPB)
	app.Get("/tgx", tgx.GetTGX)

	_ = app.Listen(":5050")
}
