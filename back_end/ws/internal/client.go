/* ----------------------------------
*  @author suyame 2022-10-27 19:13:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package internal

import (
	"fmt"
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
		log.Printf("client [%s] disconnect", c.Id)
		if err := c.Conn.Close(); err != nil {
			log.Printf("client [%s] disconnect err: %s", c.Id, err)
		}
	}()

	for {
		messageType, message, err := c.Conn.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			// 读取完毕
			if DEBUG {
				log.Println(err)
			}
			break
		}
		if DEBUG {
			log.Printf("client [%s] receive message: %s", c.Id, string(message))
		}
		// 反序列划
		var msg Message
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(c.Id, " 接收到消息：", msg)
		// 如果收到调整 读写模式
		if msg.Type == ModeChangeType {
			m.ReadOnly = msg.Attr.(bool)
			continue
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
