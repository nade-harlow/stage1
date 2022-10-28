package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type slackUser struct {
	SlackUsername string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	user := newUser()
	r.Use(CORSMiddleware())
	r.GET("/", user.GetSlackUser())
	err := r.Run()
	if err != nil {
		log.Println("An error occurred while starting up server")
		return
	}
}

func newUser() *slackUser {
	return &slackUser{
		SlackUsername: "Nade",
		Backend:       true,
		Age:           24,
		Bio:           "nil guy",
	}
}

func (user *slackUser) GetSlackUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, user)
	}
}
