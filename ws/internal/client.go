/* ----------------------------------
*  @author suyame 2022-10-27 19:13:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package internal

import (
	"github.com/gorilla/websocket"
	"k8s.io/apimachinery/pkg/util/json"
	"log"
	"sync"
)

// Client 客户端信息
type Client struct {
	sync.Mutex                 // 并发安全
	Id         string          // 用户标识
	Conn       *websocket.Conn // 用户连接
}

// Read 从websocket连接读取数据
func (c *Client) Read(m *Manager) {
	defer func() {
		m.UnRegister <- c
		if DEBUG {
			log.Printf("[%s]client [%s] disconnect\n", m.PageName, c.Id)
		}
		if err := c.Conn.Close(); err != nil {
			if DEBUG {
				log.Printf("[%s]client [%s] disconnect err: %s\n", m.PageName, c.Id, err)
			}
		}
	}()

	for {
		messageType, message, err := c.Conn.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			// 读取完毕
			if DEBUG {
				log.Printf("[%s]%s\n", m.PageName, err)
			}
			break
		}
		if DEBUG {
			log.Printf("[%s]client [%s] receive message: %s\n", m.PageName, c.Id, string(message))
		}
		var msg Message
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Printf("[%s]%s\n", m.PageName, err)
		}
		if DEBUG {
			log.Printf("[%s] client: %s 接收到消息：%s\n",
				m.PageName, c.Id, message)
		}
		// 如果收到调整 读写模式
		if c.Id == m.AuthorId && msg.Type == ModeChangeType {
			m.ReadOnly = msg.Attr.(bool)
		}
		// 如果当前客户端是所有者
		// 或则是可读可写模式
		// 允许将非所有者的操作广播
		if c.Id == m.AuthorId || !m.ReadOnly {
			// 向广播消息写入数据
			m.BroadCastMessage <- &BroadCastMessageData{Id: c.Id, Message: message}
		}
	}
}
