import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import vuetify from './plugins/vuetify';

import axios from 'axios'
import VueAxios from 'vue-axios'
import VueRouter from 'vue-router'

import router from "./router";

Vue.config.productionTip = false

new Vue({
  vuetify,
  axios,
  VueAxios,
  router,
  render: h => h(App)
}).$mount('#app')
