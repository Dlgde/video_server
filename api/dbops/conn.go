package dbops

import (
	"database/sql"
	"fmt"
	//"github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	fmt.Println("Entering conn.go init function...")
	dbConn, err = sql.Open("mysql", "root:909923@/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("dbConn +%v\n", dbConn)
}
