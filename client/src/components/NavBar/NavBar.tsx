import React from 'react';
import { Link } from 'react-router-dom';
import './NavBar.css'; 
import SearchBar from '../SearchBar/SearchBar';
import CartIcon from '../Cart/CartIcon';

const NavBar: React.FC = () => {
  return (
    <nav>
      <div className="left-container">
        <button className="menu-button" onClick={() => alert('Abrir menú')}>
          ☰
        </button>
        <Link to="/" className="logo">
          <h2>MicroserviceApp</h2>
        </Link>
        <SearchBar />
      </div>
      
      <div className="nav-items">
        <CartIcon />
        <Link to="/login" className="account-link">
          Mi cuenta
          <div className="user-icon"></div>
        </Link>
      </div>
    </nav>
  );
};

export default NavBar;
