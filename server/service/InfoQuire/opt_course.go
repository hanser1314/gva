package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    InfoQuireReq "github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire/request"
)

type OptCourseService struct {
}

// CreateOptCourse 创建选课表记录
// Author [piexlmax](https://github.com/piexlmax)
func (optCourseService *OptCourseService) CreateOptCourse(optCourse *InfoQuire.OptCourse) (err error) {
	err = global.GVA_DB.Create(optCourse).Error
	return err
}

// DeleteOptCourse 删除选课表记录
// Author [piexlmax](https://github.com/piexlmax)
func (optCourseService *OptCourseService)DeleteOptCourse(optCourse InfoQuire.OptCourse) (err error) {
	err = global.GVA_DB.Delete(&optCourse).Error
	return err
}

// DeleteOptCourseByIds 批量删除选课表记录
// Author [piexlmax](https://github.com/piexlmax)
func (optCourseService *OptCourseService)DeleteOptCourseByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]InfoQuire.OptCourse{},"id in ?",ids.Ids).Error
	return err
}

// UpdateOptCourse 更新选课表记录
// Author [piexlmax](https://github.com/piexlmax)
func (optCourseService *OptCourseService)UpdateOptCourse(optCourse InfoQuire.OptCourse) (err error) {
	err = global.GVA_DB.Save(&optCourse).Error
	return err
}

// GetOptCourse 根据id获取选课表记录
// Author [piexlmax](https://github.com/piexlmax)
func (optCourseService *OptCourseService)GetOptCourse(id uint) (optCourse InfoQuire.OptCourse, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&optCourse).Error
	return
}

// GetOptCourseInfoList 分页获取选课表记录
// Author [piexlmax](https://github.com/piexlmax)
func (optCourseService *OptCourseService)GetOptCourseInfoList(info InfoQuireReq.OptCourseSearch) (list []InfoQuire.OptCourse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&InfoQuire.OptCourse{})
    var optCourses []InfoQuire.OptCourse
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["cno"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&optCourses).Error
	return  optCourses, total, err
}
