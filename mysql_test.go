package main

import (
	"sqltostruct/config"
	"sqltostruct/global"
	"sqltostruct/internal/service"
	"sqltostruct/internal/templates"
	"testing"
)

func TestGetTable(t *testing.T) {

}

func TestTempMain(t *testing.T) {
	templates.TempMain()
}

//生成一个表的结构体
func TestTableOne(t *testing.T) {
	config.Conf.IsAllDataBase = false
	config.Conf.TableName = "eb_user_invite_log"

	service.SqlTwoStruct()

	// 如果发生错误则关闭数据库连接
	defer func() {
		_ = global.DBEngine.Close()
	}()
}
