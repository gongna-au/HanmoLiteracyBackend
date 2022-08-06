package router

import (
	"net/http"
	"strings"

	"github.com/HanmoLiteracyBackend/handler/character"
	"github.com/HanmoLiteracyBackend/handler/login"
	"github.com/HanmoLiteracyBackend/handler/signup"
	"github.com/HanmoLiteracyBackend/handler/user"
	"github.com/HanmoLiteracyBackend/handler/video"
	"github.com/HanmoLiteracyBackend/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	//  注册 API 路由
	RegisterAPIRoutes(router)

	//  配置 404 路由
	setup404Handler(router)
}

//注册全局中间件
func registerGlobalMiddleWare(g *gin.Engine) {
	g.Use(middlewares.Logger())

}

//  注册 API 路由
func RegisterAPIRoutes(g *gin.Engine) {

	g1 := g.Group("/api/v1/upload")
	{
		g1.POST("/video", video.UpdateVideo)
		g1.POST("/videos", video.UpdateVideos)

	}

	g2 := g.Group("/api/v1/auth")
	{
		//使用手机号和密码注册
		g2.POST("/signup/usingphone", signup.SignupUsingPhone)

	}

	g3 := g.Group("/api/v1/login")
	{
		g3.POST("/usingphone", login.LoginByPhone)

	}
	g4 := g.Group("/api/v1/user")
	{
		//修改用户信息
		g4.POST("/password", user.UpdatePassword)
		g4.POST("/name", user.UpdateName)
		//注销用户
		g4.DELETE("/", user.GetUsers)
		//修改个人资料
	}
	g5 := g.Group("/api/v1/character")
	{
		//修改用户信息
		g5.POST("/", character.UpdateCharacter)
		g5.POST("/default", character.DefaultCharacterInit)
	}
	g6 := g.Group("/api/v1/download")
	{
		//修改用户信息
		g6.GET("/video", video.DownloadVideo)
		g6.GET("/videos", video.DownloadVideos)
	}
}

//  配置 404 路由
func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

}
