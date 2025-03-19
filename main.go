package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "modernc.org/sqlite"
	"golang.org/x/crypto/bcrypt"
	"html/template"
)

type InfoUser struct {
	ID int
	Nom string
	Email string
	Mdp string
}

var db *sql.DB

func creerDB() (*sql.DB, error) {
	database, err := sql.Open("sqlite", "./test.db")
	if err != nil {
		return nil, err
	}

	creerDB := 
	`CREATE TABLE IF NOT EXISTS users (
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

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", renderTemplate("index.html"))
	http.HandleFunc("/login", renderTemplate("login.html"))
	http.HandleFunc("/signup", renderTemplate("signup.html"))
	http.HandleFunc("/actualite", renderTemplate("actualite.html"))
	http.HandleFunc("/contact", renderTemplate("contact.html"))
	http.HandleFunc("/soumettre-signup", soumettreSignupHandler)
	http.HandleFunc("/soumettre-login", soumettreLoginHandler)

	log.Println("Serveur démarré sur : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderTemplate(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("template/" + filename)
		if err != nil {
			http.Error(w, "Erreur de chargement de la page", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}
}

func soumettreSignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

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
}

func soumettreLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	email := r.FormValue("email")
	mdp := r.FormValue("mdp")

	var user InfoUser
	err := db.QueryRow("SELECT id, nom, email, mdp FROM users WHERE email = ?", email).Scan(&user.ID, &user.Nom, &user.Email, &user.Mdp)

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
}