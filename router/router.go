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
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var Router *gin.Engine

func init() {
	Router = gin.New()
	Router.MaxMultipartMemory = 200 << 20 // 64 MiB

}

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
	//g.Use(middlewares.Cors()) //开启中间件 允许使用跨域请求

}

//  注册 API 路由
func RegisterAPIRoutes(g *gin.Engine) {

	g.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
		g4.POST("/password", user.UpdatePassword)
		g4.POST("/name", user.UpdateName)
	}
	g5 := g.Group("/api/v1/character")
	{
		g5.POST("/", character.UpdateCharacter)
		g5.POST("/default", character.DefaultCharacterInit)
		g5.POST("/records", character.UpdateStudyRecords)
		g5.GET("/records/limit", character.GetStudyRecordsByTime)
		g5.GET("/records/num", character.GetStudyRecordsNum)
	}
	g6 := g.Group("/api/v1/download")
	{

		g6.GET("/video", video.DownloadVideo)
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
