package user

import (
	"errors"
	//"fmt"
	"github.com/HanmoLiteracyBackend/handler"
	"github.com/HanmoLiteracyBackend/log"
	//"github.com/HanmoLiteracyBackend/model"
	"github.com/HanmoLiteracyBackend/model/requests"
	"github.com/HanmoLiteracyBackend/model/response"
	"github.com/HanmoLiteracyBackend/model/user"
	"github.com/HanmoLiteracyBackend/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	handler.BaseAPIController
}

// ShowAccount godoc
// @Summary      UpdatePassword
// @Description  update  user password by token 、 old password and new password
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.SignupUsingPhoneRequest  true  "OldPassword--旧密码 || NewPassword--新密码"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /password   [post]
// UpdatePassword   修改用户密码
func UpdatePassword(c *gin.Context) {
	// 1. 解析token
	claim, err := jwt.NewJWT().ParserToken(c)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	//验证表单
	request := requests.UpdatePassword{}
	if ok := handler.Validate(c, &request, requests.UpdateUserPassword); !ok {
		response.VerificationFailed(c, errors.New("请求格式不正确"))
		return
	}
	//根据用户id 去比较密码
	user, err := user.GetUserById(claim.UserIdToInt())
	if err != nil {
		response.BadRequest(c, err, "user not exist")
		return
	}
	correct := user.ComparePassword(request.OldPassword)
	if !correct {
		response.BadRequest(c, err, "The original password is incorrect, please re-enter")
		return
	}
	//如果密码正确
	user.Password = request.NewPassword
	err = user.Save()
	if err != nil {
		response.Abort500(c, "update password in  database err")
		return
	}

	response.Success(c)

}

// ShowAccount godoc
// @Summary      UpdateName
// @Description  update user name  by  token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.UpdateName   true  "Name--新昵称"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /name    [post]
//UpdateName 修改用户昵称
func UpdateName(c *gin.Context) {
	// 1. 解析token
	claim, err := jwt.NewJWT().ParserToken(c)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	//验证表单
	request := requests.UpdateName{}
	if ok := handler.Validate(c, &request, requests.UpdateUserName); !ok {
		response.VerificationFailed(c, errors.New("请求格式不正确"))
		return
	}
	//根据用户id 去比较密码
	user, err := user.GetUserById(claim.UserIdToInt())
	if err != nil {
		response.BadRequest(c, err, "user not exist")
		return
	}
	user.Name = request.Name
	err = user.Save()
	if err != nil {
		response.Abort500(c, "update name in database err")
		return
	}
	response.Success(c)
}

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) (userModel user.UserModel, err error) {

	defer func() {
		if x := recover(); x != nil {
			//处理panic, 让程序从panicking状态恢复的机会
			err = errors.New("Key current_user does not exist")
			//c.AbortWithError(400, errors.New("Get CurrentUser error"))

		}
	}()

	userModel, ok := c.MustGet("current_user").(user.UserModel)
	if !ok {
		log.LogIf(errors.New("无法获取用户"))
		return userModel, errors.New("get current_user failed")
	}

	// db is now a *DB value
	return userModel, err
}

// GetUsers所有用户
func GetUsers(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := handler.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}
