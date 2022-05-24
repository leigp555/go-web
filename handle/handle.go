package deal

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

// HandleQuery 处理query
func HandleQuery(c *gin.Context) {
	name := c.Query("username") //http://localhost:8000/?name=lgp   //获取get请求的查询参数
	c.String(http.StatusOK, `<div style="color:red"">欢迎%v的驾到</div>`, name)
}

// HandleParam 处理param
func HandleParam(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, name)
}

// HandleForm 处理post请求
func HandleForm(c *gin.Context) {
	type User struct {
		username string
		password string
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.String(200, "username:%v,password:%v", username, password)
}

// HandlePage 处理文件请求
func HandlePage(c *gin.Context) {
	file, _ := os.OpenFile("index.html", os.O_RDWR, os.ModePerm)
	all, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	c.String(200, string(all))
}
