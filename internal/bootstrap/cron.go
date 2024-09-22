package bootstrap

import (
	"myapi/internal/bootstrap/database"

	"myapi/internal/bootstrap/logger"

	"github.com/robfig/cron"
)

type Cron struct {
	c 	*cron.Cron
	db	*database.Database
}

func NewCron(c *cron.Cron, db *database.Database) *Cron {
	return &Cron{c, db}
}

func (cr *Cron) DeleteNotification() {
	err := cr.c.AddFunc("@daily", func() {
		logger.Log.Info().Msg(`this function runs daily`)
	})
	
	if err != nil {
		logger.Log.Err(err).Msg(`failed to add cron job`)
	}
}