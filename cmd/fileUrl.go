package cmd

import (
	"fmt"
	"os"
	"sqltostruct/config"

	"gitee.com/nicole-go-libs/print-colors/color_print"
)

func InitFileUrl() {
	if err := os.MkdirAll(config.Conf.SaveFilePwd, 0755); err != nil {
		fmt.Println(color_print.Red(fmt.Sprintf("初始化文件路径创建失败，失败原因：\n%s", err)))
		panic(err)
	}
}
