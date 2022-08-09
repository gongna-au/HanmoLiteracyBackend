package video

import (
	//"errors"
	"fmt"
	//"path"
	//"io"
	//"io"
	//"io/ioutil"
	//"net/http"
	"os"
	//"path"
	//"github.com/HanmoLiteracyBackend/handler"
	"strings"

	//"github.com/HanmoLiteracyBackend/handler"
	"github.com/HanmoLiteracyBackend/model/character"
	"github.com/HanmoLiteracyBackend/model/requests"
	"github.com/HanmoLiteracyBackend/model/response"

	//"errors"
	"github.com/HanmoLiteracyBackend/model/video"
	mathrand "math/rand"
	multipart "mime/multipart"
	"path/filepath"
	"time"

	"github.com/HanmoLiteracyBackend/pkg/jwt"
	//"github.com/HanmoLiteracyBackend/router"
	"github.com/gin-gonic/gin"
)

// fetchAllTodo 返回所有的 todo 数据
func FetchAllVideo(c *gin.Context) {

}

// fetchSingleTodo方法返回一条 todo 数据
func FetchSingleVideo(c *gin.Context) {

}

// ShowAccount godoc
// @Summary      Update a Video
// @Description  Update a video
// @Tags         upload
// @Accept       json
// @Produce      json
// @Param        req  formData  multipart.FileHeader true "要上传的文件"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /video    [post]
// 上传单个视频
func UpdateVideo(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		response.Abort500(c, "上传视频失败，请稍后尝试~")
		return
	}
	_, err = SaveUploadVideo(c, file)
	if err != nil {
		response.Abort500(c, "上传视频失败，请稍后尝试~")
		return
	}
	response.Success(c)
}

// ShowAccount godoc
// @Summary      Update  Videos
// @Description  Update  videos
// @Tags         upload
// @Accept       json
// @Produce      json
// @Param        req  formData  multipart.FileHeader true "要上传的文件"
// @Success      200  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /videos    [post]
// 上传单个视频
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

// ShowAccount godoc
// @Summary      Update  Videos
// @Description  Update  videos
// @Tags         download
// @Accept       json
// @Produce      json
// @Param        req   body  requests.DownLoadVideoRequest  true "Name--要下载的汉字 返回该汉字对应的视频的文件"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /video    [get]
// 下载单独的一个视频
func DownloadVideo(c *gin.Context) {
	request := requests.DownLoadVideoRequest{}
	err := c.ShouldBind(&request)
	charact, err := character.GetCharacterIdByName(request.Name)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	vc, err := video.GetVideoByCharacterId(charact.BaseModel.ID)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	v, err := video.GetVideoName(vc.VideoId)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	//name := c.Param("name")
	c.File("public/" + v.Path)
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
func GetPath() string {
	dir, _ := os.Getwd()
	return dir
}
