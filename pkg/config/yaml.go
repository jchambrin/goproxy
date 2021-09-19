package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func loadFromYaml(p *Proxy, path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, p)
	return err
}
