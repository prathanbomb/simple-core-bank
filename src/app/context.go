package app

import (
	"github.com/oatsaysai/simple-core-bank/src/db"
	log "github.com/oatsaysai/simple-core-bank/src/logger"
)

type Context struct {
	Logger    log.Logger
	Config    *Config
	DB        db.DB
	RequestID string
}

func (app *App) NewContext() *Context {
	return &Context{
		Logger: app.Logger,
		Config: app.Config,
		DB:     app.DB,
	}
}

func (ctx *Context) WithLogger(logger log.Logger) *Context {
	ret := *ctx
	ret.Logger = logger
	return &ret
}

func (ctx *Context) getLogger() log.Logger {
	return ctx.Logger.WithFields(log.Fields{
		"package":    "app",
		"request_id": ctx.RequestID,
	})
}
