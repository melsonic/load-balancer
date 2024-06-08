package util

import (
	"math/rand"
	"time"

	"github.com/melsonic/load-balancer/server"
)

var (
	seed           *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	serverList                = []string{server.Server1Address, server.Server2Address, server.Server3Address, server.Server4Address}
	LBAddress      string     = ":3000"
	selectedServer int        = 0
)
