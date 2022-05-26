package route

import (
	"github.com/gin-gonic/gin"
	deal "go_web/handle"
	"net/http"
)

// Cors 跨域处理
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}
		defer func() {
			recover()
		}()
		c.Next()
	}
}

// Serve 路由
func Serve() *gin.Engine {
	app := gin.Default()
	app.Use(Cors())
	app.GET("/api", deal.HandleQuery)
	app.GET("/:name", deal.HandleParam)
	app.POST("/user/submit", deal.HandleForm)
	app.GET("/get/page", deal.HandlePage)
	app.POST("/upload", deal.HandleFile)
	return app
}
