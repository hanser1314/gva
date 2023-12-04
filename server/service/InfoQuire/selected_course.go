package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    InfoQuireReq "github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire/request"
)

type SelectedCourseService struct {
}

// CreateSelectedCourse 创建已选课程表记录
// Author [piexlmax](https://github.com/piexlmax)
func (selectedCourseService *SelectedCourseService) CreateSelectedCourse(selectedCourse *InfoQuire.SelectedCourse) (err error) {
	err = global.GVA_DB.Create(selectedCourse).Error
	return err
}

// DeleteSelectedCourse 删除已选课程表记录
// Author [piexlmax](https://github.com/piexlmax)
func (selectedCourseService *SelectedCourseService)DeleteSelectedCourse(selectedCourse InfoQuire.SelectedCourse) (err error) {
	err = global.GVA_DB.Delete(&selectedCourse).Error
	return err
}

// DeleteSelectedCourseByIds 批量删除已选课程表记录
// Author [piexlmax](https://github.com/piexlmax)
func (selectedCourseService *SelectedCourseService)DeleteSelectedCourseByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]InfoQuire.SelectedCourse{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSelectedCourse 更新已选课程表记录
// Author [piexlmax](https://github.com/piexlmax)
func (selectedCourseService *SelectedCourseService)UpdateSelectedCourse(selectedCourse InfoQuire.SelectedCourse) (err error) {
	err = global.GVA_DB.Save(&selectedCourse).Error
	return err
}

// GetSelectedCourse 根据id获取已选课程表记录
// Author [piexlmax](https://github.com/piexlmax)
func (selectedCourseService *SelectedCourseService)GetSelectedCourse(id uint) (selectedCourse InfoQuire.SelectedCourse, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&selectedCourse).Error
	return
}

// GetSelectedCourseInfoList 分页获取已选课程表记录
// Author [piexlmax](https://github.com/piexlmax)
func (selectedCourseService *SelectedCourseService)GetSelectedCourseInfoList(info InfoQuireReq.SelectedCourseSearch) (list []InfoQuire.SelectedCourse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&InfoQuire.SelectedCourse{})
    var selectedCourses []InfoQuire.SelectedCourse
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
         	orderMap["sno"] = true
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
	
	err = db.Find(&selectedCourses).Error
	return  selectedCourses, total, err
}
