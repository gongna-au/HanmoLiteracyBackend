package main

/*
   gin框架实现文件下载功能
*/

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strings"
)

//主函数
func main() {
	r := gin.Default()

	//Get路由，动态路由
	r.GET("/GetFile/:name", DowFile)
	//监听端口
	err := r.Run(":8088")
	if err != nil {
		fmt.Println("error")
	}

}

//文件下载功能实现
func DowFile(c *gin.Context) {
	//通过动态路由方式获取文件名，以实现下载不同文件的功能
	name := c.Param("name")
	//拼接路径,如果没有这一步，则默认在当前路径下寻找
	filename := path.Join("./", name)
	//响应一个文件
	c.File(filename)
	return
}

func GetPublicFilePath() string {
	dir, _ := os.Getwd()
	array := strings.Split(dir, "/")
	fmt.Print(array)
	sum := ""
	for _, v := range array {
		if v == "HanmoLiteracyBackend" {
			sum = sum + v + "/"
			break
		} else {
			sum = sum + v + "/"
		}

	}
	sum = sum + "public/video/test.mp4"
	return sum
}
