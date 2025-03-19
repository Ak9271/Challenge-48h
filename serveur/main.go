package main

import (
	"fmt"
	"net/http"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	)

func CreerBaseDonnee() *sql.DB {
	baseDonnee, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	creerBaseDonnee := 
	`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT,
		mdp TEXT
	);

	CREATE TABLE IF NOT EXISTS evenement (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		titre TEXT,
		date TEXT,
		desciption TEXT,
	);`

	_, err = baseDonnee.Exec(creerBaseDonnee)
	if err != nil {
		log.Fatal(err)
	}
	return baseDonnee
}

func GestionRoutes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/inscription":
			if r.Method == "POST" {
				nom := r.FormValue("nom")
				email := r.FormValue("email")
				mdp := r.FormValue("mdp")
				_, err := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", name, email, password)
				if err != nil {
					http.Error(w, "Erreur d'inscription", http.StatusInternalServerError)
					return
				}

				w.Write([]byte("Inscription réussie"))
			}
		case "/connexion":
			if r.Method == "POST" {
				email := r.FormValue("email")
				mdp := r.FormValue("mdp")

				a := db.QueryRow("SELECT password FROM users WHERE email = ?", email)
				var mdpStocke string
				err := a.Scan(&storedPassword)
				if err != nil || mdpStocke != mdp {
					w.Write([]byte("Identifiants incorrects"))
					return
				}
				w.Write([]byte("Connexion réussie"))
			}
		default:
			http.NotFound(w, r)
		}
	}
}

func main() {
	db := CreerBaseDonnee()
	defer db.Close()

	http.HandleFunc("/", GestionRoutes(db))
	fmt.Println("Serveur Go sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
