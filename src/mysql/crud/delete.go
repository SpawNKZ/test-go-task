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

	sql := "DELETE FROM catogories WHERE `id` = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(1)
	if err != nil {

		panic(err)
	}

	tx.Commit()
}
