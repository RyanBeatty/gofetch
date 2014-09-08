package gofetch

import (
	"net"
	"fmt"
)

func Connect(address string) net.Conn {
	address += ":21"
	conn, err := net.Dial("tcp", address)
	RaiseOnError(err, fmt.Sprintf("error: can't connect to <%s>", address))
	return conn
}