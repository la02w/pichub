package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	LOCAL_BASE_FOLDER     string
	TENCENT_COS_URL       string
	TENCENT_COS_SECRETID  string
	TENCENT_COS_SECRETKEY string
	PICK_SERVICE          string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	PICK_SERVICE = file.Section("base").Key("PICK_SERVICE").MustString("local")
	LoadLocal(file)
	LoadTencent(file)
}

func LoadLocal(file *ini.File) {
	LOCAL_BASE_FOLDER = file.Section("local").Key("LOCAL_BASE_FOLDER").MustString("./upload/")
}
func LoadTencent(file *ini.File) {
	TENCENT_COS_URL = file.Section("tencent").Key("TENCENT_COS_URL").MustString("")
	TENCENT_COS_SECRETID = file.Section("tencent").Key("TENCENT_COS_SECRETID").MustString("")
	TENCENT_COS_SECRETKEY = file.Section("tencent").Key("TENCENT_COS_SECRETKEY").MustString("")
}
