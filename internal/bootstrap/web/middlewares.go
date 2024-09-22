package web

import (
	"os"
	"strings"
	"time"

	"myapi/configs"

	"myapi/internal/bootstrap/database"

	"myapi/internal/bootstrap/logger"

	"myapi/internal/repository/sessions"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Middlewares struct {
	app *fiber.App
	db *database.Database
}

func NewMiddlewares(app *fiber.App, db *database.Database) *Middlewares {
	return &Middlewares{app, db}
}

func (_m *Middlewares) Init() {
	_m.RateLimiter()
	_m.CORS()
	_m.Logger()
	_m.Recover()
}

func (_m *Middlewares) RateLimiter() {
	_m.app.Use(limiter.New(limiter.Config{
		Max:        300,
		Expiration: 2 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			var errMessage string = "you have exceeded your rate limit, please try again a few moments later"

			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				`errors`: errMessage,
			})
		},
	}))
}

func (_m *Middlewares) CORS() {
	allowOrigins := []string{
		`http://localhost:5173`,
		`http://127.0.0.1:5173`,
	}

	allowHeaders := []string{
		fiber.HeaderOrigin,
		fiber.HeaderContentType,
		fiber.HeaderAccept,
		fiber.HeaderContentLength,
		fiber.HeaderAcceptLanguage,
		fiber.HeaderAcceptEncoding,
		fiber.HeaderConnection,
		fiber.HeaderAccessControlAllowOrigin,
		fiber.HeaderSetCookie,
		fiber.HeaderCookie,
	}
	
	_m.app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(allowOrigins, `, `),
		AllowCredentials: true,
		AllowHeaders:     strings.Join(allowHeaders, `, `),
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, `, `),
		MaxAge: 3600,
	}))
}

func (_m *Middlewares) Logger() {
	var conf fiberLogger.Config

	if os.Getenv("PROJECT_ENV") == `prod` {
		file, _ := os.OpenFile(
			configs.OS_PATH_WEBACCESS_LOG, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666,
		)
		conf = fiberLogger.Config{
			Format:        "{\"time\": \"${time}\", \"status\": \"${status}\", \"ip\": \"${ip}\", \"ips\": \"${ips}\", \"latency\": \"${latency}\", \"method\": \"${method}\", \"path\": \"${path}\", \"user_agent\": \"${user_agent}\", \"error\": \"${error}\"}\n",
			TimeFormat:    "2006-01-02T03:00:55+08:00",
			TimeZone:      "Asia/Makassar",
			Output:        file,
			DisableColors: true,
			CustomTags: map[string]fiberLogger.LogFunc{
				"user_agent": func(output fiberLogger.Buffer, c *fiber.Ctx, data *fiberLogger.Data, extraParam string) (int, error) {
					return output.WriteString(c.Get(fiber.HeaderUserAgent))
				},
			},
		}
	} else {
		conf = fiberLogger.Config{
			Format:     "${time} | ${status} | ${latency} | ${method} | ${path}\n",
			TimeFormat: "2006/01/02 03:04 PM",
			TimeZone:   "Asia/Makassar",
			Output:     os.Stdout,
		}
	}

	_m.app.Use(fiberLogger.New(conf))
}

func (_m *Middlewares) Recover() {
	_m.app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			logger.Log.Error().Str("path", c.Path()).Err(e.(error)).Msg("received unexpected panic error")
		},
	}))
}

// #################################################### //
// ############## OPTIONAL MIDDLEWARES ################ //
// #################################################### //

// function name use format: OPT_<MiddlewareName> for any optional middlewares

func (_m *Middlewares) OPT_Auth(c *fiber.Ctx) error {
	sessionId := c.Cookies(configs.AUTH_COOKIE, ``)
	apiKey := c.Get(configs.HEADER_API_KEY, ``)

	var SKEY string = sessionId
	if sessionId == `` { SKEY = apiKey }
	
	if SKEY == `` {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			`errors`: `you are unauthorized to process this operation`,
		})
	}

	session := sessions.NewSessionMutator(_m.db)
	
	err := session.GetSession(SKEY);
	if err != nil {
		c.ClearCookie(configs.AUTH_COOKIE)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			`errors`: `invalid session key`,
		})
	}

	return c.Next()
}

func (_m *Middlewares) OPT_WebSocket(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}

	return fiber.ErrUpgradeRequired
}