package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "modernc.org/sqlite"
	"os"
	"golang.org/x/crypto/bcrypt"
)

type InfoUser struct {
	ID int
	Nom string
	Email string
	Mdp string
}

var db *sql.DB

func creerDB() {
	database, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		return nil, err
	}

	creerDB := `CREATE TABLE IF NOT EXISTS users (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"nom" TEXT,
		"email" TEXT,
		"mdp" TEXT
	);`

	_, err = database.Exec(creerDB)
	if err != nil {
		return nil, err
	}
	return database, nil
}

func main() {
	db, err := creerDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/signup.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func soumettreSignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		nom := r.FormValue("nom")
		email := r.FormValue("email")
		mdp := r.FormValue("mdp")
		hash, err := bcrypt.GenerateFromPassword([]byte(mdp), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Erreur lors du hash du mot de passe", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO users (nom, email, mdp) VALUES (?, ?, ?)", nom, email, hash)
		if err != nil {
			http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}


