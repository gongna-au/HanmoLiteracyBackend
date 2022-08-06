package video

import (
	//"errors"
	"fmt"
	//"io"
	//"io"
	//"io/ioutil"
	//"net/http"
	"os"
	//"path"
	"github.com/HanmoLiteracyBackend/handler"

	//"github.com/HanmoLiteracyBackend/handler"
	//"github.com/HanmoLiteracyBackend/model/character"
	"github.com/HanmoLiteracyBackend/model/requests"
	"github.com/HanmoLiteracyBackend/model/response"
	//"github.com/HanmoLiteracyBackend/model/video"
	"errors"
	"github.com/HanmoLiteracyBackend/pkg/jwt"
	"github.com/gin-gonic/gin"
	mathrand "math/rand"
	"mime/multipart"
	"path/filepath"
	"time"
)

// fetchAllTodo 返回所有的 todo 数据
func FetchAllVideo(c *gin.Context) {

}

// fetchSingleTodo方法返回一条 todo 数据
func FetchSingleVideo(c *gin.Context) {

}

// 上传单个视频
func UpdateVideo(c *gin.Context) {
	file, err := c.FormFile("file")
	request := requests.UpdateVideoRequest{
		Video: file,
	}
	if err != nil {
		response.Abort500(c, "上传视频失败，请稍后尝试~")
		return
	}
	_, err = SaveUploadVideo(c, request.Video)
	if err != nil {
		response.Abort500(c, "上传视频失败，请稍后尝试~")
		return
	}
	response.Success(c)
}

//上传很多视频
func UpdateVideos(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.Abort500(c, "上传视频失败，请稍后尝试~")
	}
	// 获取所有视频
	files := form.File["files"]
	// 遍历所有视频
	for _, file := range files {
		// 逐个存
		_, err = SaveUploadVideo(c, file)
		if err != nil {
			response.Abort500(c, "上传视频失败，请稍后尝试~")
			return
		}
	}
	response.Success(c)
}

//下载单个视频
func DownloadVideo(c *gin.Context) {
	//表单验证
	request := requests.DownLoadVideoRequest{}
	//requests.SignupUsingPhone 验证函数
	if ok := handler.Validate(c, &request, requests.DownLoadVideo); !ok {
		response.VerificationFailed(c, errors.New("请求格式不正确，下载失败"))
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		response.Abort500(c, "下载视频失败，请稍后尝试~")
	}
	// 获取所有视频
	files := form.File["files"]
	// 遍历所有视频
	for _, file := range files {
		// 逐个存
		_, err = SaveUploadVideo(c, file)
		if err != nil {
			response.Abort500(c, "下载视频失败，请稍后尝试~")
			return
		}
	}
	response.Success(c)
}

//下载很多视频
func DownloadVideos(c *gin.Context) {
	request := requests.DownLoadVideosRequest{}
	if ok := handler.Validate(c, &request, requests.DownLoadVideos); !ok {
		response.VerificationFailed(c, errors.New("请求格式不正确,下载失败"))
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		response.Abort500(c, "下载视频失败，请稍后尝试~")
	}
	// 获取所有视频
	files := form.File["files"]
	// 遍历所有视频
	for _, file := range files {
		// 逐个存
		_, err = SaveUploadVideo(c, file)
		if err != nil {
			response.Abort500(c, "下载视频失败，请稍后尝试~")
			return
		}
	}
	response.Success(c)
}

func SaveUploadVideo(c *gin.Context, file *multipart.FileHeader) (string, error) {
	var video string
	// 确保目录存在，不存在创建
	publicPath := "public"
	dirName := fmt.Sprintf("/uploads/videos/%s/", jwt.TimenowInTimezone().Format("2006/01/02"))
	os.MkdirAll(publicPath+dirName, 0755)

	// 保存文件
	fileName := randomNameFromUploadFile(file)
	// public/uploads/videos/2021/12/22/nFDacgaWKpWWOmOt.png
	videoPath := publicPath + dirName + fileName
	if err := c.SaveUploadedFile(file, videoPath); err != nil {
		return video, err
	}
	return dirName + fileName, nil
}

func randomNameFromUploadFile(file *multipart.FileHeader) string {
	return RandomString(16) + filepath.Ext(file.Filename)
}

//RandomString 生成长度为 length 的随机字符串
func RandomString(length int) string {
	mathrand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[mathrand.Intn(len(letters))]
	}
	return string(b)
}
