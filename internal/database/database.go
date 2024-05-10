package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func CreateUser(nome string) {
	db, err := sql.Open("sqlite3", "C:/Users/richa/OneDrive/Área de Trabalho/go/usuarios.db")
	if err != nil {
		fmt.Println("Erro ao abrir o banco de dados:", err)
	}
	_, err = db.Exec("INSERT INTO usuarios (nome) VALUES (?)", nome)
	if err != nil {
		fmt.Println("Erro ao criar o usário:", err)
		return
	}
	fmt.Println("Usário criado com sucesso.")
}

func CreateDatabaseIfNecessary() {
	db, err := sql.Open("sqlite3", "C:/Users/richa/OneDrive/Área de Trabalho/go/usuarios.db")
	if err != nil {
		fmt.Println("Erro ao abrir o banco de dados:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS usuarios (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				nome TEXT
			)
		`)
	if err != nil {
		fmt.Println("Erro ao criar a tabela:", err)
		return
	}

	fmt.Println("Banco de dados e tabela prontos para uso.")
}

func SelectAllUsers() string {
	db, err := sql.Open("sqlite3", "C:/Users/richa/OneDrive/Área de Trabalho/go/usuarios.db")
	if err != nil {
		fmt.Println("Erro ao abrir o banco de dados:", err)
		return ""
	}
	defer db.Close()

	stmt, _ := db.Prepare("SELECT * FROM usuarios")

	rows, _ := stmt.Query()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Nome)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	return string(jsonData)
}
