package loadbalance

import (
	"hash/crc32"
	"net"

	"github.com/missdeer/ProxyClient"
)

func NewHash(proxies []proxyclient.Dial) proxyclient.Dial {
	return func(network, address string) (net.Conn, error) {
		checksum := crc32.ChecksumIEEE([]byte(address))
		dial := proxies[int(checksum) % len(proxies)]
		return dial(network, address)
	}
}
