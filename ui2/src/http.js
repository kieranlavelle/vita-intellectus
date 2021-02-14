import axios from 'axios';

export const API = axios.create({
    baseURL: `https://gateway.node404.com/vita`
})

export const AUTH = axios.create({
    baseURL: `https://auth.node404.com/`
})
