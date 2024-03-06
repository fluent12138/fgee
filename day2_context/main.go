package main

import (
	"fmt"
	"net/http"

	"fgee"
)

func main() {
	fmt.Printf("fgee 启动....")
	r := fgee.New()
	r.GET("/", func(c *fgee.Context) {
		c.HTML(http.StatusOK, "<div>hello fgee</div>")
	})

	r.GET("/hello", func(c *fgee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *fgee.Context) {
		c.JSON(http.StatusOK, fgee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
