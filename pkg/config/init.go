package config

import "log"

type Proxy struct {
	Name        string
	Source      string
	Destination string
}

func Init(path string) Proxy {
	res, err := loadFromYaml(path)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
