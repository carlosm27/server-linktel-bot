package main

import (
  "net/http"
  

  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  
)
  


func setupRouter() *gin.Engine {
  r := gin.Default()

  config := cors.DefaultConfig()

  config.AllowOrigins = []string{"https://link-sharer.vercel.app/"}

  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })

  r.POST("/handler", Handler)
  r.POST("/link", LinkHandler)
 
  r.Use(cors.New(config))
  return r
}

func main() {
	
  r := setupRouter()
  r.Run(":8000")
}

