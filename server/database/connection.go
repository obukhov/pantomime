package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func NewConnection() Connection {
	return Connection{
		dsn: "root:root@tcp(" + os.Getenv("DB_HOST") + ")/pantomime",
	}
}

type Connection struct {
	dsn string
}

func (c *Connection) Open() *gorm.DB {
	db, err := gorm.Open("mysql", c.dsn)
	if err != nil {
		panic("failed to connect database")
	}

	//db.LogMode("1" == os.Getenv("DB_DEBUG"))

	return db
}
