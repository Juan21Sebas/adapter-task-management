package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./task.db")
	if err != nil {
		fmt.Println("Error al abrir la base de datos:", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Error de conexión a la base de datos:", err)
		db.Close()
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
	id TEXT PRIMARY KEY,
    title TEXT,
    description TEXT,
    status BOOLEAN,
    FOREIGN KEY (status) REFERENCES status(name)
)`)
	if err != nil {
		fmt.Println("Error al crear la tabla tasks:", err)
		db.Close()
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS status (
    name TEXT PRIMARY KEY,
    description TEXT,
    active BOOLEAN
)`)
	if err != nil {
		fmt.Println("Error al crear la tabla status:", err)
		db.Close()
		return nil, err
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM status").Scan(&count)
	if err != nil {
		fmt.Println("Error al verificar la tabla status:", err)
		db.Close()
		return nil, err
	}

	if count == 0 {
		_, err = db.Exec(`INSERT INTO status (name, description, active) VALUES (?, ?, ?), (?, ?, ?), (?, ?, ?)`,
			"Pendiente", "Estado Pendiente", true,
			"En Progreso", "Estado En Progreso", true,
			"Completada", "Estado Completada", true)
		if err != nil {
			fmt.Println("Error al insertar datos en la tabla status:", err)
			db.Close()
			return nil, err
		}
	}

	return db, nil
}
