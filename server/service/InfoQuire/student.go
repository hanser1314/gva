package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire"
	InfoQuireReq "github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type StudentService struct {
}

// CreateStudent 创建student表记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) CreateStudent(student *InfoQuire.Student) (err error) {
	err = global.GVA_DB.Create(student).Error
	return err
}

// DeleteStudent 删除student表记录
// Author [piexlmax](https://github.com/piexlmax)
//
//	func (studentService *StudentService) DeleteStudent(student InfoQuire.Student) (err error) {
//		err = global.GVA_DB.Delete(&student).Error
//		return err
//	}
func (studentService *StudentService) DeleteStudent(student InfoQuire.Student) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&student).Error
	return err
}

// DeleteStudentByIds 批量删除student表记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) DeleteStudentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]InfoQuire.Student{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateStudent 更新student表记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) UpdateStudent(student InfoQuire.Student) (err error) {
	err = global.GVA_DB.Save(&student).Error
	return err
}

// GetStudent 根据id获取student表记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) GetStudent(id uint) (student InfoQuire.Student, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&student).Error
	return
}

// GetStudentInfoList 分页获取student表记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentService *StudentService) GetStudentInfoList(info InfoQuireReq.StudentSearch) (list []InfoQuire.Student, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&InfoQuire.Student{})
	var students []InfoQuire.Student
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&students).Error
	return students, total, err
}

func (studentService *StudentService) GetStudentInfoListByName(info InfoQuireReq.StudentSearchByName) (list []InfoQuire.Student, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&InfoQuire.Student{})
	var students []InfoQuire.Student
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Sname != "" && info.Sage != 0 {
		db = db.Where("sname like ? or sage = ?", "%"+info.Sname+"%", info.Sage)
	}
	if info.Sname == "" && info.Sage != 0 {
		db = db.Where("sage = ?", info.Sage)
	}
	if info.Sname != "" {
		db = db.Where("sname like ?", "%"+info.Sname+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&students).Error
	return students, total, err
}

func (studentService *StudentService) GetStudentInfoListBySno(info InfoQuireReq.StudentSearchBySno) (list []InfoQuire.Student, total int64, err error) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&InfoQuire.Student{})
	var students []InfoQuire.Student
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Sno != "" {
		db = db.Where("sno = ? ", info.Sno)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	//if limit != 0 {
	//	db = db.Limit(limit).Offset(offset)
	//}

	err = db.Find(&students).Error
	return students, total, err
}

func (studentService *StudentService) GetStudentInfoListByDept(info InfoQuireReq.StudentSearchByDept) (list []InfoQuire.Student, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&InfoQuire.Student{})
	var students []InfoQuire.Student
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Sdept != "" {
		db = db.Where("sdept = ? ", info.Sdept)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&students).Error
	return students, total, err
}
