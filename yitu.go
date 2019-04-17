package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/lzjluzijie/yitu/onedrive"

	"github.com/lzjluzijie/yitu/routers"
)

const VERSION = `v1.0.0-dev`

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("yitu %s by halulu", VERSION)

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"*"}
	router.Use(cors.New(corsConfig))

	routers.RegisterRouters(router)
	onedrive.LoadConfig()
	onedrive.Refresh()

	go func() {
		err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://t.halu.lu"+r.URL.String(), http.StatusMovedPermanently)
		}))
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	err := http.ListenAndServeTLS(":443", "cert", "key", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
