import Vue from "vue";
import VueRouter from "vue-router";

import LoginComponent from "../components/LoginComponent.vue"
import SignUpComponent from "../components/SignUpComponent.vue"
import Dashboard from "../components/Dashboard.vue"


Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "",
    component: Dashboard
  },
  {
    path: "/home",
    name: "home",
    component: Dashboard
  },
  {
    path: "/login",
    name: "Login",
    component: LoginComponent
  },
  {
    path: "/register",
    name: "Register",
    component: SignUpComponent
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;