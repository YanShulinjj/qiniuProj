/* ----------------------------------
*  @author suyame 2022-10-26 20:33:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

// Manager 管理所有websocket 信息
type Manager struct {
	ClientMap            map[string]*Client
	clientCount          uint
	Lock                 sync.Mutex
	Register, UnRegister chan *Client
	BroadCastMessage     chan *BroadCastMessageData
}

// 广播发送数据信息
type BroadCastMessageData struct {
	Id      string // 消息的标识符，标识指定用户
	Message []byte
}

// 向所有客户发送广播数据
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

			log.Println("广播数据：", data.Message)
		}
	}
}

// 启动 websocket 管理器
func (manager *Manager) Start() {
	log.Printf("websocket manage start")
	for {
		select {
		// 注册
		case client := <-manager.Register:
			log.Printf("client [%s] connect", client.Id)
			log.Printf("register client [%s]", client.Id)

			manager.Lock.Lock()
			manager.ClientMap[client.Id] = client
			manager.clientCount += 1
			manager.Lock.Unlock()

		// 注销
		case client := <-manager.UnRegister:
			log.Printf("unregister client [%s]", client.Id)
			manager.Lock.Lock()

			if _, ok := manager.ClientMap[client.Id]; ok {
				delete(manager.ClientMap, client.Id)
				manager.clientCount -= 1
			}

			manager.Lock.Unlock()
		}
	}
}

// 注册
func (manager *Manager) RegisterClient(client *Client) {
	manager.Register <- client
}

// 注销
func (manager *Manager) UnRegisterClient(client *Client) {
	manager.UnRegister <- client
}

// 当前连接个数
func (manager *Manager) LenClient() uint {
	return manager.clientCount
}

// 获取 wsManager 管理器信息
func (manager *Manager) Info() map[string]interface{} {
	managerInfo := make(map[string]interface{})
	managerInfo["clientLen"] = manager.LenClient()
	managerInfo["chanRegisterLen"] = len(manager.Register)
	managerInfo["chanUnregisterLen"] = len(manager.UnRegister)
	managerInfo["chanBroadCastMessageLen"] = len(manager.BroadCastMessage)
	return managerInfo
}

// 初始化 wsManager 管理器
var WebsocketManager = Manager{
	ClientMap:        make(map[string]*Client),
	Register:         make(chan *Client, 128),
	UnRegister:       make(chan *Client, 128),
	BroadCastMessage: make(chan *BroadCastMessageData, 128),
	clientCount:      0,
}

// gin 处理 websocket handler
func (manager *Manager) WsClient(ctx *gin.Context) { // 参数为 ctx *gin.Context 的即为 gin的路由绑定函数
	upGrader := websocket.Upgrader{
		// cross origin domain
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// 处理 Sec-WebSocket-Protocol Header
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},
	}

	// 生成uuid，作为sessionid
	// id := strings.ToUpper(strings.Join(strings.Split(uuid.NewV4().String(), "-"), ""))
	id := ctx.Request.RemoteAddr
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

	manager.RegisterClient(client) // 将client对象添加到管理器中
	go client.Read(manager)        // 从一个客户端读取数据
	go manager.WriteToAll()        // 将数据写入所有客户端
}
