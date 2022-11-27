package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GormDb() *gorm.DB {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/androidmarket?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // use singular table name, table for `User` would be `user` with this option enabled
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	return db
}
