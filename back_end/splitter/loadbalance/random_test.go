/* ----------------------------------
*  @author suyame 2022-08-29 15:22:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package loadbalance

import "testing"

func TestNewRandom(t *testing.T) {
	_, err := NewRandom(servers)
	if err != nil {
		t.Error("NewRandom error, ", err)
	}
	_, err = NewRandom(servers, WithRandomWeights([]int{1}))
	if err != WeightsNotMatchErr {
		t.Error("NewRR(with weights) error, ", err)
	}
	_, err = NewRandom(servers, WithRandomWeights([]int{1, 1, 1, 1, 1, 1, 1, 1}))
	if err != WeightsNotMatchErr {
		t.Error("NewRR(with weights) error, ", err)
	}
	// 合法的
	_, err = NewRandom(servers, WithRandomWeights([]int{1, 5, 1, 2, 3, 1}))
	if err != nil {
		t.Error("NewRR(with weights) error, ", err)
	}
}

func TestRandom_Do(t *testing.T) {
	// 无权重的轮询
	rr, err := NewRandom(servers)
	if err != nil {
		t.Error("NewRandom error, ", err)
	}
	// 模拟20次请求
	for i := 0; i < 20; i++ {
		server, err := rr.Do()
		if err != nil {
			t.Error("Random Do error, ", err)
		}
		t.Log(server)
	}
}

func TestRandom_DoWithWeights(t *testing.T) {
	// 带权重的轮询
	rr, err := NewRandom(servers, WithRandomWeights([]int{1, 5, 1, 2, 3, 1}))
	if err != nil {
		t.Error("NewRandom(with weights) error, ", err)
	}
	// 模拟20次请求
	for i := 0; i < 20; i++ {
		server, err := rr.Do()
		if err != nil {
			t.Error("RandomDo error, ", err)
		}
		t.Log(server)
	}
}
