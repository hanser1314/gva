import service from '@/utils/request'

// @Tags Student
// @Summary 创建student表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Student true "创建student表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /student/createStudent [post]
export const createStudent = (data) => {
  return service({
    url: '/student/createStudent',
    method: 'post',
    data
  })
}

// @Tags Student
// @Summary 删除student表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Student true "删除student表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /student/deleteStudent [delete]
export const deleteStudent = (data) => {
  return service({
    url: '/student/deleteStudent',
    method: 'delete',
    data
  })
}

// @Tags Student
// @Summary 批量删除student表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除student表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /student/deleteStudent [delete]
export const deleteStudentByIds = (data) => {
  return service({
    url: '/student/deleteStudentByIds',
    method: 'delete',
    data
  })
}

// @Tags Student
// @Summary 更新student表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Student true "更新student表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /student/updateStudent [put]
export const updateStudent = (data) => {
  return service({
    url: '/student/updateStudent',
    method: 'put',
    data
  })
}

// @Tags Student
// @Summary 用id查询student表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Student true "用id查询student表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /student/findStudent [get]
export const findStudent = (params) => {
  return service({
    url: '/student/findStudent',
    method: 'get',
    params
  })
}

// @Tags Student
// @Summary 分页获取student表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取student表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/getStudentList [get]
export const getStudentList = (params) => {
  return service({
    url: '/student/getStudentList',
    method: 'get',
    params
  })
}
// 根据学生姓名和年龄获取student表列表
export const getStudentByName = (params) => {
  return service({
    url: '/student/getStudentByName',
    method: 'get',
    params
  })
}

// 根据学生学号获取student表列表
export const getStudentBySno = (params) => {
  return service({
    url: '/student/getStudentBySno',
    method: 'get',
    params
  })
}

export const getStudentByDept = (params) => {
  return service({
    url: '/student/getStudentByDept',
    method: 'get',
    params
  })
}

