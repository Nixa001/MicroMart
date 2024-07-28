#!/bin/bash

# Nom du projet
PROJECT_NAME="frontend"

# Création du projet React
npx create-react-app $PROJECT_NAME
cd $PROJECT_NAME

# Installation des dépendances supplémentaires
npm install react-router-dom axios

# Suppression des fichiers par défaut non nécessaires
rm src/App.css src/logo.svg src/App.test.js src/setupTests.js

# Création de la structure des dossiers
mkdir -p src/{components,pages,services,styles}

# Création des fichiers
touch src/components/{Header.js,Footer.js,ProductCard.js,CartItem.js}
touch src/pages/{HomePage.js,ProductListPage.js,ProductDetailPage.js,CartPage.js,CheckoutPage.js,UserProfilePage.js,LoginPage.js}
touch src/services/{api.js,auth.js,cart.js}
touch src/styles/{global.css,variables.css}

# Contenu de base pour App.js
cat > src/App.js << EOL
import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Header from './components/Header';
import Footer from './components/Footer';
import HomePage from './pages/HomePage';
import ProductListPage from './pages/ProductListPage';
import ProductDetailPage from './pages/ProductDetailPage';
import CartPage from './pages/CartPage';
import CheckoutPage from './pages/CheckoutPage';
import UserProfilePage from './pages/UserProfilePage';
import LoginPage from './pages/LoginPage';

function App() {
  return (
    <Router>
      <div className="App">
        <Header />
        <Switch>
          <Route exact path="/" component={HomePage} />
          <Route path="/products" component={ProductListPage} />
          <Route path="/product/:id" component={ProductDetailPage} />
          <Route path="/cart" component={CartPage} />
          <Route path="/checkout" component={CheckoutPage} />
          <Route path="/profile" component={UserProfilePage} />
          <Route path="/login" component={LoginPage} />
        </Switch>
        <Footer />
      </div>
    </Router>
  );
}

export default App;
EOL

# Contenu de base pour HomePage.js
cat > src/pages/HomePage.js << EOL
import React from 'react';

function HomePage() {
  return (
    <div className="home-page">
      <h1>Welcome to Our E-commerce Store</h1>
      {/* Add featured products, categories, etc. */}
    </div>
  );
}

export default HomePage;
EOL

# Contenu de base pour ProductListPage.js
cat > src/pages/ProductListPage.js << EOL
import React, { useState, useEffect } from 'react';
import { getProducts } from '../services/api';
import ProductCard from '../components/ProductCard';

function ProductListPage() {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    async function fetchProducts() {
      const data = await getProducts();
      setProducts(data);
    }
    fetchProducts();
  }, []);

  return (
    <div className="product-list-page">
      <h2>Our Products</h2>
      <div className="product-grid">
        {products.map(product => (
          <ProductCard key={product.id} product={product} />
        ))}
      </div>
    </div>
  );
}

export default ProductListPage;
EOL

# Contenu de base pour ProductCard.js
cat > src/components/ProductCard.js << EOL
import React from 'react';
import { Link } from 'react-router-dom';

function ProductCard({ product }) {
  return (
    <div className="product-card">
      <img src={product.image} alt={product.name} />
      <h3>{product.name}</h3>
      <p>\${product.price}</p>
      <Link to={\`/product/\${product.id}\`}>View Details</Link>
    </div>
  );
}

export default ProductCard;
EOL

# Contenu de base pour api.js
cat > src/services/api.js << EOL
const API_BASE_URL = 'http://localhost:8080';

export const getProducts = async () => {
  const response = await fetch(\`\${API_BASE_URL}/products\`);
  return response.json();
};

// Ajoutez d'autres fonctions pour les différentes opérations (getProduct, createOrder, etc.)
EOL

echo "Projet React '$PROJECT_NAME' créé avec succès!"
echo "Pour démarrer le projet, exécutez:"
echo "cd $PROJECT_NAME && npm start"