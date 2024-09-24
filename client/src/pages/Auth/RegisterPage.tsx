import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom'; // Asumiendo que usas React Router
import './RegisterPage.css';

const RegisterPage: React.FC = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    // Validación de campos vacíos
    if (!username || !email || !password) {
      setErrorMessage('Por favor, completa todos los campos.');
      return;
    }

    // Lógica de registro (aquí deberías implementar la verificación con backend)
    // Ejemplo de registro exitoso
    console.log('Registro exitoso:', { username, email, password });
    navigate('/home'); // Redirige a la página de inicio después de registrarse
  };

  return (
    <div className="register-container">
      <h1>Registrar</h1>
      {errorMessage && <p className="error-message">{errorMessage}</p>}
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="username">Nombre de Usuario:</label>
          <input 
            type="text" 
            id="username" 
            value={username} 
            onChange={(e) => setUsername(e.target.value)} 
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
          <input 
            type="password" 
            id="password" 
            value={password} 
            onChange={(e) => setPassword(e.target.value)} 
            required 
          />
        </div>
        <button type="submit" className="btn">Registrar</button>
      </form>
      <p>¿Ya tienes una cuenta? <button onClick={() => navigate('/login')} className="link-button">Inicia sesión aquí</button></p>
    </div>
  );
};

export default RegisterPage;
