package util

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "mysql:dias12345@tcp(127.0.0.1:3306)/testtask")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(10)
	// Устанавливаем максимальное количество неактивных соединений в базе данных
	db.SetMaxIdleConns(5)
	// Проверяем соединение
	if err := db.Ping(); err != nil {
		panic(err)
	}
	// Возвращаем ссылку на указатель соединения с базой данных
	return db

}
