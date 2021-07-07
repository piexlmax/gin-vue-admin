package service

import (
	"database/sql"
	"fmt"
	"gin-vue-admin/config"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/source"
	"gin-vue-admin/utils"
	"gorm.io/driver/postgres"
	"path/filepath"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: writeConfig
//@description: 回写配置
//@param: viper *viper.Viper, mysql config.Mysql
//@return: error

func writeConfig(viper *viper.Viper, mysql config.Mysql) error {
	global.GVA_CONFIG.Mysql = mysql
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

//@author:
//@function: writePostgresConfig
//@description: 回写配置
//@param: viper *viper.Viper, postgres config.Postgres
//@return: error

func writePostgresConfig(viper *viper.Viper, postgres config.Postgres) error {
	global.GVA_CONFIG.Postgres = postgres
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

//@author:
//@function: changeDbType
//@description: 回写配置,主要修改DbType参数
//@param: viper *viper.Viper, configSystem config.System
//@return: error

func changeDbType(viper *viper.Viper, configSystem config.System) error {
	global.GVA_CONFIG.System = configSystem
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: createTable
//@description: 创建数据库(mysql)
//@param: dsn string, driver string, createSql
//@return: error

func createTable(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

//@author:
//@function: createPostgresDB
//@description: 创建数据库(postgres)
//@param: dsn string, driver string, createSql
//@return: error

func createPostgresDB(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, _ = db.Exec(createSql)
	return err
}

func initDB(InitDBFunctions ...model.InitDBFunc) (err error) {
	for _, v := range InitDBFunctions {
		err = v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: InitDB
//@description: 创建数据库并初始化
//@param: conf request.InitDB
//@return: error

func InitDB(conf request.InitDB) error {

	switch conf.SqlType {
	case "mysql":
		if conf.Host == "" {
			conf.Host = "127.0.0.1"
		}

		if conf.Port == "" {
			conf.Port = "3306"
		}

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.UserName, conf.Password, conf.Host, conf.Port)
		createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)
		if err := createTable(dsn, "mysql", createSql); err != nil {
			return err
		}

		MysqlConfig := config.Mysql{
			Path:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
			Dbname:   conf.DBName,
			Username: conf.UserName,
			Password: conf.Password,
			Config:   "charset=utf8mb4&parseTime=True&loc=Local",
		}

		if MysqlConfig.Dbname == "" {
			return nil
		}

		linkDns := MysqlConfig.Username + ":" + MysqlConfig.Password + "@tcp(" + MysqlConfig.Path + ")/" + MysqlConfig.Dbname + "?" + MysqlConfig.Config
		mysqlConfig := mysql.Config{
			DSN:                       linkDns, // DSN data source name
			DefaultStringSize:         191,     // string 类型字段的默认长度
			DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false,   // 根据版本自动配置
		}
		if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
			return nil
		} else {
			sqlDB, _ := db.DB()
			sqlDB.SetMaxIdleConns(MysqlConfig.MaxIdleConns)
			sqlDB.SetMaxOpenConns(MysqlConfig.MaxOpenConns)
			global.GVA_DB = db
		}

		err := global.GVA_DB.AutoMigrate(
			model.SysUser{},
			model.SysAuthority{},
			model.SysApi{},
			model.SysBaseMenu{},
			model.SysBaseMenuParameter{},
			model.JwtBlacklist{},
			model.SysDictionary{},
			model.SysDictionaryDetail{},
			model.ExaFileUploadAndDownload{},
			model.ExaFile{},
			model.ExaFileChunk{},
			model.ExaSimpleUploader{},
			model.ExaCustomer{},
			model.SysOperationRecord{},
		)
		if err != nil {
			global.GVA_DB = nil
			return err
		}
		err = initDB(
			source.Admin,
			source.Api,
			source.AuthorityMenu,
			source.Authority,
			source.AuthoritiesMenus,
			source.Casbin,
			source.DataAuthorities,
			source.Dictionary,
			source.DictionaryDetail,
			source.File,
			source.BaseMenu)
		if err != nil {
			global.GVA_DB = nil
			return err
		}
		if err = writeConfig(global.GVA_VP, MysqlConfig); err != nil {
			return err
		}
		global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	case "postgres":
        if conf.Host == "" {
			conf.Host = "127.0.0.1"
		}

		port,err := strconv.Atoi(conf.Port)
		if err != nil {
			return err
		}

		// Get Mysql config, set gorm MaxIdleConns and MaxOpenConns
		mysql := global.GVA_CONFIG.Mysql

		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
			conf.Host, port, conf.UserName, conf.Password)
		createSql := fmt.Sprintf("CREATE DATABASE %s ENCODING 'UTF8';", conf.DBName)
		if err := createPostgresDB(dsn, "postgres", createSql); err != nil {
			return err
		}

		linkDns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			conf.Host, port, conf.UserName, conf.Password, conf.DBName)
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN: linkDns,
			PreferSimpleProtocol: true, // 禁用隐式 prepared statement
		}), &gorm.Config{})
		if err != nil {
			return nil
		} else {
			sqlDB, _ := db.DB()
			sqlDB.SetMaxIdleConns(mysql.MaxIdleConns)
			sqlDB.SetMaxOpenConns(mysql.MaxOpenConns)
			global.GVA_DB = db
		}

		err = global.GVA_DB.AutoMigrate(
			model.SysUser{},
			model.SysAuthority{},
			model.SysApi{},
			model.SysBaseMenu{},
			model.SysBaseMenuParameter{},
			model.JwtBlacklist{},
			model.SysDictionary{},
			model.SysDictionaryDetail{},
			model.ExaFileUploadAndDownload{},
			model.ExaFile{},
			model.ExaFileChunk{},
			model.ExaSimpleUploader{},
			model.ExaCustomer{},
			model.SysOperationRecord{},
		)
		if err != nil {
			global.GVA_DB = nil
			return err
		}

		err = initDB(
			source.Authority,
			source.Admin,
			source.Api,
			source.AuthorityMenu,
			source.BaseMenu,
			source.AuthoritiesMenus,
			source.Casbin,
			source.DataAuthorities,
			source.Dictionary,
			source.DictionaryDetail,
			source.File)
		if err != nil {
			global.GVA_DB = nil
			return err
		}

		PostgresConfig := config.Postgres{
			Path:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
			Dbname:   conf.DBName,
			Username: conf.UserName,
			Password: conf.Password,
			Config:   "",
			MaxIdleConns: mysql.MaxIdleConns,
			MaxOpenConns: mysql.MaxOpenConns,
		}
		if err = writePostgresConfig(global.GVA_VP, PostgresConfig); err != nil {
			return err
		}

		// Change system.db-type: postgres
		var confServer config.Server
		err, confServer = GetSystemConfig()
		if err != nil {
			return err
		}
		configSystem := config.System{
			Env: confServer.System.Env,
			Addr: confServer.System.Addr,
			DbType: "postgres",
			OssType: confServer.System.OssType,
			UseMultipoint: confServer.System.UseMultipoint,
		}
		err = changeDbType(global.GVA_VP, configSystem)
		if err != nil {
			return err
		}
		global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	default:
		global.GVA_LOG.Info("No SqlType setup.")
		return global.GVA_DB.Error
	}
	return nil
	
}
