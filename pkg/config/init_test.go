package config

import (
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg := Init("./test/config.yaml")
	if cfg.Cache.Enable == false {
		t.Errorf("cache should be disabled: got %v want %v", cfg.Cache.Enable, false)
	}
	expected := []string{"GET", "HEAD"}
	if !reflect.DeepEqual(cfg.Cache.AllowedMethods, expected) {
		t.Errorf("allowedMethods dont match: got %v want %v", cfg.Cache.AllowedMethods, expected)
	}
}
