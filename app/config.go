package app

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/mocheer/golib/tray"
)

type AppConfig struct {
	File       string
	TrayConfig *TrayConfig
	LeftMenu   []MenuItem
	RightMenu  []MenuItem
}

//AppConfig
type TrayConfig struct {
	Name      string `json:"name"` //托盘名称
	Ico       string `json:"ico"`  //托盘显示图标
	LeftMenu  string `json:"lm"`   //左键菜单项
	RightMenu string `json:"rm"`   //右键菜单项
}

//Init
func (this *AppConfig) Init() error {
	if this.File == "" {
		this.TrayConfig = this.GetDefaultConfig()
		return nil
	}
	err := readJSON(this.File, &this.TrayConfig)
	if err != nil {
		return err
	}
	c := this.TrayConfig
	readJSON(c.LeftMenu, &this.LeftMenu)
	readJSON(c.RightMenu, &this.RightMenu)
	return nil
}

func readJSON(path string, v interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}

// GetDefaultConfig 获取默认配置
func (this *AppConfig) GetDefaultConfig() *TrayConfig {
	return &TrayConfig{
		Name:      "gotray",
		Ico:       "app.ico",
		RightMenu: "rm.json",
		LeftMenu:  "lm.json"}
}
