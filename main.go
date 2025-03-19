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

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/soumettre-signup", soumettreSignupHandler)
	http.HandleFunc("/soumettre-login", soumettreLoginHandler)

	log.Println("Serveur démarré sur : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
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

func soumettreLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		email := r.FormValue("email")
		mdp := r.FormValue("mdp")

		var user InfoUser
		row := db.QueryRow("SELECT id, nom, email, mdp FROM users WHERE email = ?", email).Scan(&user.ID, &user.Nom, &user.Email, &user.Mdp)
		err := row.Scan(&user.ID, &user.Nom, &user.Email, &user.Mdp)
		if err != nil {
			http.Error(w, "Utilisateur introuvable", http.StatusUnauthorized)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Mdp), []byte(mdp))
		if err != nil {
			http.Error(w, "Mot de passe incorrect", http.StatusUnauthorized)
			return
		}
		fmt.Fprintf(w, "Bienvenue %s", user.Nom)

	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

