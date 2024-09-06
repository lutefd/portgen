package port

import (
	"fmt"
	"math/rand"
	"net"
)

func Generate(min, max int) int {
	for {
		port := rand.Intn(max-min+1) + min
		if !isInUse(port) {
			return port
		}
	}
}

func isInUse(port int) bool {
	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return true
	}
	listener.Close()
	return false
}
