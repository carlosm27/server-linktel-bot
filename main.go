package main

import (
  "net/http"
  

  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  
)
  


func setupRouter() *gin.Engine {
  r := gin.Default()

  config := cors.Config{
    AllowOrigins: []string{"https://link-sharer.vercel.app"},
    AllowMethods: []string{"POST"},
    AllowHeaders: []string{"Origin", "Content-Type", "Access-Control-Allow-Origin"},
  }
  
  r.Use(cors.New(config))	
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })

  r.POST("/handler", Handler)
  r.POST("/link", LinkHandler)
 
  
  return r
}

func main() {
	
  r := setupRouter()
  r.Run()
}

