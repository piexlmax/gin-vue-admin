package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	uuid "github.com/satori/go.uuid"
)

// @title    Register
// @description   register, 用户注册
// @auth                     （2020/04/05  20:22）
// @param     u               model.SysUser
// @return    err             error
// @return    userInter       *SysUser

func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	// 判断用户名是否注册
	notRegister := global.GVA_DB.Where("username = ?", u.Username).First(&user).RecordNotFound()
	// notRegister为false表明读取到了 不能注册
	if !notRegister {
		return errors.New("用户名已注册"), userInter
	} else {
		// 否则 附加uuid 密码md5简单加密 注册
		u.Password = utils.MD5V([]byte(u.Password))
		u.UUID = uuid.NewV4()
		err = global.GVA_DB.Create(&u).Error
	}
	return err, u
}

// @title    Login
// @description   login, 用户登录
// @auth                     （2020/04/05  20:22）
// @param     u               *model.SysUser
// @return    err             error
// @return    userInter       *SysUser

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

// @title    ChangePassword
// @description   change the password of a certain user, 修改用户密码
// @auth                     （2020/04/05  20:22）
// @param     u               *model.SysUser
// @param     newPassword     string
// @return    err             error
// @return    userInter       *SysUser

func ChangePassword(u *model.SysUser, newPassword string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

// @title    GetInfoList
// @description   get user list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysUser{})
	var userList []model.SysUser
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	return err, userList, total
}

// @title    SetUserAuthority
// @description   set the authority of a certain user, 设置一个用户的权限
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     authorityId     string
// @return    err             error

func SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&model.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

// @title    SetUserAuthority
// @description   set the authority of a certain user, 设置一个用户的权限
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     authorityId     string
// @return    err             error

func DeleteUser(id float64) (err error) {
	var user model.SysUser
	err = global.GVA_DB.Where("id = ?", id).Delete(&user).Error
	return err
}

// @title    UploadHeaderImg
// @description   upload avatar, 用户头像上传更新地址
// @auth                     （2020/04/05  20:22）
// @param     uuid            UUID
// @param     filePath        string
// @return    err             error
// @return    userInter       *SysUser

func UploadHeaderImg(uuid uuid.UUID, filePath string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&user).Update("header_img", filePath).First(&user).Error
	return err, &user
}

// @title    BindMFA
// @description   bind google mfa for a certain user, 绑定用户多因子认证
// @auth                     （2020/07/09  11:21）
// @param     u               *model.SysUser
// @return    err             error
// @return    userInter       *SysUser

func BindMFA(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	err = global.GVA_DB.Where("username = ?", u.Username).First(&user).Update("IsOpen", true).First(&user).Error
	return err, &user
}

// @title    CloseMFA
// @description   bind google mfa for a certain user, 关闭用户多因子认证
// @auth                     （2020/07/09  11:25）
// @param     u               *model.SysUser
// @return    err             error
// @return    userInter       *SysUser

func CloseMFA(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	err = global.GVA_DB.Where("username = ?", u.Username).First(&user).Update("IsOpen", false).First(&user).Error
	return err, &user
}

// @title    GetSecret
// @description   get google mfa secret for a certain user, 获取用户多因子认证秘钥
// @auth                     （2020/07/09  14:28）
// @param     u               *model.SysUser
// @return    err             error
// @return    secret          string

func GetSecret(uid float64) (err error, secret string) {
	var user model.SysUser
	err = global.GVA_DB.Where("id = ?", uid).First(&user).Error
	return err, user.Secret
}

// @title    UpdateSecret
// @description   update google mfa secret for a certain user, 更新用户多因子认证秘钥信息
// @auth                     （2020/07/10  10:37）
// @param     uid               float64
// @param     secret            string
// @return    err             error

func UpdateSecret(uid float64,secret string) (err error, userInfo *model.SysUser) {
	var user model.SysUser
	err = global.GVA_DB.Where("id = ?", uid).First(&user).Update("secret",secret).Error
	return err,&user
}

// @title    CheckMfaIsOpen
// @description   get google mfa secret for a certain user, 获取用户是否开启多因子认证
// @auth                     （2020/07/09  21:34）
// @param     u               *model.SysUser
// @return    err             error
// @return    isopen          boolean

func CheckMfaIsOpen(uid float64) (err error, isopen bool) {
	var user model.SysUser
	err = global.GVA_DB.Where("id = ?", uid).First(&user).Error
	return err, user.IsOpen
}

// @title    GetUserInfoById
// @description   get user information by id, 根据用户id获取用户信息
// @auth                     （2020/07/12  12:00）
// @param     u               *model.SysUser
// @return    err             error
// @return    userInfo       *SysUser

func GetUserInfoById(uid float64) (err error, userInfo *model.SysUser) {
	var user model.SysUser
	err = global.GVA_DB.Where("id = ?", uid).First(&user).Error
	return err, &user
}