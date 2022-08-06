package video

import (
	db "github.com/HanmoLiteracyBackend/database"
	"github.com/HanmoLiteracyBackend/model"
)

type VideoModel struct {
	model.BaseModel
	Path        string `json:"name,omitempty" gorm:"column:path;" binding:"required"`
	CharacterId int    `json:"character_id" gorm:"column:character_id;" binding:"required"`
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
