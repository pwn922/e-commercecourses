import { ApolloClient , InMemoryCache, ApolloProvider , createHttpLink, } from '@apollo/client';
import { setContext } from '@apollo/client/link/context';
import App from '../App';

// Crear un link para el HTTP
const httpLink = createHttpLink({
  uri: import.meta.env.URL_GQL, // Asegúrate de que la URL sea correcta
});

// Configurar el contexto de la solicitud para agregar el token
const authLink = setContext((_, { headers }) => {
  const token = localStorage.getItem('token'); // Cambia esto según tu implementación

return {
    headers: {
        ...headers,
        authorization: token ? `Bearer ${token}` : '', // Agregar el token a las cabeceras
        },
    };
});

// Crear el cliente de Apollo
const client = new ApolloClient({
    link: authLink.concat(httpLink), // Combina el authLink con httpLink
    cache: new InMemoryCache(),
});

// Renderizar la aplicación
const ApolloProviderComponent: React.FC = () => {
    return (
        <ApolloProvider client={client}>
            <App />
        </ApolloProvider>
    );
};

export default ApolloProviderComponent;