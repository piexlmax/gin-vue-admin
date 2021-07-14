package captcha

import (
	"gin-vue-admin/global"
	"time"

	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

func NewDefaultRedisStore() base64Captcha.Store {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
}

func (rs *RedisStore) Set(id string, value string) {
	err := global.GVA_REDIS.Set(rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		global.GVA_LOG.Error("RedisStoreSetError!", zap.Error(err))
	}
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.GVA_REDIS.Get(key).Result()
	if err != nil {
		global.GVA_LOG.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := global.GVA_REDIS.Del(key).Err()
		if err != nil {
			global.GVA_LOG.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
