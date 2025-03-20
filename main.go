package main

import (
	"database/sql"
	"log"
	"net/http"
	_ "modernc.org/sqlite"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"encoding/json"
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
    if r.Method != http.MethodPost {
        http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
        return
    }

    var user struct {
        Nom     string `json:"nom"`
        Email   string `json:"email"`
        Mdp     string `json:"mdp"`
    }

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Erreur lors de la lecture des données", http.StatusBadRequest)
        return
    }

    var existingUser string
    err = db.QueryRow("SELECT email FROM users WHERE email = ?", user.Email).Scan(&existingUser)
    if err == nil {
        http.Error(w, "Cet email est déjà utilisé", http.StatusConflict)
        return
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(user.Mdp), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Erreur lors du hash du mot de passe", http.StatusInternalServerError)
        return
    }

    _, err = db.Exec("INSERT INTO users (nom, email, mdp) VALUES (?, ?, ?)", user.Nom, user.Email, hash)
    if err != nil {
        http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Inscription réussie"})
}

func soumettreLoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
        return
    }

    var user struct {
        Email string `json:"email"`
        Mdp   string `json:"mdp"`
    }

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Erreur lors de la lecture des données", http.StatusBadRequest)
        return
    }

    var dbUser InfoUser
    err = db.QueryRow("SELECT id, nom, email, mdp FROM users WHERE email = ?", user.Email).Scan(&dbUser.ID, &dbUser.Nom, &dbUser.Email, &dbUser.Mdp)

    if err != nil {
        http.Error(w, "Utilisateur introuvable", http.StatusUnauthorized)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(dbUser.Mdp), []byte(user.Mdp))
    if err != nil {
        http.Error(w, "Mot de passe incorrect", http.StatusUnauthorized)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Connexion réussie"})
}