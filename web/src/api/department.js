import service from '@/utils/request'

// @Tags Department
// @Summary 创建院系信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Department true "创建院系信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /department/createDepartment [post]
export const createDepartment = (data) => {
  return service({
    url: '/department/createDepartment',
    method: 'post',
    data
  })
}

// @Tags Department
// @Summary 删除院系信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Department true "删除院系信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /department/deleteDepartment [delete]
export const deleteDepartment = (data) => {
  return service({
    url: '/department/deleteDepartment',
    method: 'delete',
    data
  })
}

// @Tags Department
// @Summary 批量删除院系信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除院系信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /department/deleteDepartment [delete]
export const deleteDepartmentByIds = (data) => {
  return service({
    url: '/department/deleteDepartmentByIds',
    method: 'delete',
    data
  })
}

// @Tags Department
// @Summary 更新院系信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Department true "更新院系信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /department/updateDepartment [put]
export const updateDepartment = (data) => {
  return service({
    url: '/department/updateDepartment',
    method: 'put',
    data
  })
}

// @Tags Department
// @Summary 用id查询院系信息表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Department true "用id查询院系信息表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /department/findDepartment [get]
export const findDepartment = (params) => {
  return service({
    url: '/department/findDepartment',
    method: 'get',
    params
  })
}

// @Tags Department
// @Summary 分页获取院系信息表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取院系信息表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /department/getDepartmentList [get]
export const getDepartmentList = (params) => {
  return service({
    url: '/department/getDepartmentList',
    method: 'get',
    params
  })
}
