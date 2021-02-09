import axios from 'axios';

export const API = axios.create({
    baseURL: `https://node404.com/gateway/vita`
})

export const AUTH = axios.create({
    baseURL: `https://node404.com/auth`
})
