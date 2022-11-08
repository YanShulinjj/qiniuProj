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
	sync.RWMutex
	PageName             string
	AuthorId             string
	ReadOnly             bool
	ClientMap            map[string]*Client
	clientCount          uint
	Register, UnRegister chan *Client
	BroadCastMessage     chan *BroadCastMessageData
	LastMessage          *BroadCastMessageData
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
				log.Printf("[%s]没有取到广播数据。\n", m.PageName)
			}
			// 使用多个协程广播消息
			wg := sync.WaitGroup{}
			sender, ok := m.ClientMap[data.Id]
			for _, client := range m.ClientMap {
				wg.Add(1)
				go func(c *Client) {
					defer wg.Done()
					if !ok || sender.Id != c.Id {
						c.Lock()
						c.Conn.WriteMessage(websocket.TextMessage, data.Message)
						c.Unlock()
					}
				}(client)
			}
			wg.Wait()

			if DEBUG {
				log.Printf("[%s]广播数据: %v\n", m.PageName, data.Message)
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

func (m *Manager) IsExist(username string) bool {
	m.RLock()
	defer m.RUnlock()
	_, ok := m.ClientMap[username]
	return ok
}
