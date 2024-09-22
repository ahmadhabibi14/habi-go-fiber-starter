package helper

import (
	"myapi/internal/bootstrap/logger"
)

func Recover() {
	if rec := recover(); rec != nil {
		logger.Log.Panic().Msg(rec.(string))
	}
}
