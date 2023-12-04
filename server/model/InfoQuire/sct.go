// 自动生成模板Sct
package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 选课表 结构体  Sct
type Sct struct {
	global.GVA_MODEL
	Sno     string `json:"sno" form:"sno" gorm:"column:sno;comment:sno;size:20;"`                //学号
	Cno     string `json:"cno" form:"cno" gorm:"column:cno;comment:cno;size:20;"`                //课程名
	Cname   string `json:"cname" form:"cname" gorm:"column:cname;comment:cname;size:20;"`        //课程名
	Ccredit *int   `json:"ccredit" form:"ccredit" gorm:"column:ccredit;comment:ccredit;size:2;"` //学分
	Tno     string `json:"tno" form:"tno" gorm:"column:tno;comment:tno;size:20;"`                //教工号
	Cplace  string `json:"cplace" form:"cplace" gorm:"column:cplace;comment:cplace;size:50;"`    //上课地点
	Grade   int    `json:"grade" form:"grade" gorm:"column:grade;comment:grade;size:10;"`        //成绩
}

// TableName 选课表 Sct自定义表名 sct
func (Sct) TableName() string {
	return "sct"
}
