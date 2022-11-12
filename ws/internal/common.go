/* ----------------------------------
*  @author suyame 2022-11-01 19:59:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package internal

const (
	ReadWriteAllow = iota
	ReadOnly
	ModeChangeType = 12
	NeedSyncType   = 13
)

var DEBUG = false
var IPMode = false
var MaxClientNum = 128
var MaxMessageNum = 128

type Message struct {
	Type int         `json:"type"`
	Attr interface{} `json:"Attr"`
	Mode int         `json:"Mode"`
}

// 广播发送数据信息
type BroadCastMessageData struct {
	Id      string // 消息的标识符，标识指定用户
	Message []byte
}
