package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TeacherSearchByNameOrTno struct {
	InfoQuire.Teacher
	//StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	//EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Tname string `json:"tname" form:"tname"`
	Tno   string `json:"tno" form:"tno"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
