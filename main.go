package main

import (
	//"fmt"
	//_ "github.com/HanmoLiteracyBackend/handler/login"
	//_ "github.com/HanmoLiteracyBackend/handler/signup"
	_ "github.com/HanmoLiteracyBackend/handler/todo"
	//_ "github.com/HanmoLiteracyBackend/handler/user"

	_ "github.com/HanmoLiteracyBackend/config"
	_ "github.com/HanmoLiteracyBackend/log"

	_ "github.com/HanmoLiteracyBackend/database"
	_ "github.com/HanmoLiteracyBackend/model/character"
	"github.com/HanmoLiteracyBackend/router"

	"github.com/gin-gonic/gin"
	//"github.com/HanmoLiteracyBackend/model"
	//"go.uber.org/zap"
	//"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin"
	//"time"
)

func main() {

	g := gin.New()
	//路由初始化
	router.SetupRoute(g)
	g.Run()

}
