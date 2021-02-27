import axios from 'axios';

const AUTH = axios.create({
    baseURL: 'https://auth.node404.com'
})

const API = axios.create({
    baseURL: 'https://gateway.node404.com/vita'
})

function getHeaders() {
    return {
        headers: {
            Authorization: window.localStorage.getItem('token')
        }
    }
}

export {
    AUTH,
    API,
    getHeaders
}