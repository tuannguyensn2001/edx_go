package socket

import (
	app_ctx "edx_go/component"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"log"
	"sync"
)

type engine struct {
	server *socketio.Server
}

var instance *engine
var mutex sync.Mutex

func NewEngine() *engine {
	mutex.Lock()

	defer mutex.Unlock()

	if instance == nil {
		instance = &engine{}
	}

	return instance
}

func (e *engine) Run(appCtx app_ctx.AppContext, r *gin.Engine) error {

	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	e.server = server

	server.OnConnect("/", func(conn socketio.Conn) error {
		log.Println("connected", conn)
		return nil
	})

	server.OnError("", func(conn socketio.Conn, err error) {
		log.Println("error", err)
	})

	go func() {
		err := e.server.Serve()
		if err != nil {

		}
	}()

	r.GET("/socket.io/*any", gin.WrapH(e.server))
	r.POST("/socket.io/*any", gin.WrapH(e.server))

	return nil
}
