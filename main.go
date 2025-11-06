package main

import (
	"fmt"
	"log"
	"net/http"
	"small_imgbed/config"
	"small_imgbed/logger"
	"small_imgbed/router"
)

func main() {
	logger.Init()
	config.Init()
	r := router.SetUpRouter()
	addrAndPort := fmt.Sprintf("%s:%d", config.C.Server.Addr, config.C.Server.Port)
	log.Printf("The server is listening on %s", addrAndPort)
	err := http.ListenAndServe(addrAndPort, r)
	if err != nil {
		log.Panicln(err)
	}
}
