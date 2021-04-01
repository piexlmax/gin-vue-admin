import service from '@/utils/request'

// @Tags ExaSensitiveWord
// @Summary 创建ExaSensitiveWord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaSensitiveWord true "创建ExaSensitiveWord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exaSensitiveWord/createExaSensitiveWord [post]
export const createExaSensitiveWord = (data) => {
     return service({
         url: "/exaSensitiveWord/createExaSensitiveWord",
         method: 'post',
         data
     })
 }


// @Tags ExaSensitiveWord
// @Summary 删除ExaSensitiveWord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaSensitiveWord true "删除ExaSensitiveWord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exaSensitiveWord/deleteExaSensitiveWord [delete]
 export const deleteExaSensitiveWord = (data) => {
     return service({
         url: "/exaSensitiveWord/deleteExaSensitiveWord",
         method: 'delete',
         data
     })
 }

// @Tags ExaSensitiveWord
// @Summary 删除ExaSensitiveWord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExaSensitiveWord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exaSensitiveWord/deleteExaSensitiveWord [delete]
 export const deleteExaSensitiveWordByIds = (data) => {
     return service({
         url: "/exaSensitiveWord/deleteExaSensitiveWordByIds",
         method: 'delete',
         data
     })
 }

// @Tags ExaSensitiveWord
// @Summary 更新ExaSensitiveWord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaSensitiveWord true "更新ExaSensitiveWord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exaSensitiveWord/updateExaSensitiveWord [put]
 export const updateExaSensitiveWord = (data) => {
     return service({
         url: "/exaSensitiveWord/updateExaSensitiveWord",
         method: 'put',
         data
     })
 }


// @Tags ExaSensitiveWord
// @Summary 用id查询ExaSensitiveWord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaSensitiveWord true "用id查询ExaSensitiveWord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exaSensitiveWord/findExaSensitiveWord [get]
 export const findExaSensitiveWord = (params) => {
     return service({
         url: "/exaSensitiveWord/findExaSensitiveWord",
         method: 'get',
         params
     })
 }


// @Tags ExaSensitiveWord
// @Summary 分页获取ExaSensitiveWord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取ExaSensitiveWord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exaSensitiveWord/getExaSensitiveWordList [get]
 export const getExaSensitiveWordList = (params) => {
     return service({
         url: "/exaSensitiveWord/getExaSensitiveWordList",
         method: 'get',
         params
     })
 }