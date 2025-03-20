const info = document.querySelector("#formulaire-inscription");
info.addEventListener("submit", function(e) {
    e.preventDefault();

    const infoUtilisateur = {
        nom: document.querySelector("#nom-inscription").value,
        email: document.querySelector("#email-inscription").value,
        mdp: document.querySelector("#mdp-inscription").value
    };

    const ajouterUtilisateur = async () => {
        try {
            const reponse = await fetch("http://localhost:8080/client", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(infoUtilisateur)
            });

            if (reponse.status === 201) {   //201 = Creéé
                alert("Utilisateur ajoute");
                info.reset();
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

const infoConnect = document.querySelector("#formaulaire-connexion")
infoConnect.addEventListener("submit", function(e) {
    e.preventDefault();

    const infoConnexion = {
        email: document.querySelector("#email-connexion").value,
        mdp: document.querySelector("#mdp-connexion").value
    };

    const connecterUtilisateur = async () => {
        try {
            const reponse = await fetch("http://localhost:8080/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(infoConnexion)
            });

            if (reponse.ok) {
                const rep = await reponse.json();
                localStorage.setItem("utilisateur", JSON.stringify(rep));
                window.location.href = "index.html";
            } else {
                alert("Email ou mot de passe incorrect");
            }
        } catch (error) {
            console.error("Erreur:", error);
            alert("Impossible se connecter au serveur.");
        }
    };
    connecterUtilisateur();
});