package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SctBySnoAndCno struct {
	InfoQuire.Sct
	Sno string `json:"sno" form:"sno"`
	Cno string `json:"cno" form:"cno"`
	request.PageInfo
}
