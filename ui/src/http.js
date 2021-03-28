import axios from 'axios';

const AUTH = axios.create({
    baseURL: 'https://auth.node404.com/',
    timeout: 1000,
    headers: {}
});

const API = axios.create({
    baseURL: 'https://gateway.node404.com/vita/',
    timeout: 1000,
    headers: {}
});

export {API, AUTH}