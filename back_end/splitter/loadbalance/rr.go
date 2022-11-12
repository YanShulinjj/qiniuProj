/* ----------------------------------
*  @author suyame 2022-08-29 9:36:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

// 实现负载均衡的常见算法
package loadbalance

// RR 实现轮询方式的负载均衡
// servers 表示所有的可选服务器
// loc 表示本次分配的服务器下标
type RR struct {
	servers []Server
	loc     int
	weights []int
	indexs  []Server
}

// 可选参数
type RROption func(*RR)

func WithRRWeights(weights []int) RROption {
	return func(rr *RR) {
		rr.weights = weights
		// 根据权重扩展index
		s := 0
		for i := range rr.weights {
			s += rr.weights[i]
		}
		rr.indexs = make([]Server, 0, s)
		for i := 0; i < len(rr.servers) && i < len(weights); i++ {
			for j := 0; j < weights[i]; j++ {
				rr.indexs = append(rr.indexs, rr.servers[i])
			}
		}
	}
}

func NewRR(servers []Server, options ...RROption) (*RR, error) {
	rr := &RR{
		servers: servers,
	}
	for _, op := range options {
		op(rr)
	}
	// 校验weight和server是否满足
	if rr.weights != nil && len(rr.servers) != len(rr.weights) {
		return nil, WeightsNotMatchErr
	}
	if rr.weights == nil {
		rr.indexs = rr.servers
		rr.weights = make([]int, len(rr.servers))
		for i := range rr.weights {
			rr.weights[i] = 1
		}
	}
	return rr, nil
}

func (rr *RR) Do() (Server, error) {
	loc := rr.loc % len(rr.indexs)
	cnt := 0
	for !rr.indexs[loc].IsAlive() {
		// 分配下一个
		// loc 跳转到下一个node的开始
		currentNode := rr.indexs[loc]
		for rr.indexs[loc] == currentNode {
			loc = (loc + 1) % len(rr.indexs)
		}
		cnt++
		if cnt >= len(rr.servers) {
			// 转了一圈都没有找到
			return nil, NoAliveServerErr
		}
	}
	// 开始分配
	rr.loc = loc + 1
	return rr.indexs[loc], nil
}
