package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire"
	InfoQuireReq "github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TeacherService struct {
}

// CreateTeacher 创建教师表记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherService *TeacherService) CreateTeacher(teacher *InfoQuire.Teacher) (err error) {
	err = global.GVA_DB.Create(teacher).Error
	return err
}

// DeleteTeacher 删除教师表记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherService *TeacherService) DeleteTeacher(teacher InfoQuire.Teacher) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&teacher).Error
	return err
}

// DeleteTeacherByIds 批量删除教师表记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherService *TeacherService) DeleteTeacherByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]InfoQuire.Teacher{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTeacher 更新教师表记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherService *TeacherService) UpdateTeacher(teacher InfoQuire.Teacher) (err error) {
	err = global.GVA_DB.Save(&teacher).Error
	return err
}

// GetTeacher 根据id获取教师表记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherService *TeacherService) GetTeacher(id uint) (teacher InfoQuire.Teacher, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&teacher).Error
	return
}

// GetTeacherInfoList 分页获取教师表记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherService *TeacherService) GetTeacherInfoList(info InfoQuireReq.TeacherSearch) (list []InfoQuire.Teacher, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&InfoQuire.Teacher{})
	var teachers []InfoQuire.Teacher
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["tno"] = true
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

	err = db.Find(&teachers).Error
	return teachers, total, err
}

func (teacherService *TeacherService) GetTeacherInfoListByNameOrTno(info InfoQuireReq.TeacherSearchByNameOrTno) (list []InfoQuire.Teacher, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&InfoQuire.Teacher{})
	var teachers []InfoQuire.Teacher
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Tname != "" && info.Tno != "" {
		db = db.Where("tname = ? or tno = ?", info.Tname, info.Tno)
	}
	if info.Tname == "" && info.Tno != "" {
		db = db.Where("tno = ?", info.Tno)
	}
	if info.Tname != "" && info.Tno == "" {
		db = db.Where("tname = ? ", info.Tname)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["tno"] = true
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

	err = db.Find(&teachers).Error
	return teachers, total, err
}
