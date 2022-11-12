/* ----------------------------------
*  @author suyame 2022-08-29 10:04:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package loadbalance

import "testing"

func TestNewRR(t *testing.T) {
	_, err := NewRR(servers)
	if err != nil {
		t.Error("NewRR error, ", err)
	}
	_, err = NewRR(servers, WithRRWeights([]int{1}))
	if err != WeightsNotMatchErr {
		t.Error("NewRR(with weights) error, ", err)
	}
	_, err = NewRR(servers, WithRRWeights([]int{1, 1, 1, 1, 1, 1, 1, 1}))
	if err != WeightsNotMatchErr {
		t.Error("NewRR(with weights) error, ", err)
	}
	// 合法的
	_, err = NewRR(servers, WithRRWeights([]int{1, 5, 1, 2, 3, 1}))
	if err != nil {
		t.Error("NewRR(with weights) error, ", err)
	}
}

func TestRR_Do(t *testing.T) {
	// 无权重的轮询
	rr, err := NewRR(servers)
	if err != nil {
		t.Error("NewRR error, ", err)
	}
	// 模拟20次请求
	for i := 0; i < 20; i++ {
		server, err := rr.Do()
		if err != nil {
			t.Error("RRDo error, ", err)
		}
		t.Log(server)
	}
}

func TestRR_DoWithWeights(t *testing.T) {
	// 带权重的轮询
	rr, err := NewRR(servers, WithRRWeights([]int{1, 5, 1, 2, 3, 1}))
	if err != nil {
		t.Error("NewRR(with weights) error, ", err)
	}
	// 模拟20次请求
	for i := 0; i < 20; i++ {
		server, err := rr.Do()
		if err != nil {
			t.Error("RRDo error, ", err)
		}
		t.Log(server)
	}
}
