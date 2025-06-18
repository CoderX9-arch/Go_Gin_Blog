package core

import (
	"blog/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"log"
)

const ConfigFile = "settings.yaml"

func InitConf() {

	//c := &global.Config
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("read settings.yaml error:%s", err))
	}
	err = yaml.Unmarshal(yamlConf, &global.Config)
	if err != nil {
		log.Fatalf("init configfile fail:%s", err)
	}
	log.Println("init configfile success!")
	//fmt.Println(c)
}

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		global.Log.Error(err)
		return err
	}
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		global.Log.Error(err)
		return err
	}
	global.Log.Info("修改成功")
	return nil
}
