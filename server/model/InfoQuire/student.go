// 自动生成模板Student
package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// student表 结构体  Student
type Student struct {
      global.GVA_MODEL
      Sno  string `json:"sno" form:"sno" gorm:"column:sno;comment:;"`  //sno字段 
      Sname  string `json:"sname" form:"sname" gorm:"column:sname;comment:;size:8;"`  //sname字段 
      Ssex  string `json:"ssex" form:"ssex" gorm:"column:ssex;comment:;"`  //ssex字段 
      Sage  *int `json:"sage" form:"sage" gorm:"column:sage;comment:;size:10;"`  //sage字段 
      Sdept  string `json:"sdept" form:"sdept" gorm:"column:sdept;comment:;size:30;"`  //sdept字段 
}


// TableName student表 Student自定义表名 student
func (Student) TableName() string {
  return "student"
}

