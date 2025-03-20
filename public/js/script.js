const estConnecte = localStorage.getItem("estConnecte") === "true";
const estAdmin = localStorage.getItem("estAdmin") === "true";

function gererAffichageSections() {
    const btnPasses = document.getElementById('btn-passes');
    const btnActuels = document.getElementById('btn-actuels');
    const btnFuturs = document.getElementById('btn-futurs');

    const sectionPasses = document.getElementById('section-passes');
    const sectionActuels = document.getElementById('section-actuels');
    const sectionFuturs = document.getElementById('section-futurs');

    const toutesLesSections = [sectionPasses, sectionActuels, sectionFuturs];

    if (btnPasses) {
        btnPasses.addEventListener('click', () => {
            afficherUneSection(sectionPasses, toutesLesSections);
        });
    }

    if (btnActuels) {
        btnActuels.addEventListener('click', () => {
            afficherUneSection(sectionActuels, toutesLesSections);
        });
    }

    if (btnFuturs) {
        btnFuturs.addEventListener('click', () => {
            afficherUneSection(sectionFuturs, toutesLesSections);
        });
    }
}

function afficherUneSection(sectionAAfficher, toutesLesSections) {
    toutesLesSections.forEach(section => {
        if (section === sectionAAfficher) {
            section.classList.remove('hidden');
        } else {
            section.classList.add('hidden');
        }
    });

    const warningMessage = document.getElementById('warning-message');
    if (warningMessage) {
        warningMessage.style.display = 'none';
    }
}

function afficherFormulaireContact() {
    const formulaire = document.getElementById("contact-form");
    const message = document.getElementById("not-connected-message");

    if (formulaire && message) {
        if (estConnecte) {
            formulaire.style.display = "block";
            message.style.display = "none";
        } else {
            formulaire.style.display = "none";
            message.style.display = "block";
        }
    }
}

function gererActionsActivites() {
    const detailButtons = document.querySelectorAll('.details-btn');
    const inscriptionButtons = document.querySelectorAll('.inscription-btn');
    const warningMessage = document.getElementById('warning-message');

    detailButtons.forEach(button => {
        button.addEventListener('click', () => {
            if (estConnecte) {
                alert("Voici les détails de l'événement !");
            } else {
                afficherMessageConnexion(warningMessage);
            }
        });
    });

    inscriptionButtons.forEach(button => {
        button.addEventListener('click', () => {
            if (estConnecte) {
                alert("Inscription réussie !");
            } else {
                afficherMessageConnexion(warningMessage);
            }
        });
    });
}

function afficherMessageConnexion(warningMessage) {
    if (warningMessage) {
        warningMessage.style.display = 'block';

        setTimeout(() => {
            warningMessage.style.display = 'none';
        }, 5000);
    }
}

function gererSoumissionFormulaire() {
    const form = document.getElementById("contact-form");
    if (form) {
        form.addEventListener("submit", function(e) {
            e.preventDefault();
            document.getElementById("response-message").textContent = "Votre message a bien été envoyé !";
            form.reset();
        });
    }
}

function gererConnexion() {
    const loginForm = document.getElementById("login-form");

    if (loginForm) {
        loginForm.addEventListener("submit", function(e) {
            e.preventDefault();

            const username = document.getElementById("username").value;
            const password = document.getElementById("password").value;

            if (username === "admin" && password === "admin123") {
                localStorage.setItem("estConnecte", "true");
                localStorage.setItem("estAdmin", "true");

                alert("Vous êtes connecté en tant qu'administrateur !");
                window.location.href = "admin.html";
            } else if (username === "user" && password === "user123") {
                localStorage.setItem("estConnecte", "true");
                localStorage.setItem("estAdmin", "false");

                alert("Vous êtes connecté !");
                window.location.href = "accueil.html";
            } else {
                alert("Identifiants incorrects !");
            }
        });
    }
}

function gererDeconnexion() {
    const logoutBtn = document.getElementById("logout-btn");

    if (logoutBtn) {
        logoutBtn.addEventListener("click", () => {
            localStorage.removeItem("estConnecte");
            localStorage.removeItem("estAdmin");
            alert("Vous êtes déconnecté !");
            window.location.reload();
        });
    }
}

function verifierAdminSurPageAdmin() {
    const estSurPageAdmin = window.location.pathname.includes("admin.html");

    if (estSurPageAdmin && !estAdmin) {
        alert("Accès interdit. Vous devez être administrateur.");
        window.location.href = "accueil.html";
    }
}

document.addEventListener('DOMContentLoaded', function() {
    gererAffichageSections();
    afficherFormulaireContact();
    gererActionsActivites();
    gererSoumissionFormulaire();
    gererConnexion();
    gererDeconnexion();
    verifierAdminSurPageAdmin();
});
