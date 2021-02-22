import { writable } from 'svelte/store';

function createToken() {
    let localStoreToken = localStorage.getItem('token');
    const { subscribe, set, update } = writable(localStoreToken);

    return {
        subscribe,
        setToken: (v) => {
            set(v);
            window.localStorage.setItem('token', v);
        },
        reset: () => {set(null)}
    };
}

export const token = createToken();