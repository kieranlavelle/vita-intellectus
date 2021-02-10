import { AUTH } from '../http'
import usePersistedState from 'use-persisted-state-hook'

function Session(){
    const [token, setToken] = usePersistedState('token', '');

    const config = {
        headers: {
            Authorization: `Bearer ${token}`
        }
    }

    let loggedIn = false;

    AUTH.put('/refresh', {}, config)
        .then(response => {
            // setToken(response.headers.token)
            console.log("logged in")
            loggedIn = true;
        })
        .catch(error => {
            loggedIn = false;
            console.log(error);
        })
    
    return loggedIn
}

function CheckToken(){
    const [token, setToken] = usePersistedState('token', '');
    return token != '';
}

export {Session, CheckToken}