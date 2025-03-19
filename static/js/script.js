document.getElementById('contactForm').addEventListener('submit', function(event) {
    event.preventDefault();
    alert('Message envoyé!');
});

document.getElementById('registerBtn').addEventListener('click', function() {
    alert('Inscription réussie!');
});

document.getElementById('loginForm').addEventListener('submit', function(event) {
    event.preventDefault();
    alert('Bienvenue, vous êtes connecté!');
});

document.getElementById('signupForm').addEventListener('submit', function(event) {
    event.preventDefault();
    alert('Inscription réussie!');
});

document.getElementById('updateForm').addEventListener('submit', function(event) {
    event.preventDefault();
    alert('Activité mise à jour!');
});
