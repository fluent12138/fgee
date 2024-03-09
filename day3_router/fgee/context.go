package fgee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// Context 上下文结构
// 在ServeHTTP方法中被创建
type Context struct {
	// 请求与响应
	Req    *http.Request
	Writer http.ResponseWriter
	// 请求信息
	Path   string
	Method string
	// 响应信息
	StatusCode int
}

// 创建上下文
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm 方法返回Params中key对应的值
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 方法查询请求路由中附带的参数 [?key=val]
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 设置上下文响应状态
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置响应头
func (c *Context) SetHeader(key string, val string) {
	c.Writer.Header().Set(key, val)
}

// String 响应格式为字符串文本
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 响应格式为JSON
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 设置响应数据
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 响应格式为HTML
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
