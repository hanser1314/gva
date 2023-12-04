package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type StudentSearch struct {
	InfoQuire.Student
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	//Sname *string `json:"sname" form:"sname"`
	//sno string `json:"sno" form:"sno"`
	request.PageInfo
}
