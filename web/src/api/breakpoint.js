import service from '@/utils/request'

// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body sysModel.SysAuthority true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/setDataAuthority [post]

export const findFile = (params) => {
  return service({
    url: '/fileUploadAndDownload/findFile',
    method: 'get',
    params
  })
}

export const breakpointContinueFinish = (params) => {
  return service({
    url: '/fileUploadAndDownload/breakpointContinueFinish',
    method: 'put',
    params
  })
}

export const removeChunk = (data, params) => {
  return service({
    url: '/fileUploadAndDownload/removeChunk',
    method: 'delete',
    data,
    params
  })
}
