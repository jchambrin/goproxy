package proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jchambrin/goproxy/pkg/config"
)

type CacheStorage interface {
	Get(key KeyCache) *CacheData
	Put(key KeyCache, data *CacheData)
}

type KeyCache struct {
	URI string
}

type CacheData struct {
	header http.Header
	body   []byte
}

type Proxy struct {
	protocol            string
	host                string
	port                int
	cacheEnabled        bool
	cacheAllowedMethods []string
}

func New(config config.Proxy) *Proxy {
	return &Proxy{
		protocol:            config.Destination.Protocol,
		host:                config.Destination.Host,
		port:                config.Destination.Port,
		cacheEnabled:        config.Cache.Enable,
		cacheAllowedMethods: config.Cache.AllowedMethods,
	}
}

func (p *Proxy) Start(w http.ResponseWriter, r *http.Request) {
	if p.cacheEnabled && containsString(r.Method, p.cacheAllowedMethods) {
		// TODO search into cache
	} else {
		data := p.httpProxy(r)
		writeResponse(w, data)
		return
	}
}

func (p *Proxy) httpProxy(r *http.Request) *CacheData {
	client := &http.Client{}
	req := r.Clone(r.Context())
	req.RequestURI = ""
	req.URL.Scheme = p.protocol
	req.URL.Host = fmt.Sprintf("%s:%d", p.host, p.port)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err) // TODO
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err) // TODO
	}
	return &CacheData{
		header: resp.Header,
		body:   body,
	}
}

func writeResponse(w http.ResponseWriter, data *CacheData) {
	copyHeaders(data.header, w)
	w.Write(data.body)
}

func copyHeaders(header http.Header, w http.ResponseWriter) {
	for key, value := range header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}
}
