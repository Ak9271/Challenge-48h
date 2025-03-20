const mongoose = require('mongoose');
const express = require('express');
const app = express();

app.use(express.json())
app.use(express.static('all'));

mongoose.connect('mongo://localhost:27017/BaseDonnee')
.then(() => {console.log("Connection à la base de données")})
.catch((err) => {console.log("Erreur de connection", err)});

const userInfo = new mongoose.Schema({
    nom: String,
    email: String,
    mdp: String,
});

const User = mongoose.model('Users', userInfo, 'Users');

app.post('/signup', async (req, res) => {
    try {
        const { nom, email, mdp } = req.body;
        const user = new User({ nom, email, mdp });
        await user.save();    
        res.status(201).json({
            message: "Inscription réussie",
            user: { nom, email }
        });
    } catch (error) {
        res.status(400).json({ message: error.message });
    }
});

app.post('/login', async (req, res) => {
    try {

    }
});