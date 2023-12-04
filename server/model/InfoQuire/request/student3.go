package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire"
)

type StudentSearchBySno struct {
	InfoQuire.Student
	Sno string `json:"sno" form:"sno"`
	//request.PageInfo
}
