// 自动生成模板ExaSensitiveWord
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type ExaSensitiveWord struct {
      global.GVA_MODEL
      Word  string `json:"word" form:"word" gorm:"column:word;comment:;type:varchar(1024);size:1024;"`
}


func (ExaSensitiveWord) TableName() string {
  return "exa_sensitive_word"
}

