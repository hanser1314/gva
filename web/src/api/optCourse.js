import service from '@/utils/request'

// @Tags OptCourse
// @Summary 创建选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OptCourse true "创建选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /optCourse/createOptCourse [post]
export const createOptCourse = (data) => {
  return service({
    url: '/optCourse/createOptCourse',
    method: 'post',
    data
  })
}

// @Tags OptCourse
// @Summary 删除选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OptCourse true "删除选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /optCourse/deleteOptCourse [delete]
export const deleteOptCourse = (data) => {
  return service({
    url: '/optCourse/deleteOptCourse',
    method: 'delete',
    data
  })
}

// @Tags OptCourse
// @Summary 批量删除选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /optCourse/deleteOptCourse [delete]
export const deleteOptCourseByIds = (data) => {
  return service({
    url: '/optCourse/deleteOptCourseByIds',
    method: 'delete',
    data
  })
}

// @Tags OptCourse
// @Summary 更新选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OptCourse true "更新选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /optCourse/updateOptCourse [put]
export const updateOptCourse = (data) => {
  return service({
    url: '/optCourse/updateOptCourse',
    method: 'put',
    data
  })
}

// @Tags OptCourse
// @Summary 用id查询选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OptCourse true "用id查询选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /optCourse/findOptCourse [get]
export const findOptCourse = (params) => {
  return service({
    url: '/optCourse/findOptCourse',
    method: 'get',
    params
  })
}

// @Tags OptCourse
// @Summary 分页获取选课表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取选课表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /optCourse/getOptCourseList [get]
export const getOptCourseList = (params) => {
  return service({
    url: '/optCourse/getOptCourseList',
    method: 'get',
    params
  })
}
