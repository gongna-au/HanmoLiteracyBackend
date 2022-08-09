package requests

import (
	//db "github.com/HanmoLiteracyBackend/database"
	//"github.com/HanmoLiteracyBackend/model/user"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone" valid:"phone"`
} //@name SignupPhoneExistRequest

func SignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
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

// SignupUsingPhoneRequest 通过手机注册的请求信息
type SignupUsingPhoneRequest struct {
	Phone    string `valid:"phone" json:"phone,omitempty" gorm:"column:phone;" binding:"required"`
	Name     string `valid:"name" json:"name,omitempty" gorm:"column:name;" binding:"required"`
	Password string `valid:"password" json:"password,omitempty" gorm:"column:password;" binding:"required"`
	Gender   string `valid:"gender" json:"gender,omitempty" gorm:"column:gender;" binding:"required"`
} //@name SignupUsingPhoneRequest

func SignupUsingPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone":    []string{"required", "digits:11"},
		"name":     []string{"required", "alpha_num", "between:3,20"},
		"password": []string{"required", "min:6"},
		"gender":   []string{"required"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度必须为 11 位的数字",
		},
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"gender": []string{
			"required:性别为必填项",
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
