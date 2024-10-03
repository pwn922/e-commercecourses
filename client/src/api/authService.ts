import axios from 'axios';

const url = import.meta.env.VITE_AUTHMS_URL;

export const loginUser = async (email: String, password: String) => {    //Hay que ver si cambiar el nombre del handler
    try{
        const response = await axios.post(`${url}/login`, { 
            email,
            password,
        },{
            headers: {
                'Content-Type': 'application/json',
            },
        });

        return response.data.accessToken; 
    }catch (error) {
        console.error('Error al iniciar sesion:', error);
        throw error;
    }
}

export const registerUser = async (first_name: String, middle_name: String, last_name: String,email: string, password: string) => {
    try {
        const role = "student"
        //Recordar, las peticiones del axios son del tipo "var": "data", que tiene que ser igual al del backend en este caso first_name, etc.
        const response = await axios.post(`${url}/register`, {
            first_name,
            middle_name,
            last_name,
            email,
            password,
            role,
        },{
            headers: {
                'Content-Type': 'application/json',
            },
    });

    return response.data.accessToken;

    } catch (error) {
        console.error('Error en register:', error);
        throw error;
    }
};