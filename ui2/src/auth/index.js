import useStickyState from '../state/store'

function GetAuthHeaders() {
    const [token, setToken] = useStickyState('token', '');

    const config = {
        headers: {
            Authorization: `${token}`
        }
    }

    return config
}

export {GetAuthHeaders}