package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
		c.SetCookie("csrfToken", randSeq(16), 3600, "/", "127.0.0.1", true, true)
	})

	// Per route middleware, you can add as many as you desire.
	r.POST("/auth", csrfCheck, benchEndpoint)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

func benchEndpoint(c *gin.Context) {
	c.SetCookie("bench", randSeq(16), 3600, "/auth", "127.0.0.1", true, true)
}

// MyBenchLogger get benchmark info
func csrfCheck(c *gin.Context) {
	tk, err := c.Cookie("csrfToken")
	if err != nil {
		c.String(http.StatusOK, "error")
	}

	c.String(http.StatusForbidden, tk)
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
