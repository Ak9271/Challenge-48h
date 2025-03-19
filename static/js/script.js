// Formulaire de contact
document.getElementById('contactForm').addEventListener('submit', function(event) {
    event.preventDefault();
    alert('Message envoyé!');
  });
  
  // Inscriptions aux événements
  document.getElementById('registerBtn').addEventListener('click', function() {
    alert('Inscription réussie!');
  });
  
  // Connexion
  document.getElementById('loginForm').addEventListener('submit', function(event) {
    event.preventDefault();
    alert('Bienvenue, vous êtes connecté!');
  });
  
  // Inscription
  document.getElementById('signupForm').addEventListener('submit', function(event) {
    event.preventDefault();
    alert('Inscription réussie!');
  });
  
  // Administrateur - Mise à jour des événements
  document.getElementById('updateForm').addEventListener('submit', function(event) {
    event.preventDefault();
    alert('Activité mise à jour!');
  });
  