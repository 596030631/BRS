package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func Connect() {
	fmt.Println("--------------- connect database -------------------")
	c, err := sql.Open("mysql", "root:Sjh596030631@tcp(Inlets.fun:3306)/brs?charset=utf8")
	Fatal(err)
	Conn = c
}

func Close() {
	_ = Conn.Close()
	fmt.Println("--------------- connect close -------------------")
}
