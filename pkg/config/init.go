package config

import "log"

type Proxy struct {
	Name        string
	Source      string
	Destination DestinationConfig
	Cache       CacheConfig
}

type DestinationConfig struct {
	Protocol string
	Host     string
	Port     int
}

type CacheConfig struct {
	Enable         bool
	TTL            string   `yaml:"ttl"`
	AllowedMethods []string `yaml:"allowedMethods"`
}

var (
	defaultConfig = Proxy{
		Cache: CacheConfig{
			Enable:         false,
			AllowedMethods: []string{"GET", "HEAD"},
		},
	}
)

func Init(path string) Proxy {
	res := defaultConfig
	if err := loadFromYaml(&res, path); err != nil {
		log.Fatal(err)
	}

	return res
}
