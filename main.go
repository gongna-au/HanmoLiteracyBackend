package main

import (
	_ "github.com/HanmoLiteracyBackend/config"
	_ "github.com/HanmoLiteracyBackend/database"
	_ "github.com/HanmoLiteracyBackend/log"
	_ "github.com/HanmoLiteracyBackend/model/character"
	"github.com/HanmoLiteracyBackend/router"
)

// 添加注释以描述 server 信息
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	//路由初始化
	router.SetupRoute(router.Router)
	router.Router.Run()

}
