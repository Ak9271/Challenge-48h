document.querySelector("#signupForm").addEventListener("submit", function(e) {
    e.preventDefault();

    const infoUtilisateur = {
        nom: this.nom.value,
        prenom: this.prenom.value,
        email: this.email.value,
        mdp: this.mdp.value,
        dateDeNaissance: this.dateDeNaissance.value
    };

    const ajouterUtilisateur = async () => {
        try {
            const reponse = await fetch("http://localhost:8080/user", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(infoUtilisateur)
            });

            if (reponse.status === 201) {   //201 = Creée
                alert("Utilisateur ajouté !");
                this.reset();
            } else {
                alert("Erreur lors de l'ajout de l'utilisateur");
            }
        } catch (error) {
            console.error("Erreur:", error);
            alert("Impossible se connecter au serveur.");
        }
    };
    ajouterUtilisateur();
});

document.querySelector("#loginForm").addEventListener("submit", function(e) {
    e.preventDefault();
});