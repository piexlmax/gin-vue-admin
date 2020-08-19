import service from '@/utils/request'

// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data) => {
    return service({
        url: "/base/login",
        method: 'post',
        data: data
    })
}

// @Summary 获取验证码
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/captcha [post]
export const captcha = (data) => {
    return service({
        url: "/base/captcha",
        method: 'post',
        data: data
    })
}

// @Summary 用户注册
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/resige [post]
export const register = (data) => {
        return service({
            url: "/base/register",
            method: 'post',
            data: data
        })
    }
    // @Summary 修改密码
    // @Produce  application/json
    // @Param data body {username:"string",password:"string",newPassword:"string"}
    // @Router /user/changePassword [post]
export const changePassword = (data) => {
        return service({
            url: "/user/changePassword",
            method: 'post',
            data: data
        })
    }
    // @Tags User
    // @Summary 分页获取用户列表
    // @Security ApiKeyAuth
    // @accept application/json
    // @Produce application/json
    // @Param data body modelInterface.PageInfo true "分页获取用户列表"
    // @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
    // @Router /user/getUserList [post]
export const getUserList = (data) => {
    return service({
        url: "/user/getUserList",
        method: 'post',
        data: data
    })
}


// @Tags User
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.SetUserAuth true "设置用户权限"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
export const setUserAuthority = (data) => {
    return service({
        url: "/user/setUserAuthority",
        method: 'post',
        data: data
    })
}


// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/deleteUser [delete]
export const deleteUser = (data) => {
    return service({
        url: "/user/deleteUser",
        method: 'delete',
        data: data
    })
}

// @Tags SysUser
// @Summary 获取用户多因子认证信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "获取用户多因子认证信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取用户多因子认证信息成功"}"
// @Router /user/getusermfainfo [post]
export const getusermfainfo = (data) => {
    return service({
        url: "/user/getusermfainfo",
        method: 'post',
        data: data
    })
}


// @Tags SysUser
// @Summary 绑定用户多因子认证
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "绑定多因子认证"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"绑定多因子认证信息成功"}"
// @Router /user/bindmfa [post]
export const bindmfa = (data) => {
    return service({
        url: "/user/bindmfa",
        method: 'post',
        data: data
    })
}


// @Tags SysUser
// @Summary 解除用户多因子认证
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "解除多因子认证"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"解除多因子认证信息成功"}"
// @Router /user/unbindmfa [post]
export const unbindmfa = (data) => {
    return service({
        url: "/user/unbindmfa",
        method: 'post',
        data: data
    })
}
