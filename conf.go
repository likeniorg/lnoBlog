package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var Key string
var Domain string
var ServerIP string
var DatabaseIP string
var DatabasePort string
var DBuser string
var DBpass string

// 获取配置
func getConf() {
	conf := Conf{}
	data, err := os.ReadFile(".conf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Key = conf.Key
	Domain = conf.Domain
	ServerIP = conf.ServerIP
	DatabaseIP = conf.DatabaseIP
	DatabasePort = conf.DatabasePort
	DBuser = conf.DBuser
	DBpass = conf.DBpass
}
