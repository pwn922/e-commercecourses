import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './LoginPage.css';

const LoginPage: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false); // Estado para controlar la visibilidad de la contraseña
  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    // Validación de campos vacíos
    if (!email || !password) {
      setErrorMessage('Por favor, completa todos los campos.');
      return;
    }

    // Lógica de autenticación (aquí deberías implementar la verificación con backend)
    if (email === 'test@example.com' && password === 'password123') {
      // Simulando autenticación exitosa
      console.log('Inicio de sesión exitoso');
      navigate('/home'); // Redirige a la página de inicio después de iniciar sesión
    } else {
      setErrorMessage('Correo o contraseña incorrectos.');
    }
  };

  const handleRegister = () => {
    navigate('/register'); // Redirige a la página de registro
  };

  const handleForgotPassword = () => {
    navigate('/forgot-password'); // Redirige a la página de recuperación de contraseña
  };

  // Alterna la visibilidad de la contraseña
  const togglePasswordVisibility = () => {
    setShowPassword(!showPassword);
  };

  return (
    <div className="login-container">
      <h1>Iniciar Sesión</h1>
      {errorMessage && <p className="error-message">{errorMessage}</p>}
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="email">Correo Electrónico:</label>
          <input 
            type="email" 
            id="email" 
            value={email} 
            onChange={(e) => setEmail(e.target.value)} 
            required 
          />
        </div>
        <div className="form-group">
          <label htmlFor="password">Contraseña:</label>
          <div className="password-wrapper">
            <input 
              type={showPassword ? 'text' : 'password'} // Cambia el tipo de input para mostrar u ocultar la contraseña
              id="password" 
              value={password} 
              onChange={(e) => setPassword(e.target.value)} 
              required 
            />
            <button
              type="button"
              className="toggle-password-btn"
              onClick={togglePasswordVisibility}
            >
              {showPassword ? 'Ocultar' : 'Mostrar'}
            </button>
          </div>
        </div>
        <button type="submit" className="btn">Iniciar Sesión</button>
      </form>
      <p>¿Olvidaste tu contraseña? <button onClick={handleForgotPassword} className="link-button">Recupérala aquí</button></p>
      <p>¿No tienes una cuenta? <button onClick={handleRegister} className="link-button">Regístrate aquí</button></p>
    </div>
  );
};

export default LoginPage;
