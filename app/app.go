package app

import (
	"flag"
	"log"
	"os"
	"runtime"

	"github.com/mocheer/golib/tray"
)

func DoMain() {
	filePath := flag.String("f", "./app.json", "jsonFile to config")
	flag.Parse()
	Run(*filePath)
}

func Run(filePath string) {
	tray := tray.New()

	appConfig := &AppConfig{File: filePath}
	err := appConfig.Init()
	if err != nil {
		log.Println("读取配置信息错误：", err.Error())
		os.Exit(0)
	}
	tray.AddLeftMenuItems(appConfig.LeftMenu)
	tray.AddRightMenuItems(appConfig.RightMenu)
	trayConfig := appConfig.TrayConfig
	err = tray.Show(trayConfig.Ico, trayConfig.Name)
	if err != nil {
		log.Println("初始化错误：", err.Error())
		os.Exit(0)
	}
	log.Println("初始化成功，开始运行！")
	runtime.LockOSThread()
	tray.Run()
	runtime.UnlockOSThread()
}
