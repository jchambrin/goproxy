package proxy

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/jchambrin/goproxy/pkg/config"
)

type Proxy struct {
	destination string
}

func New(config config.Proxy) *Proxy {
	return &Proxy{
		destination: config.Destination,
	}
}

func (p *Proxy) Start(conn net.Conn) {
	defer conn.Close()

	remote, err := net.Dial("tcp", p.destination)
	if err != nil {
		log.Println(err)
		return
	}
	defer remote.Close()

	go io.Copy(remote, conn)
	io.Copy(conn, remote)
}
