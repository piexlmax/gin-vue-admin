// +build !windows

package core

import (
	"gin-vue-admin/global"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20

	s.BeforeBegin = func(add string) {
		pidMap := make(map[string]int)
		// 记录pid
		pid := syscall.Getpid()
		pidMap["pid"] = pid

		pidYaml, _ := yaml.Marshal(pidMap)
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile(dir+"/pid", pidYaml, 0664); err != nil {
			global.GVA_LOG.Sugar().Error(err)
			panic(err)
		}
		global.GVA_LOG.Sugar().Infof("Actual pid is %d", pid)
	}

	return s
}
