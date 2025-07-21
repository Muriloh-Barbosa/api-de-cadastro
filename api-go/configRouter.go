package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func newRouter() *gin.Engine {
	router := gin.Default()

	gin.DefaultWriter = io.MultiWriter(os.Stdout)

	// Middleware de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Defina seus endpoints aqui
	router.GET("/user", userHandler())
	router.POST("/create/user", userHandler())
	router.PUT("/update/user", userHandler())

	return router
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
