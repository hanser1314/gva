import service from '@/utils/request'

// @Tags Course
// @Summary 创建课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Course true "创建课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /course/createCourse [post]
export const createCourse = (data) => {
  return service({
    url: '/course/createCourse',
    method: 'post',
    data
  })
}

// @Tags Course
// @Summary 删除课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Course true "删除课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /course/deleteCourse [delete]
export const deleteCourse = (data) => {
  return service({
    url: '/course/deleteCourse',
    method: 'delete',
    data
  })
}

// @Tags Course
// @Summary 批量删除课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /course/deleteCourse [delete]
export const deleteCourseByIds = (data) => {
  return service({
    url: '/course/deleteCourseByIds',
    method: 'delete',
    data
  })
}

// @Tags Course
// @Summary 更新课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Course true "更新课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /course/updateCourse [put]
export const updateCourse = (data) => {
  return service({
    url: '/course/updateCourse',
    method: 'put',
    data
  })
}

// @Tags Course
// @Summary 用id查询课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Course true "用id查询课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /course/findCourse [get]
export const findCourse = (params) => {
  return service({
    url: '/course/findCourse',
    method: 'get',
    params
  })
}

// @Tags Course
// @Summary 分页获取课程表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取课程表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /course/getCourseList [get]
export const getCourseList = (params) => {
  return service({
    url: '/course/getCourseList',
    method: 'get',
    params
  })
}
