package database

import (
	"database/sql"
	"fmt"
	 _ "github.com/lib/pq"
)

// const (
//
//	host     = "localhost"
//	port     = 5432
//	user     = "postgres"
//	password = "postgres"
//	dbname   = "newlibrary"
//	sslmode  = "disable"
//
// )
// var DB *sql.DB

func DbConnection()(*sql.DB,error){
	psqlconn := fmt.Sprintf("host= %s port = %d user = %s password= %s dbname= %s sslmode= %s", "localhost", 5432, "postgres", "postgres", "newlibrary", "disable")

	// connection open
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		// return nil, fmt.Errorf("failed to open database: %v", err)
		fmt.Println("connect error ", err)
	}
	fmt.Println("db connected")
	return db, nil
}
