#!/bin/bash

# Fonction pour créer la structure de fichiers du service frontend
create_frontend_service() {
    local base_dir="frontend-service"

    # Créer la structure de répertoires
    mkdir -p "$base_dir/public/css"
    mkdir -p "$base_dir/public/js"

    # Créer les fichiers HTML
    touch "$base_dir/public/index.html"
    touch "$base_dir/public/products.html"
    touch "$base_dir/public/cart.html"
    touch "$base_dir/public/profile.html"

    # Créer les fichiers CSS et JS
    touch "$base_dir/public/css/style.css"
    touch "$base_dir/public/js/main.js"
    touch "$base_dir/public/js/products.js"
    touch "$base_dir/public/js/cart.js"
    touch "$base_dir/public/js/profile.js"

    # Créer le Dockerfile
    touch "$base_dir/Dockerfile"
}

# Exécuter la fonction pour créer la structure
create_frontend_service

echo "Structure de fichiers pour le service frontend créée avec succès!"
