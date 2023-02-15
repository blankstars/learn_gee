package main

import (
	"github.com/blankstars/learn_gee/gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.Use(gee.Logger())

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

	g := r.Group("/assets")
	g.Use(gee.Logger())

	g.Post("/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"filepath": c.Param("filepath"),
		})

	})

	r.Run(":8080")
}
