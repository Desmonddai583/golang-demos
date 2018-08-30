package rpcsupport

import (
	"fmt"
	"golang-demos/crawler_distributed/config"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// ServeRPC receive rpc request
func ServeRPC(host string, service interface{}) error {
	rpc.Register(service)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}

// NewClient create a new client
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
