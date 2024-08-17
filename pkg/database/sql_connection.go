package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

func main() {

	db, err := sql.Open("godror", "admin/Babayaga2@@!@g6d2754f1a6e097_suidvr4emqjvdxuh_tp.adb.oraclecloud.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select sysdate from dual")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)
}
