package main

import (
	"alert-api-go/internal/database"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Alert struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Level   string `json:"level"`
}

type User struct {
	Nome string `json:"nome"`
}

func main() {
	http.HandleFunc("/alert", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		var alert Alert
		err := json.NewDecoder(r.Body).Decode(&alert)
		if err != nil {
			http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(alert)
	})

	http.HandleFunc("/database", func(w http.ResponseWriter, r *http.Request) {
		database.CreateDatabaseIfNecessary()
	})

	http.HandleFunc("/create_user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		var nome User
		err := json.NewDecoder(r.Body).Decode(&nome)
		if err != nil {
			http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
			return
		}
		database.CreateUser(nome.Nome)
	})

	http.HandleFunc("/select_all_users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		users := database.SelectAllUsers()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	})

	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
