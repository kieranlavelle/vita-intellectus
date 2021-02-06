import axios from 'axios';

export const API = axios.create({
    baseURL: `https://node404.com/gateway/vita`,
    headers: {
      "X-Authenticated-UserId": window.localStorage.getItem('username'),
      "Authorization": "Bearer " + window.localStorage.getItem('token'),
      "Access-Control-Allow-Headers": "*"
    }
})

export const AUTH = axios.create({
    baseURL: `https://node404.com/`,
})