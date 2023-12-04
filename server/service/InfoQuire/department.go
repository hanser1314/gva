package InfoQuire

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire"
	InfoQuireReq "github.com/flipped-aurora/gin-vue-admin/server/model/InfoQuire/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type DepartmentService struct {
}

// CreateDepartment 创建院系信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (departmentService *DepartmentService) CreateDepartment(department *InfoQuire.Department) (err error) {
	err = global.GVA_DB.Create(department).Error
	return err
}

// DeleteDepartment 删除院系信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (departmentService *DepartmentService) DeleteDepartment(department InfoQuire.Department) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&department).Error
	return err
}

// DeleteDepartmentByIds 批量删除院系信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (departmentService *DepartmentService) DeleteDepartmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]InfoQuire.Department{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDepartment 更新院系信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (departmentService *DepartmentService) UpdateDepartment(department InfoQuire.Department) (err error) {
	err = global.GVA_DB.Save(&department).Error
	return err
}

// GetDepartment 根据id获取院系信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (departmentService *DepartmentService) GetDepartment(id uint) (department InfoQuire.Department, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&department).Error
	return
}

// GetDepartmentInfoList 分页获取院系信息表记录
// Author [piexlmax](https://github.com/piexlmax)
func (departmentService *DepartmentService) GetDepartmentInfoList(info InfoQuireReq.DepartmentSearch) (list []InfoQuire.Department, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&InfoQuire.Department{})
	var departments []InfoQuire.Department
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
	orderMap["dno"] = true
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

	err = db.Find(&departments).Error
	return departments, total, err
}
