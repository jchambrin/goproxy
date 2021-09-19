package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jchambrin/goproxy/pkg/cache"
	"github.com/jchambrin/goproxy/pkg/config"
	"github.com/jchambrin/goproxy/pkg/proxy"
)

func main() {
	configLocation := flag.String("config", ".", "configuration file")
	flag.Parse()

	proxyConf := config.Init(*configLocation)
	TTL, err := time.ParseDuration(proxyConf.Cache.TTL)
	if err != nil {
		fmt.Println(err)
		TTL = 60 * time.Second
	}
	memoryCache := cache.NewMemoryCache(TTL)
	p := proxy.New(proxy.Params{
		Protocol:            proxyConf.Destination.Protocol,
		Host:                proxyConf.Destination.Host,
		Port:                proxyConf.Destination.Port,
		CacheEnabled:        proxyConf.Cache.Enable,
		CacheAllowedMethods: proxyConf.Cache.AllowedMethods,
	}, memoryCache)

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(p.Handle)
	http.Handle("/", r)
	if err := http.ListenAndServe(proxyConf.Source, r); err != nil {
		log.Fatal(err)
	}
}
