// +build windows

package core

import (
	"context"
	"gin-vue-admin/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	gServer := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT)
		<-sc

		if gServer != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := gServer.Shutdown(ctx); err != nil {
				global.GVA_LOG.Error("Server Shutdown Error: ", err.Error())
			} else {
				global.GVA_LOG.Info("Server Shutdown")
			}
		}
	}()

	return gServer
}
