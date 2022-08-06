package requests

import (
	//db "github.com/HanmoLiteracyBackend/database"
	//"github.com/HanmoLiteracyBackend/model/user"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type UpdateCharacterRequest struct {
	Name string `json:"name" valid:"name"`
}

func UpdateCharacterExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"name": []string{"required"},
	}
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"name": []string{
			"required:汉字为必填项，参数名称name",
		},
	}
	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}
