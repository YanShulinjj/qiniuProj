/* ----------------------------------
*  @author suyame 2022-11-01 20:25:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"qiniu/pkg/xerr"
)

var seq int64

// gin 处理 websocket handler
func WSHandler(ctx *gin.Context) {
	username := ctx.Query("username")
	pageName := ctx.Query("page")

	// 根据pageName 获取指定manager
	m, ok := defaultMangerGroup.Get(pageName)
	if !ok {
		// 第一次进入该页面的为所有者
		m = NewManager(pageName, username)
		// 启动m
		go m.Start()
		defaultMangerGroup.Put(pageName, m)
	}
	// 如果username的客户端已经连接, 拒绝本次连接
	if m.IsExist(username) {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status_code": xerr.ClientExistedErr,
		})
		return
	}
	upGrader := websocket.Upgrader{
		// cross origin domain
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// 处理 Sec-WebSocket-Protocol Header
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},
	}

	// 设置http头部，添加sessionid
	heq := make(http.Header)
	heq.Set("sessionid", username)

	// 建立一个websocket的连接
	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, heq)
	if err != nil {
		log.Printf("websocket connect error: %s", username)
		return
	}

	// 创建一个client对象（包装websocket连接）
	client := &Client{
		Id:   username,
		Conn: conn,
	}

	m.RegisterClient(client) // 将client对象添加到管理器中
	go client.Read(m)        // 从一个客户端读取数据
	go m.WriteToAll()        // 将数据写入所有客户端
}
