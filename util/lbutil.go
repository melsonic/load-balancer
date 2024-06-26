package util

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func CheckHealth() {
	t := time.Now()
	var activeServers []string
	for _, serverUrl := range serverList {
		response, responseErr := http.Get(serverUrl)
		if responseErr != nil || response.StatusCode != 200 {
			continue
		}
		activeServers = append(activeServers, serverUrl)
	}
	serverList = activeServers
	fmt.Printf("health check performed @ %d:%d:%d\n", t.Hour(), t.Minute(), t.Second())
}

func GetTargetURL() *url.URL {
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
	target := GetTargetURL()
	fmt.Printf("response from %s\n", target)
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}
