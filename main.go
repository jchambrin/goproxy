package main

import (
	"flag"
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
	memoryCache := cache.NewMemoryCache(time.Duration(proxyConf.Cache.TTL))
	p := proxy.New(proxyConf, memoryCache)

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(p.Handle)
	http.Handle("/", r)
	if err := http.ListenAndServe(proxyConf.Source, r); err != nil {
		log.Fatal(err)
	}
}
