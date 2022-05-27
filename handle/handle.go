package deal

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go_web/model"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// HandleQuery 处理query  /api
func HandleQuery(c *gin.Context) {
	name := c.Query("username") //http://localhost:8000/?name=lgp   //获取get请求的查询参数
	password := c.DefaultQuery("password", "123456")
	c.String(http.StatusOK, `<div style="color:red"">欢迎%v的驾到</div><div>密码%v</div>`, name, password)
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
	user := c.DefaultPostForm("user", "lgp")
	c.String(200, "username:%v,password:%v,user:%v", username, password, user)
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

func HandleFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传图片出错")
	}
	err = c.SaveUploadedFile(file, "img/"+file.Filename)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{"message": file.Filename})
	//c.String(http.StatusOK, file.Filename)
}

//自定义一个字符串
var jwtkey = []byte("www.topgoer.com")
var str string

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func setting(ctx *gin.Context) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: 2,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	str = tokenString
	ctx.JSON(200, gin.H{"token": tokenString})
}

//解析token
func getting(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	//vcalidate token formate
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}

	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}
	fmt.Println(111)
	fmt.Println(claims.UserId)
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}

func HandleUserTest(ctx *gin.Context) {
	type User struct {
		gorm.Model
		Username string `json:"username"`
		Password string `json:"password"`
	}
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	var user = User{
		Username: username,
		Password: password,
	}
	//建表
	err := model.Mydb.AutoMigrate(&User{})
	if err != nil {
		return
	}
	//插入数据
	result := model.Mydb.Create(&user)
	affected := result.RowsAffected
	fmt.Println(affected)
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: 10,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(), //生成时间
			Issuer:    "127.0.0.1",       // 签名颁发者
			Subject:   "user token",      //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	str = tokenString
	fmt.Println(user)
	ctx.JSON(200, gin.H{"token": tokenString, "user": user})
}

func HandleParse(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}

	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		ctx.Abort()
		return
	}
	fmt.Println(claims.UserId)
	//fmt.Println(time.Unix(claims.ExpiresAt,0))
	//b:=(time.Unix(claims.ExpiresAt,0)).Sub(time.Now())/24
	//fmt.Println(b)
	ctx.JSON(200, claims.UserId)
}
