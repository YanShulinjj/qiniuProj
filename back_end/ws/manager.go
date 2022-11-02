/* ----------------------------------
*  @author suyame 2022-10-26 20:33:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

// Manager 管理所有websocket 信息
type Manager struct {
	sync.Mutex
	PageName             string
	AuthorId             string
	ReadOnly             bool
	ClientMap            map[string]*Client
	clientCount          uint
	Register, UnRegister chan *Client
	BroadCastMessage     chan *BroadCastMessageData
}

func NewManager(pageName, authorId string) *Manager {
	// 初始化 wsManager 管理器
	return &Manager{
		PageName:         pageName,
		AuthorId:         authorId,
		ClientMap:        make(map[string]*Client),
		Register:         make(chan *Client, MaxClientNum),
		UnRegister:       make(chan *Client, MaxClientNum),
		BroadCastMessage: make(chan *BroadCastMessageData, MaxMessageNum),
		clientCount:      0,
	}
}

// WriteToAll 向所有客户发送广播数据
func (m *Manager) WriteToAll() {
	for {
		select {
		case data, ok := <-m.BroadCastMessage:
			if !ok {
				log.Println("没有取到广播数据。")
			}
			for _, client := range m.ClientMap {
				sender, ok := m.ClientMap[data.Id]

				// 绘图数据不会发给自己，如果这里是将绘图数据写给客户端，应该跳过正在绘图的人
				if sender.Id == client.Id {
					continue
				}

				if !ok {
					log.Println("用户不存在") // 这里应该是存在的，先判断一下
				}

				client.Lock()
				client.Conn.WriteMessage(websocket.TextMessage, data.Message)
				client.Unlock()
			}
			if DEBUG {
				log.Println("广播数据：", data.Message)
			}
		}
	}
}

// Start 启动 websocket 管理器
func (m *Manager) Start() {
	log.Printf("websocket manage start")
	for {
		select {
		// 注册
		case client := <-m.Register:
			log.Printf("client [%s] connect", client.Id)
			log.Printf("register client [%s]", client.Id)

			m.Lock()
			m.ClientMap[client.Id] = client
			m.clientCount += 1
			m.Unlock()

		// 注销
		case client := <-m.UnRegister:
			log.Printf("unregister client [%s]", client.Id)
			m.Lock()

			if _, ok := m.ClientMap[client.Id]; ok {
				delete(m.ClientMap, client.Id)
				m.clientCount -= 1
			}

			m.Unlock()
		}
	}
}

// RegisterClient 注册
func (m *Manager) RegisterClient(c *Client) {
	m.Register <- c
}

// UnRegisterClient 注销
func (m *Manager) UnRegisterClient(c *Client) {
	m.UnRegister <- c
}

// GetClientNum 获取当前连接个数
func (m *Manager) GetClientNum() uint {
	return m.clientCount
}

// Info 获取 wsManager 管理器信息
func (m *Manager) Info() map[string]interface{} {
	managerInfo := make(map[string]interface{})
	managerInfo["clientLen"] = m.GetClientNum()
	managerInfo["chanRegisterLen"] = len(m.Register)
	managerInfo["chanUnregisterLen"] = len(m.UnRegister)
	managerInfo["chanBroadCastMessageLen"] = len(m.BroadCastMessage)
	return managerInfo
}
