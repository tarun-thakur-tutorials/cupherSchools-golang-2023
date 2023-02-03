package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func RegisterDatabase(userName, pass, dbName string) {
	cfg := mysql.Config{
		User:                 userName,
		Passwd:               pass,
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306", // mysql Port : 3306, postgres: 5432
		AllowNativePasswords: true,
		DBName:               dbName,
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

	fmt.Println("\n\n\n Database Connected\n\n\n")
}
