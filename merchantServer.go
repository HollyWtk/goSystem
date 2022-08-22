package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"merchantService/config"
	"merchantService/server/web"
	"merchantService/utils"
	"net/http"
	"time"
)

func main() {

	host := config.File.MustValue("web_server", "host", "127.0.0.1")
	port := config.File.MustValue("web_server", "port", "8088")

	router := gin.Default()
	//路由
	web.Init(router)
	s := &http.Server{
		Addr:           host + ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	println(utils.Award(1, "yhh"))
	_, _, err2 := utils.ParseToken("1eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVaWQiOjEsIlVzZXJuYW1lIjoieWhoIiwiZXhwIjoxNjYxNzM5ODM1LCJpYXQiOjE2NjExMzUwMzV9.V0dw6NjYwXtbLeTRbPrbedlT1Mt3vg_aOxuOVyeLG7U ")
	if err2 != nil {
		println(err2)
	}
	err := s.ListenAndServe()
	log.Println(err)
}
