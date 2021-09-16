package main

import (
	"log"
	"net"
	"os"

	"github.com/jchambrin/goproxy/pkg/config"
	"github.com/jchambrin/goproxy/pkg/proxy"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	proxyConf := config.Init(os.Getenv("CONFIG_LOCATION"))
	process := proxy.New(proxyConf)

	l, err := net.Listen("tcp", proxyConf.Source)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go process.Start(conn)
	}
}
