package database

import (
	"database/sql"
	"fmt"
	"log"
	"udhaar/constants"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func RegisterDatabase() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306", // mysql Port : 3306, postgres: 5432
		AllowNativePasswords: true,
		DBName:               "udhaar_db",
	}
	// mysql -h localhost -u root -p root
	var err error
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal()
	}

	pingErr := Db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println(constants.Green + "\n\n\n Database Connected\n\n\n" + constants.White)
}
