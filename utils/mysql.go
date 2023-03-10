package utils

import (
	"mtv/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

// 定义数据库连接配置
type DbConfig struct {
	Debug        bool
	Alias        string
	Driver       string
	DSN          string
	MaxOpenConns int //最大连接数
	MaxIdleConns int //最大空闲连接数
}

func loadDB() DbConfig {
	dbdebug, _ := config.Bool("mysql::dbdebug")
	dbalias, _ := config.String("mysql::dbalias")
	dbdriver, _ := config.String("mysql::dbdriver")
	dbhost, _ := config.String("mysql::dbhost")
	dbport, _ := config.String("mysql::dbport")
	dbname, _ := config.String("mysql::dbname")
	dbuser, _ := config.String("mysql::dbuser")
	dbpassword, _ := config.String("mysql::dbpassword")
	dbcharset, _ := config.String("mysql::dbcharset")
	dbloc, _ := config.String("mysql::dbloc")

	maxIdleConns, _ := config.Int("mysql::maxIdleConns")
	maxOpenConns, _ := config.Int("mysql::maxOpenConns")

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=" + dbcharset + "&loc=" + dbloc

	dbConfig := DbConfig{}
	dbConfig.Debug = dbdebug
	dbConfig.Alias = dbalias
	dbConfig.Driver = dbdriver
	dbConfig.DSN = dsn
	dbConfig.MaxOpenConns = maxOpenConns
	dbConfig.MaxIdleConns = maxIdleConns
	logs.Info(dbConfig)
	return dbConfig
}

// 初始化
func InitMySQL() {
	//加载配置
	dbconfig := loadDB()
	//开启调试
	orm.Debug = dbconfig.Debug
	//注册驱动
	orm.RegisterDriver(dbconfig.Driver, orm.DRMySQL)
	//注册数据库
	err := orm.RegisterDataBase(dbconfig.Alias, dbconfig.Driver, dbconfig.DSN, orm.MaxOpenConnections(dbconfig.MaxOpenConns), orm.MaxIdleConnections(dbconfig.MaxIdleConns))
	if err != nil {
		logs.Error(err)
	}

	//注册模型
	orm.RegisterModel(new(models.Chat))
	orm.RegisterModel(new(models.NostrRelay))
	orm.RegisterModel(new(models.Question))
	orm.RegisterModel(new(models.QuestionTmp))
	orm.RegisterModel(new(models.User))

	//创建数据表
	orm.RunSyncdb("default", false, true)
}
