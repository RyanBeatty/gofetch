package gofetch

import (
	"net"
	"fmt"
)

type UPI interface {
	Connect(address, port string) error
	Close()
	
}

type Upi struct {
	Conn net.Conn
}

func (upi Upi) Connect(address string, port string) error{
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		return err
	}
	upi.Conn = conn
	return err
}

func (upi Upi) Close() {
	upi.Conn.Close()
}

