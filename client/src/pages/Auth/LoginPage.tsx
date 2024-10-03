import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { loginUser } from '../../api/authService';
import './LoginPage.css';

const LoginPage: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false); // Estado para controlar la visibilidad de la contraseña
  const [errorMessage, setErrorMessage] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // Validación de campos vacíos
    if (!email || !password) {
      setErrorMessage('Por favor, completa todos los campos.');
      return;
    }
    setLoading(true);
    //[Insertar descripcion]
    try {
      const token = await loginUser(email, password)

      if (token) {
        console.log('Token recibido:', token);
        localStorage.setItem('token', token);
        navigate('/home'); 
    } else {
        setErrorMessage('No se recibió el token.');
    }
    } catch (error: any) {
      setErrorMessage(error.message || 'Error desconocido');
    } finally {
      setLoading(false);
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
        <button type="submit" className="btn" disabled={loading}>
          {loading ? 'Iniciando sesión...' : 'Iniciar Sesión'}
        </button>
      </form>
      <p>¿Olvidaste tu contraseña? <button onClick={handleForgotPassword} className="link-button">Recupérala aquí</button></p>
      <p>¿No tienes una cuenta? <button onClick={handleRegister} className="link-button">Regístrate aquí</button></p>
    </div>
  );
};

export default LoginPage;
