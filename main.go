package main

import (
	"blog-service/global"
	"blog-service/internal/model"
	"blog-service/internal/routers"
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

// 初始化配置
func setupSetting() error {
	set, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = set.ReadSection("Server",&global.ServerSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("App",&global.AppSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("Database",&global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

// 初始化数据库引擎
func setupDBEngine() error {
	var err error
	global.DBEngine,err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

// 初始化日志
func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
			Filename: global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
			MaxSize: 600,
			MaxAge: 10,
			LocalTime: true,
	},"",log.Lshortfile).WithCaller(2)
	return nil
}

// 总初始化 会在main方法前自动执行
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v",err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v",err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v",err)
	}
}


/*
@title 博客系统
@version 1.0
@description CK‘s Go练手项目
 */
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":8080",
		Handler: router,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
