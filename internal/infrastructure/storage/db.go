package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/sdimitrenco/grammurrr/internal/config"
)

func NewPostgresDB() *sql.DB {
	config := config.GetConfig()

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s", // Убрал `$(...)`
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.DbSsl, // Должно быть строкой ("disable" или "require")
	)

	db, err := sql.Open(config.DB.Type, connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("БД недоступна:", err)
	}

	fmt.Println("Подключение к PostgreSQL успешно!")
	return db
}
