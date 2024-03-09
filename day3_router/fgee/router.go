package fgee

import (
	"log"
	"net/http"
)

// 路由结构定义
type router struct {
	handlers map[string]HandlerFunc
}

// 创建路由实例
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// 添加路径与方法
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// 找到对应的方法并执行
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUNT: %s\n", c.Path)
	}
}
