import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { registerUser } from '../../api/authService';
import './RegisterPage.css';

const RegisterPage: React.FC = () => {
  // Estados para los campos requeridos
  const [firstName, setFirstName] = useState('');
  const [middleName, setMiddleName] = useState('');
  const [lastName, setLastName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false); // Estado para controlar la visibilidad de la contraseña
  const [showConfirmPassword, setShowConfirmPassword] = useState(false); // Controlar visibilidad de confirmación de contraseña
  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // Validación de campos vacíos
    if (!firstName || !middleName || !lastName || !email || !password || !confirmPassword) {
      setErrorMessage('Por favor, completa todos los campos.');
      return;
    }

    // Verificación de que las contraseñas coincidan
    if (password !== confirmPassword) {
      setErrorMessage('Las contraseñas no coinciden.');
      return;
    }
    try {
      const token = await registerUser(firstName,middleName,lastName,email, password)

      if (token) {
        console.log('Token recibido:', token);
        //localStorage.setItem('token', token);
        navigate('/home'); 
    } else {
        setErrorMessage('No se recibió el token.');
    }
    } catch (error: any) {
      setErrorMessage(error.message || 'Error desconocido');
    }
  };

  const togglePasswordVisibility = () => {
    setShowPassword(!showPassword);
  };

  const toggleConfirmPasswordVisibility = () => {
    setShowConfirmPassword(!showConfirmPassword);
  };

  return (
    <div className="register-container">
      <h1>Registrar</h1>
      {errorMessage && <p className="error-message">{errorMessage}</p>}
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="firstName">Nombre:</label>
          <input 
            type="text" 
            id="firstName" 
            value={firstName} 
            onChange={(e) => setFirstName(e.target.value)} 
            required 
          />
        </div>
        <div className="form-group">
          <label htmlFor="middleName">Segundo Nombre:</label>
          <input 
            type="text" 
            id="middleName" 
            value={middleName} 
            onChange={(e) => setMiddleName(e.target.value)} 
          />
        </div>
        <div className="form-group">
          <label htmlFor="lastName">Apellido:</label>
          <input 
            type="text" 
            id="lastName" 
            value={lastName} 
            onChange={(e) => setLastName(e.target.value)} 
            required 
          />
        </div>
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
              type={showPassword ? 'text' : 'password'} 
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
        <div className="form-group">
          <label htmlFor="confirmPassword">Confirmar Contraseña:</label>
          <div className="password-wrapper">
            <input 
              type={showConfirmPassword ? 'text' : 'password'} 
              id="confirmPassword" 
              value={confirmPassword} 
              onChange={(e) => setConfirmPassword(e.target.value)} 
              required 
            />
            <button
              type="button"
              className="toggle-password-btn"
              onClick={toggleConfirmPasswordVisibility}
            >
              {showConfirmPassword ? 'Ocultar' : 'Mostrar'}
            </button>
          </div>
        </div>
        <button type="submit" className="btn">Registrar</button>
      </form>
      <p>¿Ya tienes una cuenta? <button onClick={() => navigate('/login')} className="link-button">Inicia sesión aquí</button></p>
    </div>
  );
};

export default RegisterPage;
