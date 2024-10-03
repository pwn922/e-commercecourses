import React from 'react';
import { Routes, Route } from 'react-router-dom'; // No importamos Router aquí
import HomePage from './pages/HomePage';
import LoginPage from './pages/Auth/LoginPage';
import RegisterPage from './pages/Auth/RegisterPage';
import UserPage from './pages/User/userPage';

const AppRoutes: React.FC = () => {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/register" element={<RegisterPage />} />
      <Route path="/user" element={<UserPage />} />
    </Routes>
  );
};

export default AppRoutes;
