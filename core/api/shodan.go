package api

import (
	"context"
	"drdos/config"
	"fmt"
	"net/http"

	"github.com/shadowscatcher/shodan"
	"github.com/shadowscatcher/shodan/search"
)

func Shodan(port int, page uint) ([]string, error) {
	iplist := []string{}
	portSearch := search.Params{
		Page: page,
		Query: search.Query{
			Port: port,
		},
	}
	client, _ := shodan.GetClient(config.ShodanApi, http.DefaultClient, true)
	ctx := context.Background()
	result, err := client.Search(ctx, portSearch)
	if err != nil {
		fmt.Println(err)
		return iplist, err
	}
	for _, match := range result.Matches {
		iplist = append(iplist, match.IPString())
	}
	return iplist, nil
}
