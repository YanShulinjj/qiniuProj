/* ----------------------------------
*  @author suyame 2022-11-01 20:25:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"k8s.io/apimachinery/pkg/util/json"
	"log"
	"net/http"
	"qiniu/pkg/xerr"
)

var seq int64

// gin 处理 websocket handler
func WSHandler(c *gin.Context) {
	userName := c.Query("username")
	authorName := c.Query("author")
	pageName := c.Query("page")

	// 根据pageName 获取指定manager
	key := authorName + "#" + pageName
	m, ok := defaultMangerGroup.Get(key)

	if !ok {
		// 第一次进入该页面的为所有者
		m = NewManager(key, authorName)
		// 启动m
		go m.Start()
		defaultMangerGroup.Put(key, m)
	}
	// 如果username的客户端已经连接, 拒绝本次连接
	if m.IsExist(userName) {
		c.JSON(http.StatusBadGateway, gin.H{
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
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
	}

	// 设置http头部，添加sessionid
	heq := make(http.Header)
	heq.Set("sessionid", userName)

	// 建立一个websocket的连接
	conn, err := upGrader.Upgrade(c.Writer, c.Request, heq)
	if err != nil {
		log.Printf("websocket connect error: %s", userName)
		return
	}

	// 创建一个client对象（包装websocket连接）
	client := &Client{
		Id:   userName,
		Conn: conn,
	}

	m.RegisterClient(client) // 将client对象添加到管理器中
	go client.Read(m)        // 从一个客户端读取数据
	go m.WriteToAll()        // 将数据写入所有客户端
}

// 同步数据的ws连接
func SyncHandler(c *gin.Context) {
	username := c.Query("username")
	authorName := c.Query("author")
	pageName := c.Query("page")

	// 根据pageName 获取指定manager
	key := authorName + "#" + pageName
	m, ok := syncMangerGroup.Get(key)
	if !ok {
		// 第一次进入该页面的为所有者
		m = NewManager(key, authorName)
		// 启动m
		go m.Start()
		syncMangerGroup.Put(key, m)
	}
	// 如果username的客户端已经连接, 拒绝本次连接
	if m.IsExist(username) {
		c.JSON(http.StatusBadGateway, gin.H{
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
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
	}

	// 设置http头部，添加sessionid
	heq := make(http.Header)
	heq.Set("sessionid", username)

	// 建立一个websocket的连接
	conn, err := upGrader.Upgrade(c.Writer, c.Request, heq)
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
	if client.Id == m.AuthorId {
		go m.WriteToAll() // 将数据写入所有客户端
	} else {
		// 有新client连接
		//
		syncMsg := Message{
			Type: NeedSyncType,
		}
		message, err := json.Marshal(syncMsg)
		if err != nil {
			log.Fatalf("websocket connect error: %s", username)
		}
		m.BroadCastMessage <- &BroadCastMessageData{Id: client.Id, Message: message}
	}
	go client.Read(m) // 从一个客户端读取数据

}
