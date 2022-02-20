package main

import (
	"GolandProjects/src/mysql/util"
	"fmt"
)

func main() {
	db := util.InitDB()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	sql := "INSERT INTO catogories (`name`, `price`, `date`, `type`, `comment`) VALUES (?, ?, ?, ?, ?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	result, err := stmt.Exec("Bought tel", 1000, "2022-02-20", "Расход", "Smth")
	if err != nil {

		panic(err)
	}

	tx.Commit()

	fmt.Println(result.LastInsertId())
}
