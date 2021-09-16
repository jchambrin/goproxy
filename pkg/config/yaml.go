package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func loadFromYaml(path string) (Proxy, error) {
	res := Proxy{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return res, err
	}
	err = yaml.Unmarshal(data, &res)
	return res, err
}
