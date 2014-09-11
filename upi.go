package gofetch

import (
	"net"
	"fmt"
)


// Defines the interface by which the UPI module
// of the FTP protocol defined by RFC 959 
// (http://www.w3.org/Protocols/rfc959/2_Overview.html)
type UPI interface {

	// Takes an address string and attempts to setup a TCP
	// connection with the specified address.
	//
	// PARAMETERS **NOT FINALIZED**:
	//	- address: address string of host we wish to connect to
	//	- port: port string we wish to connect to
	//
	// RETURNS:
	//	- error: if there is an error making the connection
	//
	Connect(address, port string) error

	// Closes the FTP connection
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

