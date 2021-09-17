package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jchambrin/goproxy/pkg/config"
	"github.com/jchambrin/goproxy/pkg/proxy"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	proxyConf := config.Init(os.Getenv("CONFIGFILE_LOCATION"))
	p := proxy.New(proxyConf)
	r := mux.NewRouter()

	r.PathPrefix("/").HandlerFunc(p.Start)
	http.Handle("/", r)
	if err := http.ListenAndServe(proxyConf.Source, r); err != nil {
		log.Fatal(err)
	}
}
