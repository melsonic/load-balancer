package util

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func GetTargetURL(serverList []string) *url.URL {
	availableServers := serverList
  serverIndex := selectedServer % len(serverList)
	target, urlParseErr := url.Parse(availableServers[serverIndex])
	selectedServer = selectedServer + 1
	if urlParseErr != nil {
		fmt.Printf("error %s\n", urlParseErr.Error())
	}
	return target
}

func LBHandler(w http.ResponseWriter, r *http.Request) {
	target := GetTargetURL(serverList)
	fmt.Printf("response from %s\n", target)
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}
