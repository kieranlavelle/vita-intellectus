import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export const store = new Vuex.Store({
    state: {
        token: "",
        user: ''
    },
    getters: {
        token: state => state.token
    },
    mutations: {
        setToken (state, token) {
            state.token = token
        },
        setUser (state, user) {
            state.user = user
        }
    }
})