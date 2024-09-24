import React from 'react';

const SearchBar: React.FC = () => {
  return (
    <input 
      type="text" 
      placeholder="Buscar servicios..." 
      style={searchBarStyle} 
    />
  );
};

// Estilo para la barra de búsqueda
const searchBarStyle: React.CSSProperties = {
  width: '400px', // Ajuste del ancho para hacerla más grande
  padding: '10px',
  fontSize: '16px', // Tamaño de letra más grande
  borderRadius: '5px',
  border: '1px solid #ccc',
};

export default SearchBar;
