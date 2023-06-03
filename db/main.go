/*db Provide Db Connection*/
package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "Arshak@32"
		dbname   = "postgres"
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

