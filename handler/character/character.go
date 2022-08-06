package character

import (
	//"github.com/HanmoLiteracyBackend/database"
	"github.com/HanmoLiteracyBackend/handler"
	//"github.com/HanmoLiteracyBackend/model"
	"errors"
	"github.com/HanmoLiteracyBackend/model/character"
	"github.com/HanmoLiteracyBackend/model/requests"
	"github.com/HanmoLiteracyBackend/model/response"
	"github.com/gin-gonic/gin"
)

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
				return
			}
		}
	}

	response.Success(c)
}

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
