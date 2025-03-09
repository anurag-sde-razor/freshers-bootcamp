package main

import "github.com/gin-gonic/gin"

func main(){
	route:=gin.Default()
	route.POST("/posting",PostData){

	})
}

func PostData(c *gin.Context){
	c.JSON(200,gin.H){
		"message":"Data Posted Successfully"
	}