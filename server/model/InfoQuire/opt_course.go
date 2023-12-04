// 自动生成模板OptCourse
package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 选课表 结构体  OptCourse
type OptCourse struct {
      global.GVA_MODEL
      Cno  string `json:"cno" form:"cno" gorm:"column:cno;comment:;size:6;"`  //课程号 
      Cname  string `json:"cname" form:"cname" gorm:"column:cname;comment:;size:50;"`  //课程名 
      Cpno  string `json:"cpno" form:"cpno" gorm:"column:cpno;comment:;size:6;"`  //先行课编号 
      Ccredit  *int `json:"ccredit" form:"ccredit" gorm:"column:ccredit;comment:;size:2;"`  //学分 
}


// TableName 选课表 OptCourse自定义表名 optcourse
func (OptCourse) TableName() string {
  return "optcourse"
}

