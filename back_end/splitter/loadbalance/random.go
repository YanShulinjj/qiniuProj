/* ----------------------------------
*  @author suyame 2022-08-29 15:03:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package loadbalance

import (
	"math/rand"
	"time"
)

type Random struct {
	servers []Server
	weights []int
	indexs  []Server
}

type RandomOption func(*Random)

func WithRandomWeights(weights []int) RandomOption {
	return func(rd *Random) {
		rd.weights = weights
		// 根据权重扩展index
		s := 0
		for i := range rd.weights {
			s += rd.weights[i]
		}
		rd.indexs = make([]Server, 0, s)
		for i := 0; i < len(rd.servers) && i < len(weights); i++ {
			for j := 0; j < weights[i]; j++ {
				rd.indexs = append(rd.indexs, rd.servers[i])
			}
		}
	}
}

func NewRandom(servers []Server, options ...RandomOption) (*Random, error) {
	rd := &Random{
		servers: servers,
	}
	for _, op := range options {
		op(rd)
	}
	// 校验weight和server是否满足
	if rd.weights != nil && len(rd.servers) != len(rd.weights) {
		return nil, WeightsNotMatchErr
	}
	if rd.weights == nil {
		rd.indexs = rd.servers
		rd.weights = make([]int, len(rd.servers))
		for i := range rd.weights {
			rd.weights[i] = 1
		}
	}
	// 固定随机，保证概率相等
	rand.Seed(time.Now().UnixNano())
	return rd, nil
}

func (rd *Random) Do() (Server, error) {
	loc := rand.Intn(len(rd.indexs))
	cnt := 0
	for !rd.indexs[loc].IsAlive() {
		loc = rand.Intn(len(rd.indexs))
		cnt++
		if cnt == len(rd.indexs) {
			// 判断是否还有alive节点
			for _, node := range rd.servers {
				if node.IsAlive() {
					return node, nil
				}
			}
			// 全部宕机
			return nil, NoAliveServerErr
		}
	}
	return rd.indexs[loc], nil
}
