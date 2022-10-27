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

func main() {
	r := gin.Default()
	user := newUser()
	r.GET("/", user.GetSlackUser())
	err := r.Run()
	if err != nil {
		log.Println("An error occurred while starting up server")
		return
	}
}

func newUser() *slackUser {
	return &slackUser{
		SlackUsername: "Nade 1️⃣",
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
