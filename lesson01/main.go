package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func SayHi(c *gin.Context)  {
	c.JSON(200,gin.H{
		"message":"自由万岁",
	})
}

func main() {

	r := gin.Default()
	//r.LoadHTMLFiles("templates/posts/index.tmpl") //解析模板
	//r.GET("/posts/index", func(c *gin.Context) {
	//		c.HTML(http.StatusOK,"index.tmpl",gin.H{
	//			"title":"哈哈哈",
	//		})
	//	})

	//设置自定义函数，必须在解析模板文件前设置
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	//加载静态文件,替换静态文件的路径
	r.Static("/xxx","./statics")

	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"posts/index.html",gin.H{
			"title":"哈哈哈",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"users/index.html",gin.H{
			"title":"<a href='https://baidu.com'>百度链接</a>",
		})
	})
	//默认以8080端口号启动
	r.Run(":9090")
}
