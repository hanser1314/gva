package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SctSearchBySno struct {
	InfoQuire.Sct
	//StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	//EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Sno string `json:"sno" form:"sno"`
	request.PageInfo
}
