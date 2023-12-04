// 自动生成模板SelectedCourse
package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 已选课程表 结构体  SelectedCourse
type SelectedCourse struct {
      global.GVA_MODEL
      Sno  string `json:"sno" form:"sno" gorm:"column:sno;comment:;size:30;"`  //学号 
      Cno  string `json:"cno" form:"cno" gorm:"column:cno;comment:;size:30;"`  //课程号 
      Tno  string `json:"tno" form:"tno" gorm:"column:tno;comment:;size:30;"`  //教工号 
      Grade  *int `json:"grade" form:"grade" gorm:"column:grade;comment:;size:4;"`  //成绩 
}


// TableName 已选课程表 SelectedCourse自定义表名 selectedcourse
func (SelectedCourse) TableName() string {
  return "selectedcourse"
}

