package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"simpleserver/internal/controllers"
	"simpleserver/internal/sse"
	"time"
)

func main() {

	//cfg := config.MustLoad()

	//setup
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(cors.Default())
	/*
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
			ExposeHeaders:    []string{"Content-Type"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))*/

	//server-sent events
	sse.Init()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			sse.Instance.Broadcast("Periodic Event at " + time.Now().String())
		}
	}()

	router.GET("/updateLinesEvent", sse.Instance.UpdateLinesEventHandler)

	router.GET("/getLines", controllers.GetLines)
	router.POST("/setLine", controllers.SetLine)
	router.POST("/clearAllLines", controllers.ClearAllLines)

	err := router.Run("185.233.200.175:8080")
	if err != nil {
		log.Fatal(err)
	}
}
