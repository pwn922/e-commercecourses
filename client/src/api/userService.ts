import { gql, useQuery } from '@apollo/client';


const GET_USERS = gql`query {
    user {
        id
        first_name
        last_name
        email
        role {
            id
            roleName
        }
    }
}
`;

// Crear el hook para obtener usuarios
export const useGetUser = () => {
    const { loading, error, data } = useQuery(GET_USERS);
    return { loading, error, data };
};