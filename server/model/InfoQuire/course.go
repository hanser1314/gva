// 自动生成模板Course
package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 课程表 结构体  Course
type Course struct {
	global.GVA_MODEL
	Cno     string `json:"cno" form:"cno" gorm:"column:cno;comment:;size:20;"`             //课程号
	Cname   string `json:"cname" form:"cname" gorm:"column:cname;comment:;size:20;"`       //课程名
	CPno    string `json:"cpno" form:"cpno" gorm:"column:cpno;comment:;size:6;"`           //先行课编号
	Ccredit *int   `json:"ccredit" form:"ccredit" gorm:"column:ccredit;comment:;size:2;"`  //学分
	Ctno    string `json:"ctno" form:"ctno" gorm:"column:ctno;comment:;size:20;"`          //教工号
	Cplace  string `json:"cplace" form:"cplace" gorm:"column:cplace;comment:;size:20;"`    //place字段
	Cmax    *int   `json:"cmax" form:"cmax" gorm:"column:cmax;comment:;size:10;"`          //最大选课人数
	Cremain *int   `json:"cremain" form:"cremain" gorm:"column:cremain;comment:;size:10;"` //已选人数
}

// TableName 课程表 Course自定义表名 course
func (Course) TableName() string {
	return "course"
}
