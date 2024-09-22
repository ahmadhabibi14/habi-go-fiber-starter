package configs

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

func ConnectClickhouse() (*sqlx.DB, error) {
	clickhouseDbName := os.Getenv("CLICKHOUSE_DB")
	clickhouseUser := os.Getenv("CLICKHOUSE_USER")
	clickhousePassword := os.Getenv("CLICKHOUSE_PASSWORD")
	clickhouseIsDebug := "true"
	if os.Getenv("PROJECT_ENV") == "prod" {
		clickhouseIsDebug = "false"
	}

	clickhouseURL := fmt.Sprintf(
		"tcp://127.0.0.1:9000?username=%s&password=%s&database=%s&debug=%s",
		clickhouseUser, clickhousePassword, clickhouseDbName, clickhouseIsDebug,
	)

	driverName := "clickhouse"
	db, err := sql.Open(driverName, clickhouseURL)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	zlogLevel, err := strconv.Atoi(os.Getenv(`LOG_LEVEL`))
	if err != nil {
		zlogLevel = int(zerolog.InfoLevel)
	}

	var output io.Writer = zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: `2006/01/02 03:04 PM`,
		PartsOrder: []string{
			zerolog.TimestampFieldName,
			zerolog.LevelFieldName,
			zerolog.CallerFieldName,
			zerolog.MessageFieldName,
		},
	}

	zlog := zerolog.New(output).Level(zerolog.Level(zlogLevel)).With().Timestamp().Logger()
	
	db = sqldblogger.OpenDriver(clickhouseURL, db.Driver(), zerologadapter.New(zlog))

	sqlxDb := sqlx.NewDb(db, driverName)

	err = sqlxDb.Ping()
	if err != nil {
		sqlxDb.Close()
		return nil, err
	}

	return sqlxDb, nil
}