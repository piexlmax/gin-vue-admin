package sensitive_word

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"sync"
)

func NewLibrary() *Library {
	return &Library{
		root: nil,
		mu:   &sync.RWMutex{},
	}
}

func LoadMySQLWords() ([]model.ExaSensitiveWord, error) {
	var words []model.ExaSensitiveWord
	if err := global.GVA_DB.Find(&words).Error; err != nil {
		return nil, err
	}
	return words, nil
}
