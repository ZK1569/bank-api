# 💰 bank-api

**bank-api** est une API bancaire développée en **Go (Golang)** dans le cadre d’un projet de cours.  
L'application permet d'effectuer les opérations bancaires de base via des endpoints RESTful :

- **Créer un compte**
- **Effacer un compte**
- **Déposer de l'argent**
- **Retirer de l'argent**
- **Consulter un compte**

L'API utilise une base de données **MongoDB** pour stocker les informations des comptes.

---

## 📂 Contenu

- `bank-api.postman_collection.json`  
  Ce fichier contient **toutes les routes de l'API** avec des exemples de requêtes.  
  Tu peux l'importer directement dans **Postman** pour tester rapidement les endpoints.

---

## 🚀 Lancer l'application

### En local (sans Docker)

Assure-toi d’avoir **Go** installé sur ta machine (version 1.24+ recommandée).

1. Clone le dépôt :
   ```bash
   git clone https://github.com/ZK1569/bank-api.git
   cd bank-api
   ```
2. Lance l’application : 
    ```bash 
    go run .
    ```
> 👉 Note : Par défaut, l'application attend une base de données MongoDB accessible via la variable d'environnement MONGODB_URI.
Tu peux utiliser une instance locale ou modifier l'URI dans ton environnement.

--- 

## 🐳 Lancer avec Docker Compose
Une configuration Docker Compose est fournie pour lancer à la fois l’application bank-api et une base de données MongoDB.

### Étapes :
1. Assure-toi d’avoir Docker et Docker Compose installés.
2. Clone le projet :
    ```bash 
    git clone https://github.com/ZK1569/bank-api.git
    cd bank-api
    ```
3. Build et lance les services :
    ```bash
    docker-compose up --build 
    ```

### Une fois les containers lancés :
L’API est disponible sur `http://localhost:8080`.
La base MongoDB est accessible sur `mongodb://root:passwordRoot@localhost:27017/`.

--- 

## Endpoints de l'API

Retrouve la liste complète des routes dans le fichier bank-api.postman_collection.json.

Quelques exemples d'actions disponibles :
- POST v1/bank → créer un compte
- GET v1/bank/{id} → consulter un compte
- GET v1/bank/{id}/total → consulter un le sold total
- DELETE v1/bank/{id} → supprimer un compte
- POST v1/bank/{id}/money → déposer de l'argent
- DELETE v1/bank/{id}/money → retirer de l'argent

--- 

## 🛠️ Technologies utilisées
- Golang
- MongoDB
- Docker / Docker Compose

