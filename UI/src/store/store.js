import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)

export const store = new Vuex.Store({
    plugins: [createPersistedState({
        storage: window.sessionStorage,
    })],
    state: {
        token: null,
        user: null,
        loggedIn: false
    },
    getters: {
        token: state => state.token,
        username: state => state.user,
        loggedIn: state => state.loggedIn
    },
    mutations: {
        setToken (state, token) {
            state.token = token
        },
        setUser (state, user) {
            state.user = user
        },
        login (state) {
            state.loggedIn = true
        },
        logout (state) {
            state.loggedIn = false
        }
    }
})