const FormumlaireConnexion = document.getElementById('formumlaire-connexion');
const FormumlaireInscription = document.getElementById('formumlaire-inscription');
const MsgConnexion = document.getElementById('msg-connexion');
const MsgInscription = document.getElementById('msg-inscription');

loginForm.addEventListener('submit', function(e) {
    e.preventDefault(); //Empêche rafraichissement de la page pendant la soumission  mail et mdp

    const email = document.getElementById('email-connexion').value;
    const mdp = document.getElementById('mdp-connexion').value;
    const infos = new URLSearchParams();
    infos.append('email', email);
    infos.append('mdp', mdp);

    fetch ('/connexion', {
        method: 'POST',
        header: {
            'Content-Type': 'application/x-www-form-urlencoded', //Envoi des données en format URL
        },
    })
    .then(response => response.text())
    .then(data => {
        if (data ==='Connexion réussie') {
        MsgConnexion.textContent = 'Connexion réussie';
        } else {
            MsgConnexion.textContent = 'Identifiants Incorrects';
        }
    })
    .catch(error =>{
        MsgConnexion.textContent = 'Erreur de connexion';
    });
});