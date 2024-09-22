package controller

import (
	"errors"
	"os"
	"time"

	"myapi/configs"

	"myapi/internal/bootstrap/database"
	"myapi/internal/repository/sessions"

	"github.com/gofiber/fiber/v2"
)

func getSession(db *database.Database, c *fiber.Ctx) (session *sessions.Session, key string, err error) {
	sessionId := c.Cookies(configs.AUTH_COOKIE, ``)
	apiKey := c.Get(configs.HEADER_API_KEY, ``)

	var SKEY string = sessionId
	if sessionId == `` { SKEY = apiKey }
	
	if SKEY == `` {
		err = errors.New(`you are unauthorized to process this operation`)
		return
	}

	session = sessions.NewSessionMutator(db)
	
	err = session.GetSession(SKEY);
	if err != nil {
		c.ClearCookie(configs.AUTH_COOKIE)
		err = errors.New(`invalid session key`)
		return
	}

	key = SKEY

	return
}

func setCookie(c *fiber.Ctx, sessionKey string) {
	// 2 months expired
	expiration := time.Now().AddDate(0, 2, 0)

	c.Cookie(&fiber.Cookie{
		Name:     configs.AUTH_COOKIE,
		Value:    sessionKey,
		Expires:  expiration,
		SameSite: `Lax`,
		Secure:   os.Getenv(`PROJECT_ENV`) == `prod`,
		HTTPOnly: true,
	})
}

type DebugController struct {
	actionPrefix string
	db *database.Database
}

func NewDebugController(app *fiber.App, db *database.Database) {
	debugController := &DebugController{
		actionPrefix: `/debug`,
		db: db,
	}

	app.Get(debugController.actionPrefix, debugController.Debug)
}

func (_dc *DebugController) Debug(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Hello World !")
}