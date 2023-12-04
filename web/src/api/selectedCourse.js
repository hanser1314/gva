import service from '@/utils/request'

// @Tags SelectedCourse
// @Summary 创建已选课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SelectedCourse true "创建已选课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /selectedCourse/createSelectedCourse [post]
export const createSelectedCourse = (data) => {
  return service({
    url: '/selectedCourse/createSelectedCourse',
    method: 'post',
    data
  })
}

// @Tags SelectedCourse
// @Summary 删除已选课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SelectedCourse true "删除已选课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /selectedCourse/deleteSelectedCourse [delete]
export const deleteSelectedCourse = (data) => {
  return service({
    url: '/selectedCourse/deleteSelectedCourse',
    method: 'delete',
    data
  })
}

// @Tags SelectedCourse
// @Summary 批量删除已选课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除已选课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /selectedCourse/deleteSelectedCourse [delete]
export const deleteSelectedCourseByIds = (data) => {
  return service({
    url: '/selectedCourse/deleteSelectedCourseByIds',
    method: 'delete',
    data
  })
}

// @Tags SelectedCourse
// @Summary 更新已选课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SelectedCourse true "更新已选课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /selectedCourse/updateSelectedCourse [put]
export const updateSelectedCourse = (data) => {
  return service({
    url: '/selectedCourse/updateSelectedCourse',
    method: 'put',
    data
  })
}

// @Tags SelectedCourse
// @Summary 用id查询已选课程表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SelectedCourse true "用id查询已选课程表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /selectedCourse/findSelectedCourse [get]
export const findSelectedCourse = (params) => {
  return service({
    url: '/selectedCourse/findSelectedCourse',
    method: 'get',
    params
  })
}

// @Tags SelectedCourse
// @Summary 分页获取已选课程表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取已选课程表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /selectedCourse/getSelectedCourseList [get]
export const getSelectedCourseList = (params) => {
  return service({
    url: '/selectedCourse/getSelectedCourseList',
    method: 'get',
    params
  })
}

// export const getSelectCourseListByUsername = (params) => {
//   return service({
//     url: '/selectedCourse/getSelectCourseListByUsername',
//     method: 'get',
//     params
//   })
// }
