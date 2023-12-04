package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type StudentSearchByDept struct {
	InfoQuire.Student
	Sdept string `json:"sdept" form:"sdept"`
	request.PageInfo
}
