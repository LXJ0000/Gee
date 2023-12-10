package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

func ForV1() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.GET("/", func(context *gee.Context) {
		context.HTML(http.StatusOK, "<h1>Index Page With Logger</h1>")
	})

	v1 := r.Group("/v1")
	v1.Use(ForV1())
	{
		v1.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}
	_ = r.Run(":80")
}
