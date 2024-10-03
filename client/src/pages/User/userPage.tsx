import React from 'react';
import './UserPage.css';
import { useGetUser } from '../../api/userService';

const UserPage: React.FC = () => {
    const { loading, error, data } = useGetUser(); 

    const handleShowUser = () => {
        console.log(data)
        if (data && data.user) {
        console.log('Usuario:', data.user.id); 
    } else {
        console.error('No se pudo obtener el usuario');
    }};

    //Por ahora el unico problema que existe en con la peticiones hacia el backend, debido a que las manda siempre
    //desde la url de vite.

    return (
    <div className="page-container">
    <button onClick={handleShowUser} className="user-btn">
        {loading ? 'Cargando...' : 'Mostrar Usuario'}
    </button>
      {error && <p>Error al cargar el usuario: {error.message}</p>} {/* Mensaje de error */}
    </div>
    );
};

export default UserPage;
