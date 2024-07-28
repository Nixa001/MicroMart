const API_BASE_URL = "http://localhost:8080";

export const getProducts = async () => {
  const response = await fetch(`${API_BASE_URL}/products`);
  console.log(response.json());
  return response.json();
};

// Ajoutez d'autres fonctions pour les différentes opérations (getProduct, createOrder, etc.)
