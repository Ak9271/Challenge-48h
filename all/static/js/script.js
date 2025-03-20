document.getElementById('contactForm').addEventListener('submit', function(event) {
    event.preventDefault();
    alert('Message envoyé!');
});

document.getElementById('registerBtn').addEventListener('click', function() {
    alert('Inscription réussie!');
});

document.getElementById('formaulaire-connexion').addEventListener('submit', function(event) {
    event.preventDefault();

    const email = document.getElementById('email-connexion').value;
    const mdp = document.getElementById('mdp-connexion').value;

    fetch('/soumettre-login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, mdp })
    })
    .then(response => response.json())
    .then(data => {
        alert(data.message);
        window.location.href = "/dashboard";
    })
    .catch(error => {
        alert("Erreur de connexion : " + error);
    });
});

document.getElementById('signupForm').addEventListener('submit', function(event) {
    event.preventDefault();

    const nom = document.getElementById('nom-inscription').value;
    const email = document.getElementById('email-inscription').value;
    const mdp = document.getElementById('mdp-inscription').value;
    const mdpConfirmation = document.getElementById('mdp-confirmation').value;

    if (mdp !== mdpConfirmation) {
        alert("Les mots de passe ne correspondent pas.");
        return;
    }

    const data = {
        nom: nom,
        email: email,
        mdp: mdp
    };

    fetch('/soumettre-signup', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
    .then(response => response.json())
    .then(data => {
        alert(data.message);
        window.location.href = '/login';
    })
    .catch((error) => {
        alert("Erreur lors de l'inscription : " + error);
    });
});

document.getElementById('updateForm').addEventListener('submit', function(event) {
    event.preventDefault();
    alert('Activité mise à jour!');
});