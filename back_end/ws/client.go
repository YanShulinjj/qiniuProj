/* ----------------------------------
*  @author suyame 2022-10-27 19:13:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

// Client 单个 websocket 信息
type Client struct {
	sync.Mutex                 // 加一把锁
	Id         string          // 用户标识
	Conn       *websocket.Conn // 用户连接
}

// Read 从websocket连接读取数据
func (c *Client) Read(manager *Manager) {
	defer func() {
		WebsocketManager.UnRegister <- c
		log.Printf("client [%s] disconnect", c.Id)
		if err := c.Conn.Close(); err != nil {
			log.Printf("client [%s] disconnect err: %s", c.Id, err)
		}
	}()

	for {
		messageType, message, err := c.Conn.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			log.Println(err)
			break
		}
		log.Printf("client [%s] receive message: %s", c.Id, string(message))

		// 向广播消息写入数据
		manager.BroadCastMessage <- &BroadCastMessageData{Id: c.Id, Message: message}
	}
}
