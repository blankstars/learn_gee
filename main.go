package main

import (
	"github.com/blankstars/learn_gee/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Get("/", func(c *gee.Context) {
		c.String(http.StatusOK, "hello world: %v", c.Path)
	})

	r.Post("/hello", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})

	})

	r.Run(":8080")
}
