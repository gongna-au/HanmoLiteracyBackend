package user

import (
	"errors"
	"fmt"

	db "github.com/HanmoLiteracyBackend/database"
	"github.com/HanmoLiteracyBackend/model"
	"github.com/HanmoLiteracyBackend/pkg/paginator"
	"github.com/gin-gonic/gin"
)

type UserModel struct {
	model.BaseModel
	Name     string `json:"name,omitempty" gorm:"column:name;" binding:"required"`
	Phone    string `json:"phone" gorm:"column:phone;" binding:"required"`
	Password string `json:"password" gorm:"column:password;" binding:"required"`
	Gender   string `json:"gender" gorm:"column:gender;" binding:"required"`
}

func (u *UserModel) TableName() string {
	return "tbl_user"
}

func (u *UserModel) Create() error {
	return db.DB.
		Table("tbl_user").
		Create(u).Error
}

func (u *UserModel) Save() (err error) {
	return db.DB.
		Table("tbl_user").
		Save(u).Error
}

// ComparePassword 密码是否正确
func (u *UserModel) ComparePassword(password string) bool {
	if password == u.Password {
		return true
	} else {
		return false
	}
}

func GetUserByName(phone string) (*UserModel, error) {
	u := &UserModel{}
	d := db.DB.
		Table("tbl_user").
		Where("phone = ?", phone).First(u)
	return u, d.Error
}

func GetUserByPhoneAndPassword(phone string, password string) (*UserModel, error) {
	u := &UserModel{}
	d := db.DB.
		Table("tbl_user").
		Where("phone = ? AND password = ?", phone, password).First(u)
	return u, d.Error
}

func GetUserById(uid int) (*UserModel, error) {
	u := &UserModel{}
	d := db.DB.
		Table("tbl_user").
		Where("id = ?", uid).First(u)
	return u, d.Error
}

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string) (userModel UserModel) {
	db.DB.
		Table("tbl_user").
		Where("phone = ?", phone).First(&userModel)
	return userModel
}

// GetByPhone 通过手机号来判断手机号是非被注册
func JudgePhoneExist(phone string) error {
	var userModel UserModel
	db.DB.Where("phone = ?", phone).First(&userModel)
	if userModel.BaseModel.ID > 0 {
		fmt.Println("this phone has been registered")
		return errors.New("this phone has been registered")
	} else {
		return nil
	}
}

// GetByPhone 通过手机号来判断手机号是非被注册
func JudgeGender(gender string) error {
	if gender == "male" || gender == "female" {
		return nil
	} else {
		return errors.New("Incorrect gender format")
	}
}

// GetByMulti 通过 手机号 密码来获取用户
func GetByMulti(phone string) (userModel UserModel) {
	db.DB.
		Table("tbl_user").
		Where("phone = ?", phone).
		First(&userModel)
	return
}

// Get 通过 ID 获取用户
func Get(id string) (userModel UserModel) {
	db.DB.
		Table("tbl_user").
		Where("id", id).
		First(&userModel)
	return
}

// All 获取所有用户数据
func All() (users []UserModel) {
	db.DB.
		Table("tbl_user").
		Find(&users)
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []UserModel, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		db.DB.Model(UserModel{}),
		&users,
		"http://localhost:8080",
		perPage,
	)
	return
}
