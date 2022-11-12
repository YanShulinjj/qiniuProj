/* ----------------------------------
*  @author suyame 2022-08-29 11:27:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package loadbalance

type Node struct {
	name  string
	alive bool
}

func (n *Node) IsAlive() bool {
	return n.alive
}

// 定义一组servers
var servers = []Server{
	&Node{
		"node1",
		true,
	},
	&Node{
		"node2",
		true,
	},
	&Node{
		"node3",
		true,
	},
	&Node{
		"node4",
		true,
	},
	&Node{
		"node5",
		true,
	},
	&Node{
		"node6",
		false,
	},
}
