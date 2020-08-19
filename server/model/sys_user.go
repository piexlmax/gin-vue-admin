package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	gorm.Model
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:'用户UUID'"`
	Username    string       `json:"userName" gorm:"comment:'用户登录名'"`
	Password    string       `json:"-"  gorm:"comment:'用户登录密码'"`
	NickName    string       `json:"nickName" gorm:"default:'系统用户';comment:'用户昵称'" `
	HeaderImg   string       `json:"headerImg" gorm:"default:'http://qmplusimg.henrongyi.top/head.png';comment:'用户头像'"`
	Authority   SysAuthority `json:"authority" gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId;comment:'用户角色'"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:'用户角色ID'"`
	IsOpen      bool         `json:"isopen" gorm:"default:false;comment:'是否开启多因子认证',type:boolean,column:isopen"`
	Secret      string       `json:"secret" gorm:"comment:'用户多因子认证秘钥'"`
}
