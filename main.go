package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	Connect()
	Listener()
	Close()
}

func Fatal(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
