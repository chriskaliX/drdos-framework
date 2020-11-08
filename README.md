# Drdos Framework

[README-EN](https://github.com/chriskaliX/drdos-framework/blob/master/README-EN.md)

Drdos Framework是我学习反射型ddos的产出，我尝试将它设计为一个真正意义上的框架。这是它运行的图片。

![image](https://github.com/chriskaliX/drdos-framework/blob/master/imgs/demo.png)

## 郑重申明

**该框架仅用于学习，禁止用于其他任何非法用途**

## 介绍

Drdos框架是一个既可以用来校验IP是否存在drdos漏洞，又可以用来进行一些攻击测试的框架

## 环境准备

1. 一台在公网的Linux服务器
2. Golang环境

## 使用

1. `go get github.com/google/gopacket`
2. 运行`go run main.go --help`来查看帮助

check模式

```shell
go run main.go -m c -s xx.xx.xx.xx -type dns -api -o test.txt # 使用api查询，需要在config.go中进行修改
go run main.go -m c -s xx.xx.xx.xx -type dns -range xx.xx.xx.xx/24 -o test.txt # 扫描指定网段
go run main.go -m c -s xx.xx.xx.xx -type dns -f input.txt -o test.txt # 从文件中获取IP地址
```

attack模式

```shell
go run main.go -m a -f xxx -type dns -t xx.xx.xx.xx -p xx  # 要注意如果攻击没有效果，可能是前面有NAT
```

mix模式(觉得没有必要了，后期会逐步移除掉)

[*] 注意 : 在check和mix模式下，-o输出的文件在`/data/results/`目录下.

### 一些帮助

1. `--type` 应该为 `dns`,`mem`,`ntp`,`snmp`,`ssdp`,`portmap`,`ldap` 中的一种
2. 如果在阿里云上使用，记得在安全组上允许UDP，并把IP(即`-s`选项)设为`eth0`的网卡IP(而不是公网IP)

### 配置

默认的配置文件在 `config/config.go` ，下面是默认的配置

```go
package config

const (
	ListenPort     = 50000 // Check ip列表的时候的监听端口
	Threshold      = 100   // 当包的大小大于阈值的时候，计数接受
	WaitTime       = 10    // 全部发包完毕后，等待其余数据包的时间
	Blacklists     = "/data/blacklists/blacklists"
	MaxAtktime     = 300 // 最大攻击时间
	AttackInterval = 0
	ShodanApi      = ""
	ShodanPage     = 10 // 默认搜索页数，10页=1000个
)
```

## 支持的协议

|Port|Protocol|
|:-:|:-:|
|53|dns|
|111|portmap|
|123|ntp|
|161|snmp|
|389|ldap|
|1900|ssdp|
|11211|memcache|

## 更新计划

1. HTTP API(考虑用Gin，直接模板渲染做个简单的WEB - UI)
2. ~~SHODAN~~| FOFA API
3. ~~Blacklist of attack~~
4. Improve response check
5. Support more protocol
6. Use sqlite to save data
