import axios from 'axios';

const AUTH = axios.create({
    baseURL: 'https://auth.node404.com/',
    timeout: 1000,
    headers: {}
});

const API = axios.create({
    baseURL: 'https://gateway.node404.com/vita/',
    timeout: 2000,
    headers: {}
});

// const API = axios.create({
//     baseURL: 'http://127.0.0.1:8004/',
//     headers: {
//         'X-Authenticated-UserId': 'kmplavelle@gmail.com'
//     }
// });

export {API, AUTH}