package main

import (
	"GolandProjects/src/mysql/util"
)

func main() {

	db := util.InitDB()

	tx, err := db.Begin()
	if err != nil {

		panic(err)
	}

	sql := "UPDATE categories SET `price` = ? WHERE `id` = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec("2000", 2)
	if err != nil {

		panic(err)
	}

	tx.Commit()
}
