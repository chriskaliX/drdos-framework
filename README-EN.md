# Drdos Framework

Drdos Framework is my outcome of drdos working. I try to write this like a real framework. Here is the pic of the demo.

![image](https://github.com/chriskaliX/drdos-framework/blob/master/imgs/demo.png)

## Declaration

**This tool is for learning only. Not for illegal use.**

## Intruduction

Drdos Framework can check the vulner of drdos iplist. Also it's got the attack mode to start drdos attack.

## Prepare

1. Linux server
2. Golang env

## Usage

1. `go get github.com/google/gopacket`
2. Just run `go run main.go --help`

[*] Attention : In check mode and mix mode, -o in /data/results/.

### Some help

`--type` should be one of `dns`,`mem`,`ntp`,`snmp`,`ssdp`,`portmap`,`ldap`
`-o` the outputfile is saved in `/data/results/`

### Configuration

The configuration file is in `config/config.go`，and here is the default

```go
package config

const (
	ListenPort     = 50000 // Check ip列表的时候的监听端口
	Threshold      = 100   // 当包的大小大于阈值的时候，计数接受
	WaitTime       = 10    // 全部发包完毕后，等待其余数据包的时间
	Blacklists     = "/data/blacklists/blacklists"
	MaxAtktime     = 300 // 最大攻击时间
	AttackInterval = 0
)
```

## Q&A

Q: How much flow can this do?
A: It depends on how many source ip you get, and the type of attack. Memcached nearly 200 times, ntp 10~15 times, ssdp 10 times, dns 5 times...

## Protocol Supported

|Port|Protocol|
|:-:|:-:|
|53|dns|
|111|portmap|
|123|ntp|
|161|snmp|
|389|ldap|
|1900|ssdp|
|11211|memcache|

## Update plan

It will be deleted once I achieve

1. HTTP API
2. SHODAN | FOFA API
3. ~~Blacklist of attack~~
4. Improve response check
5. Support more protocol
6. Use sqlite to save data

## New

See, this work is kind of interestring. In fact, the most important part of this program is the vulnerable ip address. Once you get the ip address you just get them all. Since I've alreday done this work on my own, I may upload this one later...
