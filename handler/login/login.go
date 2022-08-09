package login

import (
	"errors"
)

import (
	"github.com/HanmoLiteracyBackend/handler"
	"github.com/HanmoLiteracyBackend/model/requests"
	"github.com/HanmoLiteracyBackend/model/response"
	"github.com/HanmoLiteracyBackend/model/user"
	"github.com/HanmoLiteracyBackend/pkg/jwt"
)

import (
	"github.com/gin-gonic/gin"
)

// LoginController 用户控制器
type LoginController struct {
	handler.BaseAPIController
}

// ShowAccount godoc
// @Summary      Login By Phone
// @Description  get user study records by start time and end time
// @Tags         login
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.LoginByPhoneRequest  true  "Phone--电话号码 和Password--密码"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Router       /usingphone    [post]
// LoginByPhone 手机登录
func LoginByPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.LoginByPhoneRequest{}

	if ok := handler.Validate(c, &request, requests.LoginByPhone); !ok {
		response.VerificationFailed(c, errors.New("请求格式不正确"))
		return
	}

	// 2. 尝试登录
	user, err := user.GetUserByPhoneAndPassword(request.Phone, request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在或密码错误")
	} else {
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
	response.Success(c)

}
