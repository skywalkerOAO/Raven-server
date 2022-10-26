package config

import (
	"fmt"
	"github.com/skywalkerOAO/Gotos"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var ConfigInfo Cfg

type Cfg struct {
	Mode       string `yaml:"mode"`
	ServerPort string `yaml:"server-port"`
}

func ReadOptions() {
	yamlPath := getOptionsDir()
	//根据路径读文件内容
	content, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	//yaml解析，并赋值给ConfigInfo
	err = yaml.Unmarshal(content, &ConfigInfo)
}
func getOptionsDir() string {
	ginMode := "release"
	//返回绝对路径 filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir := GetCurrentDirectory()
	Exists := Gotos.PathExists(dir + "/bin")
	if Exists {
		ginMode = "debug"
	}
	if ginMode == "debug" {
		dir = fmt.Sprintf("%s/bin/config/config.yaml", dir)
		return dir
	}
	if ginMode == "release" {
		dir = fmt.Sprintf("%s/config/config.yaml", dir)
		return dir
	}
	if ginMode == "test" {
		return dir
	}
	return dir
}

func GetCurrentDirectory() string {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1)
}
