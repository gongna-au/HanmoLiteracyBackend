package requests

import (
	//db "github.com/HanmoLiteracyBackend/database"
	//"github.com/HanmoLiteracyBackend/model/user"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"mime/multipart"
)

type UpdateVideoRequest struct {
	Video *multipart.FileHeader `valid:"video" form:"video"`
} //@name UpdateVideoRequest

// LoginByPhone 验证表单，返回长度等于零即通过
func UpdateVideo(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// size 的单位为 bytes
		// - 1024 bytes 为 1kb
		// - 1048576 bytes 为 1mb
		// - 20971520 bytes 为 20mb
		"file:video": []string{"ext:mp4", "size:100000000", "required"},
	}
	messages := govalidator.MapData{
		"file:video": []string{
			"ext:只能上传mp4",
			"size:文件大小不能超过100MB",
			"required:必须上传文件",
		},
	}

	// 配置选项
	opts := govalidator.Options{
		Request:       c.Request,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	return govalidator.New(opts).ValidateStruct()
}

type DownLoadVideoRequest struct {
	Name string `json:"name"`
} //@name DownLoadVideoRequest
