package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/products", getProducts)
	router.POST("/products", addProduct)
	router.PUT("/products/:id", updateProduct)
	router.POST("/orders", placeOrder)
	router.GET("/orders/:customerId", getOrderHistory)
	router.GET("/transactions", getTransactions)

	router.Run("localhost:8080")
}

func getProducts(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Get all products"})
}

func addProduct(c *gin.Context) {
	
	c.JSON(http.StatusOK, gin.H{"message": "Add a new product"})
}

func updateProduct(c *gin.Context) {
	
	c.JSON(http.StatusOK, gin.H{"message": "Update product"})
}

func placeOrder(c *gin.Context) {
	
	c.JSON(http.StatusOK, gin.H{"message": "Place an order"})
}

func getOrderHistory(c *gin.Context) {
	
	c.JSON(http.StatusOK, gin.H{"message": "Get order history"})
}

func getTransactions(c *gin.Context) {
	
	c.JSON(http.StatusOK, gin.H{"message": "Get all transactions"})
}
