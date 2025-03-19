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

func EnregistrerUser(baseDonnee *sql.DB, email string, mdp string) bool {
	_, err := baseDonnee.Exec("INSERT INTO users (email, mdp) VALUES (?, ?)", email, mdp)
	if err != nil {
		return false
	}
	return true
}

func ConnexionUser (baseDonnee *sql.DB, email string, mdp string) bool {
	err := baseDonnee.QueryRow("SELECT * FROM users WHERE email = ? AND mdp = ?", email, mdp).Scan(&email, &mdp)
	if err != nil {
		return false
	}
	return true
}

func AjoutEvent (baseDonnee *sql.DB, titre string, date string, description string) bool {
	_, err := baseDonnee.Exec("INSERT INTO evenement (titre, date, description) VALUES (?, ?, ?)", titre, date, description)
	if err != nil {
		return false
	}
	return true
}

func main() {
	db := CreerBaseDonnee()
	defer db.Close()

	http.HandleFunc("/", GestionRoutes(db))
	fmt.Println("Serveur Go sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
