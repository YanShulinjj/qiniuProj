/* ----------------------------------
*  @author suyame 2022-11-01 20:25:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync/atomic"
)

var seq int64

// gin 处理 websocket handler
func WSHandler(ctx *gin.Context) { // 参数为 ctx *gin.Context 的即为 gin的路由绑定函数

	pageName := ctx.Query("page")

	upGrader := websocket.Upgrader{
		// cross origin domain
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// 处理 Sec-WebSocket-Protocol Header
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},
	}

	// 生成uuid，作为sessionid
	// id := fmt.Sprintf("%s_%d", ctx.Request.RemoteAddr, seq)
	// atomic.AddInt64(&seq, 1)
	var id string
	if IPMode {
		id = ctx.ClientIP()
	} else {
		id = fmt.Sprintf("%s_%d", ctx.ClientIP(), seq)
		atomic.AddInt64(&seq, 1)
	}

	// 设置http头部，添加sessionid
	heq := make(http.Header)
	heq.Set("sessionid", id)

	// 建立一个websocket的连接
	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, heq)
	if err != nil {
		log.Printf("websocket connect error: %s", id)
		return
	}

	// 创建一个client对象（包装websocket连接）
	client := &Client{
		Id:   id,
		Conn: conn,
	}

	// 根据pageName 获取指定manager
	m, ok := defaultMangerGroup.Get(pageName)
	if !ok {
		m = NewManager(pageName, id)
		// 启动m
		go m.Start()
		// 第一次进入该页面的为所有者
		defaultMangerGroup.Put(pageName, m)
	}
	m.RegisterClient(client) // 将client对象添加到管理器中
	go client.Read(m)        // 从一个客户端读取数据
	go m.WriteToAll()        // 将数据写入所有客户端
}
