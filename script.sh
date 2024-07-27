#!/bin/bash

# Fonction pour créer la structure d'un service
create_service() {
    local service_name=$1
    mkdir -p MicroMart/services/$service_name/{cmd,internal/{handler,repository,model}}
    touch MicroMart/services/$service_name/cmd/main.go
    touch MicroMart/services/$service_name/Dockerfile
}

# Création de la structure principale du projet
mkdir -p MicroMart/{services,api-gateway,pkg,deploy,proto,scripts}

# Création des services
create_service "product"
create_service "user"
create_service "cart"
create_service "payment"

# Création de la structure de l'API Gateway
mkdir -p MicroMart/api-gateway/{cmd,internal/proxy}
touch MicroMart/api-gateway/cmd/main.go
touch MicroMart/api-gateway/Dockerfile

# Création des packages partagés
mkdir -p MicroMart/pkg/{auth,database,logger,messaging}

# Création des fichiers de déploiement
touch MicroMart/deploy/docker-compose.yml
mkdir -p MicroMart/deploy/kubernetes

# Création des fichiers proto
touch MicroMart/proto/{product,user,cart,payment}.proto

# Création des fichiers go.mod et go.sum à la racine du projet
touch MicroMart/go.mod MicroMart/go.sum

# Création d'un fichier README.md
cat << EOF > MicroMart/README.md
# MicroMart

MicroMart est une plateforme de commerce électronique basée sur une architecture de microservices, développée en Go.

## Structure du projet

- \`services/\`: Contient les microservices individuels (produit, utilisateur, panier, paiement)
- \`api-gateway/\`: Point d'entrée API pour les clients
- \`pkg/\`: Packages partagés entre les services
- \`deploy/\`: Fichiers de configuration pour le déploiement
- \`proto/\`: Définitions de protocole pour gRPC
- \`scripts/\`: Scripts utilitaires

## Démarrage

(Ajoutez ici les instructions pour configurer et exécuter le projet)

## Contribution

(Ajoutez ici les guidelines pour contribuer au projet)

## Licence

(Spécifiez ici la licence du projet)
EOF

echo "Structure du projet MicroMart créée avec succès!"
