package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// 如果需要多实例部署，可使用 Redis 来存储图片验证码。需初始化 Redis
// 使用 `var store = NewDefaultRedisStore() ` 替换下面的 `var store = base64Captcha.DefaultMemStore` 即可
//
// var store = NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]
func Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.GVA_LOG.Error("验证码获取失败!", zap.Any("err", err))
		response.FailWithMessage("验证码获取失败", c)
	} else {
		response.OkWithDetailed(response.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}

func NewDefaultRedisStore() base64Captcha.Store {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_KEY_",
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
}

func (rs *RedisStore) Set(id string, value string) {
	err := global.GVA_REDIS.Set(rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		global.GVA_LOG.Error("RedisStoreSetError!", zap.Any("err", err))
	}
}

func (rs *RedisStore) Get(id string, clear bool) string {
	key := rs.PreKey + id
	val, err := global.GVA_REDIS.Get(key).Result()
	if err != nil {
		global.GVA_LOG.Error("RedisStoreGetError!", zap.Any("err", err))
		return ""
	}
	if clear {
		err := global.GVA_REDIS.Del(key).Err()
		if err != nil {
			global.GVA_LOG.Error("RedisStoreClearError!", zap.Any("err", err))
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
