package video

import (
	db "github.com/HanmoLiteracyBackend/database"
	"github.com/HanmoLiteracyBackend/model"
)

type VideoModel struct {
	model.BaseModel
	Path string `json:"name,omitempty" gorm:"column:path;" binding:"required"`
}

func (v *VideoModel) TableName() string {
	return "tbl_video"
}

func (v *VideoModel) Create() error {
	return db.DB.
		Table("tbl_video").
		Create(v).Error
}

func (u *VideoModel) Save() (err error) {
	return db.DB.
		Table("tbl_user").
		Save(u).Error
}

func GetVideoName(id int) (v *VideoModel, err error) {
	d := db.DB.
		Table("tbl_video").
		Where("id = ?", id).First(&v)
	return v, d.Error
}

type VideoCharacterModel struct {
	model.BaseModel
	VideoId     int `json:"vid" gorm:"column:vid;" binding:"required"`
	CharacterId int `json:"cid,omitempty" gorm:"column:cid;" binding:"required"`
}

func GetVideoByCharacterId(cid int) (v *VideoCharacterModel, err error) {
	d := db.DB.
		Table("tbl_video_records").
		Where("cid = ?", cid).First(v)
	return v, d.Error
}

type UserCharacterModel struct {
	model.BaseModel
	UserId        int    `json:"uid" gorm:"column:uid;" binding:"required"`
	CharacterId   int    `json:"cid" gorm:"column:cid;" binding:"required"`
	CharacterName string `json:"cname" gorm:"column:cname;" binding:"required"`
	StudyTime     string `json:"time" gorm:"column:time;" binding:"required"`
}

func (v *UserCharacterModel) Create() error {
	return db.DB.
		Table("tbl_study_records").
		Create(v).Error
}

func CharacterRecordsByUserId(uid int) (v []UserCharacterModel, err error) {
	d := db.DB.
		Table("tbl_study_records").
		Where("uid = ?", uid).Find(&v)
	return v, d.Error
}

func GetRecordsByTime(uid int, start string, end string) (v []UserCharacterModel, err error) {
	d := db.DB.
		Table("tbl_study_records").
		Where("time BETWEEN ? AND ? AND id = ?", start, end, uid).Find(&v)
	return v, d.Error
}
func GetRecordsNum(uid int) (v []UserCharacterModel, err error) {
	var count int64
	d := db.DB.
		Table("tbl_study_records").
		Where("id = ?", uid).Count(&count)
	return v, d.Error
}
