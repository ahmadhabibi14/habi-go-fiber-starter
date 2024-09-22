package bootstrap

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"myapi/configs"

	"myapi/internal/bootstrap/database"
	"myapi/internal/repository/users"

	"myapi/internal/bootstrap/logger"

	"myapi/internal/bootstrap/web"

	"myapi/internal/controller"

	"myapi/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron"
)

type App struct {
	// Common
	httpServer 					*fiber.App
	db         					*database.Database
	cronjob    					*Cron
	// Services
	authService 				*service.AuthService
	userService 				*service.UserService
}

func NewApp() *App {
	return &App{}
}

func (_a *App) Run() {
	waits := make(chan int)

	_a.setupEnv()
	_a.setupLogger()
	_a.setupDatabases()
	_a.setupServices()
	_a.setupHTTP()
	_a.setupCron()
	go _a.shutdown()

	go func() {
		port := ":" + os.Getenv("WEB_PORT")
		if port == ":" {
			port = ":3000"
		}

		if err := _a.httpServer.Listen(port); err != nil {
			logger.Log.Err(err).Msg("failed to start http server")
			_a.closeServices()
			os.Exit(1)
		}
	}()

	<-waits
}

func (_a *App) setupHTTP() {
	app := web.NewWebserver()
	middleware := web.NewMiddlewares(app, _a.db)
	middleware.Init()

	controller.NewDebugController(app, _a.db)
	controller.NewAuthController(app, _a.authService)
	controller.NewUserController(app, middleware, _a.userService)

	// Serve public static files
	app.Static("/", configs.OS_PATH_STATIC_FILES, fiber.Static{
		Download:      false,
		CacheDuration: 20 * time.Second,
		Browse:        false,
	})

	_a.httpServer = app
}

func (_a *App) setupServices() {
	userRepo := users.NewUserImpl_ClickHouse(_a.db)

	_a.authService = service.NewAuthService(_a.db, &userRepo)
	_a.userService = service.NewUserService(_a.db, &userRepo)
}

func (_a *App) setupDatabases() {
	pq, err := configs.ConnectClickhouse()
	if err != nil {
		logger.Log.Panic().Str("error", err.Error()).Msg("failed to connect clickhouse")
	}

	rd := configs.NewRedisClient()
	_, err = rd.Ping().Result()
	if err != nil {
		logger.Log.Panic().Str("error", err.Error()).Msg("failed to connect redis")
	}

	db := database.NewDatabase(pq, rd)
	_a.db = db
}

func (_a *App) setupEnv() {
	configs.LoadEnv()
}

func (_a *App) setupLogger() {
	logger.InitLogger()
}

func (_a *App) shutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s

		logger.Log.Info().Msg("shutting down...")
		_a.closeServices()
		_a.cronjob.c.Stop()

		os.Exit(0)
	}()
}

func (_a *App) setupCron() {
	c := cron.New()
	_a.cronjob = NewCron(c, _a.db)
	_a.cronjob.c.Start()
}

func (_a *App) closeServices() {
	if err := _a.httpServer.Shutdown(); err != nil {
		logger.Log.Err(err).Msg("failed to shutdown [httpserver]")
	} else {
		logger.Log.Info().Msg("cleaned up [httpserver]")
	}

	if err := _a.db.DB.Close(); err != nil {
		logger.Log.Err(err).Msg("failed to shutdown [postgresql]")
	} else {
		logger.Log.Info().Msg("cleaned up [postgresql]")
	}

	if err := _a.db.RD.Close(); err != nil {
		logger.Log.Err(err).Msg("failed to shutdown [redis]")
	} else {
		logger.Log.Info().Msg("cleaned up [redis]")
	}
}
