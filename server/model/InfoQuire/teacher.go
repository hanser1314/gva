// 自动生成模板Teacher
package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 教师表 结构体  Teacher
type Teacher struct {
      global.GVA_MODEL
      Tno  string `json:"tno" form:"tno" gorm:"column:tno;comment:;size:8;"`  //教工号 
      Tname  string `json:"tname" form:"tname" gorm:"column:tname;comment:;size:8;"`  //姓名 
      Tsex  string `json:"tsex" form:"tsex" gorm:"column:tsex;comment:;size:1;"`  //性别 
      Tage  *int `json:"tage" form:"tage" gorm:"column:tage;comment:;size:4;"`  //年龄 
      Teb  string `json:"teb" form:"teb" gorm:"column:teb;comment:;size:10;"`  //学历 
      Tpt  string `json:"tpt" form:"tpt" gorm:"column:tpt;comment:;size:10;"`  //职称 
      Cno1  string `json:"cno1" form:"cno1" gorm:"column:cno1;comment:;size:6;"`  //主讲课程一 
      Cno2  string `json:"cno2" form:"cno2" gorm:"column:cno2;comment:;size:6;"`  //主讲课程二 
      Cno3  string `json:"cno3" form:"cno3" gorm:"column:cno3;comment:;size:6;"`  //主讲课程三 
}


// TableName 教师表 Teacher自定义表名 teacher
func (Teacher) TableName() string {
  return "teacher"
}

