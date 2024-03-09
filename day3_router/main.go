package main

import (
	"fmt"
	"net/http"

	"fgee"
)

/*
测试

curl "http://localhost:9999/"
<div>hello fgee</div>

curl "http://localhost:9999/hello"
hello , you're at /hello

curl "http://localhost:9999/hello/fluent"
hello fluent, you're at /hello/fluent

curl "http://localhost:9999/assets/css/fluent.css"
{"filepath":"css/fluent.css"}

curl "http://localhost:9999/login" -X POST -d 'username=fluent&password=1234'
{"password":"1234","username":"fluent"}
*/

func main() {
	fmt.Printf("fgee 启动....")
	r := fgee.New()
	r.GET("/", func(c *fgee.Context) {
		c.HTML(http.StatusOK, "<div>hello fgee</div>")
	})

	r.GET("/hello", func(c *fgee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *fgee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *fgee.Context) {
		c.JSON(http.StatusOK, fgee.H{"filepath": c.Param("filepath")})
	})

	r.POST("/login", func(c *fgee.Context) {
		c.JSON(http.StatusOK, fgee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
