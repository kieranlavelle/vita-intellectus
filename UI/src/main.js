import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import vuetify from './plugins/vuetify';

import VueRouter from 'vue-router'
import router from "./router";
import { store } from './store/store'

Vue.use(VueRouter);
Vue.config.productionTip = false

new Vue({
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
