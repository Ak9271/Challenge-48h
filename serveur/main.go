package main

import (
	"fmt"
	"net/http"
	"log"
	"database/sql"
	"github.com/mattn/go-sqlite3"
	)

func creerBaseDonnee() *sql.DB {
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

func enregistrerUser(baseDonnee *sql.DB, email string, mdp string) {
	_, err := baseDonnee.Exec("INSERT INTO users (email, mdp) VALUES (?, ?)", email, mdp)
	if err != nil {
		return false
	}
	return true
}

func connexionUser (baseDonnee *sql.DB, email string, mdp string) {
	rows, err := baseDonnee.Query("SELECT * FROM users WHERE email = ? AND mdp = ?", email, mdp)
	if err != nil {
		return false
	}
	return true
}


