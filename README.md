# Drdos Framework

Drdos Framework is my outcome of drdos working. I try to write this like a real framework.

## Declaration

**This tool is for learning only. Not for illegal use.**

## Prepare

1. Linux server
2. Golang env

## Usage

1. `go get github.com/google/gopacket`
2. Just run `go run main.go --help`

`[*] Attention : In check mode and mix mode, -f in /data/loadfile/, -o in /data/results/. But in attack mode, -f in /data/results/. You know, just for convenience`

### Some help

`--type` should be one of `dns`,`mem`,`ntp`,`snmp`,`ssdp`,`portmap`,`ldap`
`-f` must in `/data/loadfile/`
`-o` the outputfile is saved in `/data/results/`

## Q&A

Q: How much flow can this do?
A: It depends on how many source ip you get, and the type of attack. Memcached nearly 200 times, ntp 10~15 times, ssdp 10 times, dns 5 times...

## Intruduction

Drdos Framework can check the vulner of drdos iplist. Also it's got the attack mode to start drdos attack.

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
