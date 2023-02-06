package misc

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"icodelife.com/itorrent-gofiber/routes/tgx"
)

func SyncAll(res *fiber.Ctx) {
	log.Info("Starting Sync Manually")
	tgx.GetRecentTorrents()
}
