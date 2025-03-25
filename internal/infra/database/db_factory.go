package database

import (
	"database/sql"
	"log"
	"sync"

	_ "modernc.org/sqlite"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("sqlite", "./autenticapb.db")
		if err != nil {
			log.Fatal("Erro ao abrir o banco de dados:", err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatal("Erro ao conectar no banco de dados:", err)
		}

		sqlStmt := `CREATE TABLE IF NOT EXISTS user (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(200) NOT NULL,
            cpf VARCHAR(11) NOT NULL,
            nome VARCHAR(200) NOT NULL,
            password VARCHAR(255) NOT NULL,
            email VARCHAR(200) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			log.Fatalf("Erro ao criar tabela: %q: %s\n", err, sqlStmt)
		}
	})
	return db
}
