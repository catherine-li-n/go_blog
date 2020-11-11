package main

import (
	"github.com/catherine.li/go_blog/global"
	"github.com/catherine.li/go_blog/internal/model"
	"github.com/catherine.li/go_blog/internal/routers"
	"github.com/catherine.li/go_blog/pkg/logger"
	"github.com/catherine.li/go_blog/pkg/setting"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	_ "strings"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	//err = setupValidator()
	//if err != nil {
	//	log.Fatalf("init.setupValidator err: %v", err)
	//}
}

//func setupValidator() error {
//	global.Validator = validator.NewCustomValidator()
//	global.Validator.Engine()
//	binding.Validator = global.Validator
//
//	return nil
//}

func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" +
		global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	// 创建一个默认的路由引擎
	//r := gin.Default()
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	//For Test global.Logger.Infof("%s: go-programming-book-tour/%s", "eddycjy", "blog-service")
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	//r.GET("/hello", func(c *gin.Context) {
	//	// c.JSON：返回JSON格式的数据
	//	c.JSON(200, gin.H{
	//		"message": "Hello world!",
	//	})
	//})
	//// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	//
	//r.Run()
}
