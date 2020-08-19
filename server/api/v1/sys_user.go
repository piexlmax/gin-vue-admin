package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"mime/multipart"
	"time"
)

// @Tags Base
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.SysUser true "用户注册接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /base/register [post]
func Register(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	UserVerify := utils.Rules{
		"Username":    {utils.NotEmpty()},
		"NickName":    {utils.NotEmpty()},
		"Password":    {utils.NotEmpty()},
		"AuthorityId": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(R, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	user := &model.SysUser{Username: R.Username, NickName: R.NickName, Password: R.Password, HeaderImg: R.HeaderImg, AuthorityId: R.AuthorityId}
	err, userReturn := service.Register(*user)
	if err != nil {
		response.FailWithDetailed(response.ERROR, resp.SysUserResponse{User: userReturn}, fmt.Sprintf("%v", err), c)
	} else {
		response.OkDetailed(resp.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var L request.RegisterAndLoginStruct
	_ = c.ShouldBindJSON(&L)
	UserVerify := utils.Rules{
		"CaptchaId": {utils.NotEmpty()},
		"Captcha":   {utils.NotEmpty()},
		"Username":  {utils.NotEmpty()},
		"Password":  {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(L, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	if store.Verify(L.CaptchaId, L.Captcha, true) {
		U := &model.SysUser{Username: L.Username, Password: L.Password}
		if err, user := service.Login(U); err != nil {
			response.FailWithMessage(fmt.Sprintf("用户名密码错误或%v", err), c)
		} else {
			tokenNext(c, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}

}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user model.SysUser) {
	j := &middleware.JWT{
		SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey), // 唯一签名
	}
	clams := request.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 一周
			Issuer:    "qmPlus",                       // 签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithData(resp.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
		return
	}
	var loginJwt model.JwtBlacklist
	loginJwt.Jwt = token
	err, jwtStr := service.GetRedisJWT(user.Username)
	if err == redis.Nil {
		if err := service.SetRedisJWT(loginJwt, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithData(resp.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
	} else if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := service.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := service.SetRedisJWT(loginJwt, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithData(resp.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, c)
	}
}

// @Tags SysUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangePasswordStruct true "用户修改密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
func ChangePassword(c *gin.Context) {
	var params request.ChangePasswordStruct
	_ = c.ShouldBindJSON(&params)
	UserVerify := utils.Rules{
		"Username":    {utils.NotEmpty()},
		"Password":    {utils.NotEmpty()},
		"NewPassword": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(params, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	U := &model.SysUser{Username: params.Username, Password: params.Password}
	if err, _ := service.ChangePassword(U, params.NewPassword); err != nil {
		response.FailWithMessage("修改失败，请检查用户名密码", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

type UserHeaderImg struct {
	HeaderImg multipart.File `json:"headerImg"`
}

// @Tags SysUser
// @Summary 用户上传头像
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param headerImg formData file true "用户上传头像"
// @Param username formData string true "用户上传头像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /user/uploadHeaderImg [post]
func UploadHeaderImg(c *gin.Context) {
	claims, _ := c.Get("claims")
	// 获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	waitUse := claims.(*request.CustomClaims)
	uuid := waitUse.UUID
	_, header, err := c.Request.FormFile("headerImg")
	// 便于找到用户 以后从jwt中取
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("上传文件失败，%v", err), c)
	} else {
		// 文件上传后拿到文件路径
		err, filePath, _ := utils.Upload(header)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("接收返回值失败，%v", err), c)
		} else {
			// 修改数据库后得到修改后的user并且返回供前端使用
			err, user := service.UploadHeaderImg(uuid, filePath)
			if err != nil {
				response.FailWithMessage(fmt.Sprintf("修改数据库链接失败，%v", err), c)
			} else {
				response.OkWithData(resp.SysUserResponse{User: *user}, c)
			}
		}
	}
}

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, list, total := service.GetUserInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}

// @Tags SysUser
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "设置用户权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func SetUserAuthority(c *gin.Context) {
	var sua request.SetUserAuth
	_ = c.ShouldBindJSON(&sua)
	UserVerify := utils.Rules{
		"UUID":        {utils.NotEmpty()},
		"AuthorityId": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(sua, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	err := service.SetUserAuthority(sua.UUID, sua.AuthorityId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("修改失败，%v", err), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	IdVerifyErr := utils.Verify(reqId, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err := service.DeleteUser(reqId.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}


// @Tags SysUser
// @Summary 检查用户是否开启二次认证
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "检查用户是否开启二次认证"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"用户已开启二次认证"}"
// @Router /user/checkusermfaisopen [post]
func CheckUserMFAIsOpen(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	IdVerifyErr := utils.Verify(reqId, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err,isopen := service.CheckMfaIsOpen(reqId.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("检查用户是否开启二次认证失败，%v", err), c)
	} else {
		response.OkWithData(map[string]bool{"isopen":isopen,"success":true},c)
	}
}


// @Tags SysUser
// @Summary 获取用户是否开启二次认证及相关信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "获取用户是否开启二次认证及相关信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取用户是否开启二次认证及相关信息"}"
// @Router /user/getusermfainfo [post]
func GetUserMFAInfo(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	IdVerifyErr := utils.Verify(reqId, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err,secret := service.GetSecret(reqId.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取用户多因子认证相关信息失败，%v", err), c)
		return
	}
	accountname := ""
	isOpen := false
	if secret == "" {
		usersecret,err := utils.GenerateSecret()
		if err!= nil{
			response.FailWithMessage(fmt.Sprintf("生成动态秘钥失败，%v", err), c)
			return
		}
		//if user has no secret then update to database
		err1,_ := service.UpdateSecret(reqId.Id,usersecret)
		if err1 != nil{
			response.FailWithMessage(fmt.Sprintf("更新用户动态秘钥至数据库失败，%v", err1), c)
			return
		}
		secret = usersecret
	}
	err, userinfo := service.GetUserInfoById(reqId.Id)
	if err!= nil{
		response.FailWithMessage(fmt.Sprintf("获取用户信息失败，%v", err), c)
		return
	}
	accountname = userinfo.Username
	isOpen = userinfo.IsOpen
	err,qrcode := utils.GenerateMFAQrcode(secret,accountname)
	if err!= nil{
		response.FailWithMessage(fmt.Sprintf("生成google多因子认证二维码失败，%v", err), c)
		return
	}
	response.OkWithData(request.UserMFAInfo{
		Isopen:isOpen,
		QrCode:qrcode,
		AccountName:accountname,
		Secret:secret,
	},c)
}

// @Tags SysUser
// @Summary 绑定用户二次认证信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "绑定多因子认证"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取用户是否开启二次认证及相关信息"}"
// @Router /user/bindmfa [post]
func BindMfa(c *gin.Context) {
	var reqId request.UserBindMfainfo
	_ = c.ShouldBindJSON(&reqId)
	VerifyRule := utils.Rules{
		"Id":        {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	IdVerifyErr := utils.Verify(reqId, VerifyRule)
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err,secret := service.GetSecret(reqId.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取用户多因子认证相关信息失败，%v", err), c)
		return
	}
	err, userinfo := service.GetUserInfoById(reqId.Id)
	if err!= nil{
		response.FailWithMessage(fmt.Sprintf("获取用户信息失败，%v", err), c)
		return
	}
	isOpen := userinfo.IsOpen
	authkey := reqId.AuthKey
	if !isOpen {
		err,isauth := utils.AuthenticateMFA(secret,authkey)
		if err!= nil{
			response.FailWithMessage(fmt.Sprintf("多因子认证调用时失败，%v", err), c)
			return
		}
		if isauth{
			err, _ := service.BindMFA(userinfo)
			if err != nil{
				response.OkWithData(map[string]bool{
					"success": false,
				},c)
				return
			}
			response.OkWithData(map[string]bool{
				"success": true,
			},c)
			return
		}else {
			response.OkWithData(map[string]bool{
				"success": false,
			},c)
			return
		}
	} else {
		response.FailWithMessage(fmt.Sprintf("绑定用户多因子认证时失败,用户已经绑定多因子登录认证"), c)
		return
	}
}


// @Tags SysUser
// @Summary 解除用户二次认证信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "解除绑定多因子认证"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取用户是否开启二次认证及相关信息"}"
// @Router /user/unbindmfa [post]
func UnBindMfa(c *gin.Context) {
	var reqId request.UserBindMfainfo
	_ = c.ShouldBindJSON(&reqId)
	VerifyRule := utils.Rules{
		"Id":        {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	IdVerifyErr := utils.Verify(reqId, VerifyRule)
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err,secret := service.GetSecret(reqId.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取用户多因子认证相关信息失败，%v", err), c)
		return
	}
	err, userinfo := service.GetUserInfoById(reqId.Id)
	if err!= nil{
		response.FailWithMessage(fmt.Sprintf("获取用户信息失败，%v", err), c)
		return
	}
	isOpen := userinfo.IsOpen
	authkey := reqId.AuthKey
	if isOpen {
		err,isauth := utils.AuthenticateMFA(secret,authkey)
		if err!= nil{
			response.FailWithMessage(fmt.Sprintf("多因子认证调用时失败，%v", err), c)
			return
		}
		if isauth{
			err, _ := service.CloseMFA(userinfo)
			if err != nil{
				response.OkWithData(map[string]bool{
					"success": false,
				},c)
				return
			}
			response.OkWithData(map[string]bool{
				"success": true,
			},c)
			return
		}else {
			response.OkWithData(map[string]bool{
				"success": false,
			},c)
			return
		}
	} else {
		response.FailWithMessage(fmt.Sprintf("解除用户多因子认证时失败,用户已经解除多因子登录认证"), c)
		return
	}
}