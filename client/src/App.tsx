
import './App.css';
import NavBar from './components/NavBar/NavBar';
import AppRoutes from './Routes';  
import { BrowserRouter as Router } from 'react-router-dom'; 
import { CartProvider } from './components/Cart/CartContext';

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
