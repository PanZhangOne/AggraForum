package datasource

import (
	"fmt"
	"forum/conf"
	"forum/entitys"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

var (
	masterDB *gorm.DB
	slaveDB  *gorm.DB
	lockGorm sync.Mutex
)

func InstanceGormMaster() *gorm.DB {
	if masterDB != nil {
		return masterDB
	}
	lockGorm.Lock()
	defer lockGorm.Unlock()
	if masterDB != nil {
		return masterDB
	}

	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	DB, err := gorm.Open(conf.DriverName, driveSource)
	if err != nil {
		log.Fatal("DB ERROR.InstanceGormMaster", err)
		return nil
	}

	DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")

	masterDB = DB
	masterDB.AutoMigrate(&entitys.User{}, &entitys.Label{}, &entitys.Topic{}, &entitys.Reply{},
		&entitys.CollectTopic{})
	return DB
}

func InstanceGormSlave() *gorm.DB {
	if slaveDB != nil {
		return slaveDB
	}

	lockGorm.Lock()
	defer lockGorm.Unlock()

	if slaveDB != nil {
		return slaveDB
	}

	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	DB, err := gorm.Open(conf.DriverName, driveSource)
	if err != nil {
		log.Fatal("DB ERROR.InstanceGormMaster", err)
		return nil
	}

	DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")

	slaveDB = DB
	return DB
}
