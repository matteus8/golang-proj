// main.go is the main entry point for the auth service. It initializes 
// the router, configures routes, and starts the HTTP server.

// main initializes the Gin router, configures routes, and starts the server.
func main() {

// register handles user registration.
func register(c *gin.Context) {

// login handles user login and generates a JWT token.  
func login(c *gin.Context) {

// generateToken signs a JWT token using the provided username and secret key.
func generateToken(username string) string {
// main.go is the entry point for the auth service. It initializes the router,
// configures routes, and starts the HTTP server.

// register handles user registration.

// login handles user login and generates a JWT token. 

// generateToken signs a JWT token using the provided secret key.
// main.go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key") // Change this to a secure secret key

// main is the entry point for the application. It initializes the Gin router, 
// configures the routes for the /register and /login endpoints, and starts 
// the HTTP server on port 8081.
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
