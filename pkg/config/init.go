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
	AllowedMethods []string `yaml:"allowedMethods"`
}

var (
	defaultAllowedMethods = []string{"GET", "HEAD"}
)

func Init(path string) Proxy {
	res, err := loadFromYaml(path)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
