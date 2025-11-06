package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

type Server struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`
}
type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Data struct {
	Path string `json:"path"`
}
type Config struct {
	Server Server `json:"server"`
	Auth   Auth   `json:"auth"`
	Data   Data   `json:"data"`
}

var C = Config{}

func Init() {
	configFilePath := flag.String("c", "config.json", "the config file path")
	flag.Parse()
	b, err := os.ReadFile(*configFilePath)
	if err != nil {
		log.Panicln(err)
	}
	err = json.Unmarshal(b, &C)
	if err != nil {
		log.Panicln(err)
	}
}
