package main

import (
	"github.com/blankstars/learn_gee/gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.Use(gee.Logger(), gee.Recovery())

	r.Get("/", func(c *gee.Context) {
		c.String(http.StatusOK, "hello world: %v", c.Path)
	})

	r.Post("/hello", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})

	})

	r.Post("/hi/:name", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.Param("name"),
			"password": c.PostForm("password"),
		})

	})

	r.LoadHTMLGlob("templates/*")
	r.Static("assets", "./static")

	r.Get("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.Get("/panic", func(c *gee.Context) {
		var names []string
		c.String(http.StatusOK, names[0])
	})

	r.Run(":8080")
}
