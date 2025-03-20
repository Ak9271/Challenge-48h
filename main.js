const express = require('express');
const mongoose = require('mongoose');
const path = require('path');
const app = express();

app.use(express.json());

app.use(express.static(path.join(__dirname, 'template')));
app.use(express.static(path.join(__dirname, 'css')));
app.use(express.static(path.join(__dirname, 'js')));

const port = 8080;

mongoose.connect('mongodb://localhost:27017/MyDataBase')
    .then(() => console.log("Connecté à la database"))
    .catch(err => console.log("Erreur de connexion à la database", err));

app.get('/', (req, res) => {
    res.sendFile(path.join(__dirname, 'index.html'));
});

app.get('/actualite', (req, res) => {
    res.sendFile(path.join(__dirname, 'actualite.html'));
});

app.get('/contact', (req, res) => {
    res.sendFile(path.join(__dirname, 'contact.html'));
});

app.get('/login', (req, res) => {
    res.sendFile(path.join(__dirname, 'login.html'));
});

app.get('/signup', (req, res) => {
    res.sendFile(path.join(__dirname, 'signup.html'));
});


const clientSchema = new mongoose.Schema({
    nom: String,
    prenom: String,
    email: String,
    mdp: String,
    dateDeNaissance: Number
});

const Client = mongoose.model('Clients', clientSchema, 'Clients');

app.post('/client', async (req, res) => {
    try {
        const client = new Client(req.body);
        await client.save();
        res.status(201).json(client);
    } catch (error) {
        res.status(400).json({ message: error.message });
    }
});

app.get('/client', async (req, res) => {
    try {
        const clients = await Client.find();
        res.json(clients);
    } catch (error) {
        res.status(500).json({ message: error.message });
    }
});

app.post('/connecter', async (req, res) => {
    try {
        const { connecteremail, connectermdp } = req.body;

        const client = await Client.findOne({
            email: connecteremail,
            mdp: connectermdp
        });

        if (client) {
            res.status(200).json({
                nom: client.nom,
                prenom: client.prenom,
                email: client.email
            });
        } else {
            res.status(401).json({ message: "Email ou mot de passe incorrect" });
        }
    } catch (error) {
        res.status(500).json({ message: error.message });
    }
});

app.listen(port, () => {
    console.log(`🚀 Server is running on http://localhost:${port}`);
});
