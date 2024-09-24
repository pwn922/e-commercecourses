import React from 'react';
import './App.css';
import NavBar from './components/NavBar/NavBar';
import AppRoutes from './Routes';  // Importa las rutas desde Routes.tsx
import { BrowserRouter as Router } from 'react-router-dom'; // Router aquí
import { CartProvider } from './components/Cart/CartContext'; // Asegúrate de ajustar la ruta

function App() {
  return (
    <CartProvider>
      <Router>
        <NavBar />
        <main style={{ marginTop: '70px', padding: '20px' }}>
          <AppRoutes />  {/* Usa el componente de rutas */}
        </main>
      </Router>
    </CartProvider>
  );
}

export default App;
