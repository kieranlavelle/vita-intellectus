import axios from 'axios';

export const API = axios.create({
    baseURL: `http://127.0.0.1:8004/`,
    headers: {
      "X-Authenticated-UserId": window.localStorage.getItem('username'),
      "Access-Control-Allow-Headers": "*"
    }
})

export const AUTH = axios.create({
    baseURL: `https://node404.com/`,
})