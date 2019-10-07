package loadbalance

import (
	"math/rand"
	"net"

	".."
)

func NewRandom(proxies []proxyclient.Dial) proxyclient.Dial {
	return func(network, address string) (net.Conn, error) {
		dial := proxies[rand.Intn(len(proxies))]
		return dial(network, address)
	}
}
