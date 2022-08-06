package signup

import (
	//"fmt"

	"errors"
	"fmt"
	//"time"

	"github.com/HanmoLiteracyBackend/handler"
	//"github.com/HanmoLiteracyBackend/model"
	"github.com/HanmoLiteracyBackend/model/requests"
	"github.com/HanmoLiteracyBackend/model/response"
	"github.com/HanmoLiteracyBackend/model/user"
	"github.com/HanmoLiteracyBackend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	handler.BaseAPIController
}

// SignupUsingPhone 使用手机和密码进行注册
func SignupUsingPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignupUsingPhoneRequest{}

	//requests.SignupUsingPhone 验证函数
	if ok := handler.Validate(c, &request, requests.SignupUsingPhone); !ok {
		response.VerificationFailed(c, errors.New("请求格式不正确"))
		return
	}

	err := user.JudgePhoneExist(request.Phone)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	err = user.JudgeGender(request.Gender)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	// 2. 验证成功，创建数据
	userModel := user.UserModel{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
		Gender:   request.Gender,
	}

	err = userModel.Create()
	if err != nil {
		response.Abort500(c, err.Error())
	}

	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
		fmt.Print("token is", token)

		response.CreatedJSON(c, gin.H{
			"token": token,
			"data":  userModel,
		})

	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
	response.Success(c)

}
