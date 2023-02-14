package main

import (
	"github.com/talkincode/toughradius/app"
	"github.com/talkincode/toughradius/config"
)

// 开发环境初始化数据库
/**

CREATE USER toughradius WITH PASSWORD 'toughradius'

CREATE DATABASE toughradius OWNER postgres;

GRANT ALL PRIVILEGES ON DATABASE toughradius TO toughradius;

*/

func main() {
	app.InitGlobalApplication(config.LoadConfig(""))
	app.GApp().InitDb()
}
