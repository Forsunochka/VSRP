package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// 1. Подключаемся к PostgreSQL в Docker
	connStr := "postgres://myuser:mysecretpassword@localhost:5432/mydb?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}
	defer db.Close()

	// 2. Проверяем соединение
	err = db.Ping()
	if err != nil {
		log.Fatal("Не могу подключиться к БД:", err)
	}
	fmt.Println("✅ Успешно подключились к PostgreSQL в Docker!")

	// 3. Создаем простую таблицу
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT
		)
	`)
	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}
	fmt.Println("✅ Таблица users создана")

	// 4. Добавляем одного пользователя
	_, err = db.Exec(
		"INSERT INTO users (name, email) VALUES ($1, $2)",
		"Тестовый пользователь", "test@example.com")
	if err != nil {
		fmt.Println("⚠️ Пользователь уже существует или ошибка:", err)
	} else {
		fmt.Println("✅ Пользователь добавлен")
	}

	// 5. Читаем и выводим всех пользователей
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal("Ошибка чтения:", err)
	}
	defer rows.Close()

	fmt.Println("\n📋 Список пользователей:")
	for rows.Next() {
		var id int
		var name, email string
		rows.Scan(&id, &name, &email)
		fmt.Printf("ID: %d, Имя: %s, Email: %s\n", id, name, email)
	}

	fmt.Println("\n🎉 Всё работает!")
}
