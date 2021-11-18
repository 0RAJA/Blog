package main

import (
	"Blog/global"
	"Blog/internal/model"
	"Blog/internal/routers"
	"Blog/pkg/logger"
	"Blog/pkg/setting"
	"Blog/pkg/tracer"
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	port    string
	runMode string
	config  string
)

func init() {
	setupFlag()
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/0RAJA/Blog
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	r := routers.NewRouter()
	//设置允许读取/写入的最大时间、请求头的最大字节数
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        r,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Infof("%s: go-programming-tour-book/%s", "raja", "blog-service")
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()
	//退出通知
	quit := make(chan os.Signal)
	//等待退出通知
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Infof("ShutDown Server...")
	//给几秒完成剩余任务
	ctx, cancel := context.WithTimeout(context.Background(), global.AppSetting.DefaultContextTimeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil { //优雅退出
		global.Logger.Infof("Server forced to ShutDown:", err)
	}
	global.Logger.Infof("Server exiting")
}

//设置配置文件
func setupSetting() error {
	newSetting, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.AppSetting.DefaultContextTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
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
	//是我们使用了 lumberjack 作为日志库的 io.Writer，
	//并且设置日志文件所允许的最大占用空间为 600MB、日志文件最大生存周期为 10 天，并且设置日志文件名的时间格式为本地时间。
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupFlag() {
	//命令行参数绑定
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.Parse()
	if config == "" {
		config = "configs/"
	}
}

func SetupTracer() error {
	//将先前编写的trace注入到全局中,方便之后的使用
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog-service",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
