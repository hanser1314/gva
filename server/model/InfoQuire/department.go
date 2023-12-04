// 自动生成模板Department
package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 院系信息表 结构体  Department
type Department struct {
      global.GVA_MODEL
      Dno  string `json:"dno" form:"dno" gorm:"column:dno;comment:;size:30;"`  //系编号 
      Dname  string `json:"dname" form:"dname" gorm:"column:dname;comment:;size:30;"`  //系名 
      Dmanagerno  string `json:"dmanagerno" form:"dmanagerno" gorm:"column:dmanagerno;comment:;size:30;"`  //系主任 
}


// TableName 院系信息表 Department自定义表名 department
func (Department) TableName() string {
  return "department"
}

