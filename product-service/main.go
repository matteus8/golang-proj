// main.go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Routes
    router.GET("/products", getProducts)

    // Run the server
    router.Run(":8082")
}

func getProducts(c *gin.Context) {
    // Implementation to fetch and return product list
    // ...

    products := []string{"Product1", "Product2", "Product3"}
    c.JSON(http.StatusOK, gin.H{"products": products})
}
