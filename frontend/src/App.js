import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
// import Header from "./components/Header.js";
// import Footer from "./components/Footer.js";
import HomePage from "./pages/HomePage.js";
import ProductListPage from "./pages/ProductListPage.js";
// import ProductDetailPage from "./pages/ProductDetailPage.js";
// import CartPage from "./pages/CartPage.js";
// import CheckoutPage from "./pages/CheckoutPage.js";
// import UserProfilePage from "./pages/UserProfilePage.js";
// import LoginPage from "./pages/LoginPage.js";

function App() {
  return (
    <Router>
      <div className="App">
        {/* <Header /> */}
        <Routes>
          <Route exact path="/" component={HomePage} />
          <Route path="/products" component={ProductListPage} />
          {/* 
          <Route path="/product/:id" component={ProductDetailPage} />
          <Route path="/cart" component={CartPage} />
          <Route path="/checkout" component={CheckoutPage} />
          <Route path="/profile" component={UserProfilePage} />
          <Route path="/login" component={LoginPage} /> */}
        </Routes>
        {/* <Footer /> */}
      </div>
    </Router>
  );
}

export default App;
