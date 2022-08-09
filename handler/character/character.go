package character

import (
	"github.com/HanmoLiteracyBackend/handler"
	//"github.com/HanmoLiteracyBackend/model"
	"errors"
	"github.com/HanmoLiteracyBackend/model/character"
	"github.com/HanmoLiteracyBackend/model/requests"
	"github.com/HanmoLiteracyBackend/model/response"
	"github.com/HanmoLiteracyBackend/model/video"
	"github.com/HanmoLiteracyBackend/pkg/jwt"
	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      Update Character In Database
// @Description
// @Tags         character
// @Accept       json
// @Produce      json
// @Param        req            body    requests.UpdateCharacterRequest                 true  "请求中包含要上传的汉字的名称"
// @Success      200  {object}  response.Response
// @Failure      422  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /character/    [post]
func UpdateCharacter(c *gin.Context) {
	//初始化汉字是否存在
	count := character.JudgeDefaultCharacterIdExist()
	if count < 10 {
		character.CharacterInit()
	}
	// 1. 验证表单
	request := requests.UpdateCharacterRequest{}
	//requests.SignupUsingPhone 验证函数
	if ok := handler.Validate(c, &request, requests.UpdateCharacterExist); !ok {
		response.VerificationFailed(c, errors.New("请求格式不正确"))
		return
	}

	characters := character.RegexpCompile(request.Name)
	for _, v := range characters {
		t, _ := character.GetCharacterIdByName(v)
		if t.BaseModel.ID > 0 {
			continue
		} else {
			err := t.Save()
			if err != nil {
				response.Abort500(c, err.Error())
				return
			}
		}
	}

	response.Success(c)
}

// ShowAccount godoc
// @Summary      Insert Default Characters in Database
// @Description  use default characters
// @Tags         character
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Router       /default  [post]
// DefaultCharacterInit ... 初始化默认导入数据库的汉字
func DefaultCharacterInit(c *gin.Context) {
	//初始化汉字是否存在
	count := character.JudgeDefaultCharacterIdExist()
	if count < 10 {
		character.CharacterInit()
		response.Success(c)
	} else {
		response.BadRequest(c, errors.New("Default Character has Init"))
	}

}

// ShowAccount godoc
// @Summary      Create  A User Study Records
// @Description  cteate  by userId and post param
// @Tags         character
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string                     true  "JWT"
// @Param        req            body    requests.UpdateCharacterStudyRequest                true  "请求中包含要上传的汉字的名称"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Router       /records  [post]
// UpdateStudyRecords ... 上传用户已经学习完的一个汉字
func UpdateStudyRecords(c *gin.Context) {
	// 1. 解析token
	claim, err := jwt.NewJWT().ParserToken(c)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	//验证表单
	request := requests.UpdateCharacterStudyRequest{}
	err = c.ShouldBind(&request)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	ch, err := character.GetCharacterIdByName(request.CharacterName)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	t := jwt.TimeToString(jwt.TimenowInTimezone())
	v := video.UserCharacterModel{
		UserId:        claim.UserIdToInt(),
		CharacterId:   ch.BaseModel.ID,
		CharacterName: ch.Name,
		StudyTime:     t,
	}
	err = v.Create()
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	response.Success(c)
}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get user study records by token
// @Tags         character
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string                     true  "JWT"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Router      /records  [get]
func GetStudyRecords(c *gin.Context) {

	// 1. 解析token
	claim, err := jwt.NewJWT().ParserToken(c)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	results, err := video.CharacterRecordsByUserId(claim.UserIdToInt())
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	response.Data(c, results)

}

// ShowAccount godoc
// @Summary      Get user study records
// @Description  get user study records by start time and end time
// @Tags         character
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string                     true  "JWT"
// @Param        req  {object}  body  requests.GetStudyRecordByTimeRequest  true "StartTime-起始时间 EndTime --结束时间"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Router       /records/limit [get]
func GetStudyRecordsByTime(c *gin.Context) {

	// 1. 解析token
	claim, err := jwt.NewJWT().ParserToken(c)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	//验证表单
	request := requests.GetStudyRecordByTimeRequest{}
	err = c.ShouldBind(&request)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	if TimeCompare(request.EndTime) == false || TimeCompare(request.StartTime) == false {
		response.BadRequest(c, err, "time 格式不正确")
		return
	}
	results, err := video.GetRecordsByTime(claim.UserIdToInt(), request.StartTime, request.EndTime)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	response.Data(c, results)

}

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string    true  "JWT"
// @Success      200  {object}  response.Response
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Router       /records/num   [get]
func GetStudyRecordsNum(c *gin.Context) {

	// 1. 解析token
	claim, err := jwt.NewJWT().ParserToken(c)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	results, err := video.CharacterRecordsByUserId(claim.UserIdToInt())
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	response.Data(c, results)

}

func TimeCompare(str string) bool {
	array := []byte(str)
	for k, v := range array {
		if k == 4 || k == 7 {
			if string(v) != "-" {
				return false
			}
		}
		if k == 10 {

			if string(v) != " " {
				return false
			}
		}
		if k == 13 || k == 16 {
			if string(v) != ":" {
				return false
			}
		}
	}
	if len(str) > 19 {
		return false
	}
	return true
}
