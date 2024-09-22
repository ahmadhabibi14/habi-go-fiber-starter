package sessions

import (
	"errors"
	"time"

	"myapi/internal/bootstrap/database"
	"myapi/internal/bootstrap/logger"

	"github.com/goccy/go-json"
)

const (
	SESSION_PREFIX = `session:`
	// 2 months expiration for user session
	SESSION_EXPIRED = (24 * time.Hour) * 60
)

type Session struct {
	Db						*database.Database `json:"-"` 
	
	UserID        uint64				`json:"user_id"`
	Username      string				`json:"username"`
	Email					string 				`json:"email"`
	Role					string				`json:"role"`
}

func NewSessionMutator(Db *database.Database) *Session {
	return &Session{Db: Db}
}

func (_s *Session) SetSession(key string) error {
	sessionJSON, err := json.Marshal(&_s)
	if err != nil {
		return errors.New(`failed to marshal session data`)
	}

	sessionKey := SESSION_PREFIX + key
	err = _s.Db.RD.Set(sessionKey, sessionJSON, SESSION_EXPIRED).Err()
	if err != nil {
		errMsg := `failed to set session`
		logger.Log.Err(err).Msg(errMsg)
		return errors.New(errMsg)
	}

	return nil
}

func (_s *Session) UpdateSession(key string) error {
	sessionKey := SESSION_PREFIX + key

	sessionData, err := _s.Db.RD.Get(sessionKey).Result()
	if err != nil {
		return errors.New(`session not found`)
	}

	var sess Session
	err = json.Unmarshal([]byte(sessionData), sess)
	if err != nil {
		return errors.New(`invalid session data`)
	}

	_s.Email = sess.Email
	_s.UserID = sess.UserID
	_s.Username = sess.Username

	sessionJSON, err := json.Marshal(&_s)
	if err != nil {
		return errors.New(`failed to marshal session data`)
	}
	
	err = _s.Db.RD.Set(sessionKey, sessionJSON, SESSION_EXPIRED).Err()
	if err != nil {
		errMsg := `failed to update session`
		logger.Log.Err(err).Msg(errMsg)
		return errors.New(errMsg)
	}

	return nil
}

func (_s *Session) GetSession(key string) error {
	sessionKey := SESSION_PREFIX + key
	sessionData, err := _s.Db.RD.Get(sessionKey).Result()
	if err != nil {
		return errors.New(`session not found`)
	}

	err = json.Unmarshal([]byte(sessionData), _s)
	if err != nil {
		return errors.New(`invalid session data`)
	}

	return nil
}

func (_s *Session) DeleteSession(key string) error {
	sessionKey := SESSION_PREFIX + key
	err := _s.Db.RD.Del(sessionKey).Err()
	if err != nil {
		errMsg := `failed to delete session`
		logger.Log.Err(err).Msg(errMsg)
		return errors.New(errMsg)
	}

	return nil
}