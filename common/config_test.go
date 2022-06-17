package common

import (
	"encoding/json"
	"testing"
)

// "root:Passw0rd@tcp(127.0.0.1:3306)/microdemo?charset=utf8mb4&parseTime=True&loc=Local"
func TestGetMysqlConfig(t *testing.T) {
	cfg := &MysqlConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "Passw0rd",
		Database: "microdemo",
	}

	b, _ := json.Marshal(cfg)
	t.Log(string(b))
}
