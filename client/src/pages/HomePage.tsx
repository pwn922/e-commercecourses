import React from 'react';
import './HomePage.css';
import { useCart } from '../components/Cart/CartContext'; 

const courses = [
  { id: 1, title: 'Curso de React', price: '$999', image: 'https://via.placeholder.com/150' },
  { id: 2, title: 'Curso de Node.js', price: '$899', image: 'https://via.placeholder.com/150' },
  { id: 3, title: 'Curso de CSS Avanzado', price: '$699', image: 'https://via.placeholder.com/150' },
  { id: 4, title: 'Curso de CSS Principiante', price: '$799', image: 'https://via.placeholder.com/150' },
  { id: 5, title: 'Curso de CSS Intermedio', price: '$799', image: 'https://via.placeholder.com/150' },
  { id: 6, title: 'Curso de CSS Experto', price: '$999', image: 'https://via.placeholder.com/150' },
];

const categories = [
  'Desarrollo Web',
  'Diseño Gráfico',
  'Marketing Digital',
  'Fotografía',
  'Literatura',
  'Idiomas',
];

const HomePage: React.FC = () => {
  const { addItem } = useCart(); // Usa el hook para acceder a la función de agregar al carrito

  return (
    <div className="home-container">
      <aside className="sidebar">
        <h2>Categorías</h2>
        <ul>
          {categories.map((category, index) => (
            <li key={index}>{category}</li>
          ))}
        </ul>
      </aside>
      <div className="course-list">
        <div className="header">
          <div className="info-box">
            <h1>Bienvenido a MicroserviceApp</h1>
            <p>Esta es la página de inicio. Aquí puedes encontrar información sobre nuestros cursos disponibles.</p>
          </div>
        </div>
        <div className="courses">
          {courses.map(course => (
            <div className="course-card" key={course.id}>
              <img src={course.image} alt={course.title} />
              <h2>{course.title}</h2>
              <p>Precio: {course.price}</p>
              <button onClick={() => addItem({ id: course.id, title: course.title, price: course.price })}>
                Agregar al carrito
              </button>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default HomePage;
