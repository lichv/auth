package main

import (
	"fmt"
	"auth/utils"
	"auth/utils/setting"
)

func main() {
	setting.Setup()
	fmt.Println(utils.EncodeMD5(setting.AppSetting.SecretSalt+"123456"))
}
