package app_ctx

import (
	"edx_go/websocket"
	"gorm.io/gorm"
)

type AppContext interface {
	GetDBConnection() *gorm.DB
	GetHub() *websocket.Hub
}

type appCtx struct {
	db  *gorm.DB
	Hub *websocket.Hub
}

func NewAppContext(db *gorm.DB, hub *websocket.Hub) *appCtx {
	return &appCtx{
		db:  db,
		Hub: hub,
	}
}

func (ctx *appCtx) GetDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetHub() *websocket.Hub {
	return ctx.Hub
}
