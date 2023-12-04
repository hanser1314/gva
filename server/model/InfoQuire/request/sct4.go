package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SctByTno struct {
	InfoQuire.Sct
	Tno string `json:"tno" form:"tno"`
	request.PageInfo
}
