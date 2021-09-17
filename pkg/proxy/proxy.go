package proxy

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jchambrin/goproxy/pkg/config"
)

var (
	errorData = &CacheData{statusCode: http.StatusInternalServerError}
)

type CacheStorage interface {
	Get(key KeyCache) (*CacheData, bool)
	Put(key KeyCache, data *CacheData)
}

type KeyCache struct {
	URI string
}

type CacheData struct {
	statusCode int
	header     http.Header
	body       []byte
}

type Proxy struct {
	cache               CacheStorage
	protocol            string
	host                string
	port                int
	cacheEnabled        bool
	cacheAllowedMethods []string
}

func New(config config.Proxy, cache CacheStorage) *Proxy {
	return &Proxy{
		cache:               cache,
		protocol:            config.Destination.Protocol,
		host:                config.Destination.Host,
		port:                config.Destination.Port,
		cacheEnabled:        config.Cache.Enable,
		cacheAllowedMethods: config.Cache.AllowedMethods,
	}
}

// Handle reverse proxy
func (p *Proxy) Handle(w http.ResponseWriter, r *http.Request) {
	if p.cacheEnabled && containsString(r.Method, p.cacheAllowedMethods) {
		key := KeyCache{r.RequestURI}
		if resp, ok := p.cache.Get(key); ok {
			writeResponse(w, resp)
			return
		}
		data := p.httpProxy(r)
		if data.statusCode >= 200 && data.statusCode <= 299 {
			p.cache.Put(key, data)
		}
		writeResponse(w, data)
		return
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
		return errorData
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errorData
	}
	return &CacheData{
		statusCode: resp.StatusCode,
		header:     resp.Header,
		body:       body,
	}
}

func writeResponse(w http.ResponseWriter, data *CacheData) {
	if data == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	copyHeaders(data.header, w)
	w.WriteHeader(data.statusCode)
	w.Write(data.body)
}

func copyHeaders(header http.Header, w http.ResponseWriter) {
	for key, value := range header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}
}
