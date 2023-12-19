// main.go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key") // Change this to a secure secret key

func main() {
    router := gin.Default()

    // Routes
    router.POST("/register", register)
    router.POST("/login", login)

    // Run the server
    router.Run(":8081")
}

func register(c *gin.Context) {
    // Implementation for user registration
    // ...

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func login(c *gin.Context) {
    // Implementation for user login
    // ...

    // Generate JWT token
    token := generateToken("example-user")
    c.JSON(http.StatusOK, gin.H{"token": token})
}

func generateToken(username string) string {
    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
    })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        panic("Error generating token")
    }

    return tokenString
}
