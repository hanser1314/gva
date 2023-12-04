import service from '@/utils/request'

// @Tags Sct
// @Summary 创建选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sct true "创建选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sct/createSct [post]
export const createSct = (data) => {
  return service({
    url: '/sct/createSct',
    method: 'post',
    data
  })
}

// @Tags Sct
// @Summary 删除选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sct true "删除选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sct/deleteSct [delete]
export const deleteSct = (data) => {
  return service({
    url: '/sct/deleteSct',
    method: 'delete',
    data
  })
}

// @Tags Sct
// @Summary 批量删除选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sct/deleteSct [delete]
export const deleteSctByIds = (data) => {
  return service({
    url: '/sct/deleteSctByIds',
    method: 'delete',
    data
  })
}

// @Tags Sct
// @Summary 更新选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sct true "更新选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sct/updateSct [put]
export const updateSct = (data) => {
  return service({
    url: '/sct/updateSct',
    method: 'put',
    data
  })
}

// @Tags Sct
// @Summary 用id查询选课表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Sct true "用id查询选课表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sct/findSct [get]
export const findSct = (params) => {
  return service({
    url: '/sct/findSct',
    method: 'get',
    params
  })
}

// @Tags Sct
// @Summary 分页获取选课表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取选课表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sct/getSctList [get]
export const getSctList = (params) => {
  return service({
    url: '/sct/getSctList',
    method: 'get',
    params
  })
}

export const getSctListBySno = (params) => {
  return service({
    url: '/sct/getSctListBySno',
    method: 'get',
    params
  })
}

export const getSctListBySnoAndCno = (params) => {
  return service({
    url: '/sct/getSctListBySnoAndCno',
    method: 'get',
    params
  })
}

export const getSctListByTno = (params) => {
  return service({
    url: '/sct/getSctListByTno',
    method: 'get',
    params
  })
}
