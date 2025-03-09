package main

import "github.com/gin-gonic/gin"

func main(){
    route:=gin.Default()
	route.GET("/anurag", func(c *gin.Context){   // c is a pointer to gin.Context--->it represnt the context for the current HTTP request
		//c.Status(201)
		c.JSON(201,gin.H{   // gin.h is shortcut provided by Gin framework to create a map with key value pair
			"message":"Hello Anurag",
		})
	})
	route.Run(":8081")
}
