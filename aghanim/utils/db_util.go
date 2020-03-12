package utils

import (
	"log"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type dbUtil struct {
	db *gorm.DB
}

var dbOnce sync.Once
var dbInstance *dbUtil

// GetDBConnection gets DB connection
func GetDBConnection() *gorm.DB {
	dbOnce.Do(func() {
		log.Println("Initialize DB connection...")

		//setup connection to database
		conn := viper.GetString("DB_USER") + ":" + viper.GetString("DB_PASSWORD") + "@tcp(" + viper.GetString("DB_HOST") + ":" + viper.GetString("DB_PORT") + ")/" + viper.GetString("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"
		db, err := gorm.Open(viper.GetString("DB_TYPE"), conn)
		if err != nil {
			log.Fatal("cant connect : ", err)
		}

		/**
		* NOTES: this will set connection lifetime in connection pool to 1 minute.
		* 		  If the connection in the pool is idle > 1 min, Golang will close it
		* 		  and will create new connection if #connections in the pool < pool max num
		* 		  of connection. This to avoid invalid connection issue
		 */
		db.DB().SetConnMaxLifetime(time.Second * 60)

		// Set as singular table
		db.SingularTable(true)

		if err != nil {
			log.Println(err)
		}

		// DB setup
		db.LogMode(false)
		dbInstance = &dbUtil{
			db: db,
		}
	})

	return dbInstance.db
}
