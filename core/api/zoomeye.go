package api

import (
	"drdos/config"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
)

func Zoomeye(port int, page int) ([]string, error) {
	iplist := []string{}
	req, _ := http.NewRequest("GET", "https://api.zoomeye.org/host/search?query=port:"+strconv.Itoa(port)+"&page="+strconv.Itoa(page), nil)
	req.Header.Set("API-KEY", config.ZoomeyeApi)
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println(err)
		return iplist, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	value := gjson.Get(string(body), "matches.#.ip").Array()
	for _, every := range value {
		iplist = append(iplist, every.String())
	}
	return iplist, nil
}
