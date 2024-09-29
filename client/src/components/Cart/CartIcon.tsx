import React from 'react';
import { useCart } from './CartContext'; // Ajusta la ruta según tu estructura de carpetas

const CartIcon: React.FC = () => {
  const { cartItems } = useCart();

  return (
    <div style={{ marginRight: '20px', cursor: 'pointer' }}>
      🛒 Carrito ({cartItems.length})
    </div>
  );
};

export default CartIcon;
