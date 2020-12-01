package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

var (
	Typemap map[string]int
	Dir     string
)

func init() {
	Typemap = map[string]int{
		"ntp":       123,
		"ssdp":      1900,
		"memcached": 11211,
		"snmp":      161,
		"portmap":   111,
		"dns":       53,
		"cldap":     389,
	}
	Dir, _ = os.Getwd()
}

// PathExists 判断路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// IPCheck 判断是否为IPv4路径
func IPCheck(ip string) bool {
	regex := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regex, strings.Trim(ip, " ")); match {
		return true
	}
	return false
}

// ProcessBar 进度条
func ProcessBar(now int, all int) {
	str := "[Check] [" + bar(int((float64(now)/float64(all))*20), 20) + "] " + strconv.Itoa(now) + "/" + strconv.Itoa(all)
	fmt.Printf("\r%s", str)
	if now == all {
		fmt.Println()
	}
}

func bar(count, size int) string {
	str := ""
	for i := 0; i < size; i++ {
		if i < count {
			str += "="
		} else {
			str += " "
		}
	}
	return str
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 动态调用
func Call(m map[string]interface{}, name string, params ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("the number of input params not match!")
	}
	in := make([]reflect.Value, len(params))
	for k, v := range params {
		in[k] = reflect.ValueOf(v)
	}
	return f.Call(in), nil
}

/*
	Host函数，输入cidr返回iplist
	有个问题，应该是通用问题，我都是把所有IP存到一个list里，即在内存中
	所以太大了会出现溢出的问题，后期尝试类似python迭代器的方式实现吧
*/
func Hosts(cidr string) ([]string, error) {
	inc := func(ip net.IP) {
		for j := len(ip) - 1; j >= 0; j-- {
			ip[j]++
			if ip[j] > 0 {
				break
			}
		}
	}
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	return ips[1 : len(ips)-1], nil
}

// 抄来的代码
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

// 抄来的
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func FileNameCheck(filename string) (string, error) {
	var (
		result string
		err    error
	)
	result = filepath.Base(filename)
	if result == "." || result == "\\" || result == "/" {
		err = errors.New("filename error!")
		return "", err
	}
	return result, nil
}

func ColorPrint(input string, level string) {
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	switch level {
	case "info":
		fmt.Printf("%s%s\n", blue("[*] "), input)
		log.Info(input)
	case "warn":
		fmt.Printf("%s%s\n", yellow("[!] "), input)
		log.Warn(input)
	case "err":
		fmt.Printf("%s%s\n", red("[-] "), input)
		log.Error(input)
	case "success":
		fmt.Printf("%s%s\n", green("[+] "), input)
		log.Info(input)
	}
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}
