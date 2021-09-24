# 学习网址
https://www.liwenzhou.com/posts/Go/Gin_framework/

## 网页资源处理

## 返回json
```go
r := gin.Default()
// `json`表示当json包访问结构体时，用name去替代Name
type msg struct {
    Name string `json:"name"`
    Message string
    Age int
}

r.GET("/json", func(c *gin.Context) {
data := msg{
"笑大方",
"fjoafaosk",
12,
}
//结构体的属性要以大写开头，才能被访问
c.JSON(http.StatusOK,data)
})
```

## 获取url参数
```go
name , ok := c.GetQuery("query")
if !ok {
//取不到
name = "someone"
}
//取不到用默认值
age := c.DefaultQuery("age","18")
```

## 获取path参数
定义path参数的uri时要注意不要匹配的路径冲突
```go
//获取的path参数都是字符串类型
	r.GET("/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"name": name,
			"age":  age,
		})
	})
```

## 参数绑定
把请求的参数自动绑定到结构体中
```go
// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// 绑定JSON的示例 ({"user": "q1mi", "password": "123456"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var login Login

		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定form表单示例 (user=q1mi&password=123456)
	router.POST("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定QueryString示例 (/loginQuery?user=q1mi&password=123456)
	router.GET("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
```
