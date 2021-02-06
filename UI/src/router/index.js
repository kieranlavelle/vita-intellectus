import Vue from "vue";
import VueRouter from "vue-router";

import LoginComponent from "../components/LoginComponent.vue"
import SignUpComponent from "../components/SignUpComponent.vue"
import Habbits from "../components/Habbits.vue"

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "",
    component: Habbits
  },
  {
    path: "/home",
    name: "home",
    component: Habbits
  },
  {
    path: "/login",
    name: "login",
    component: LoginComponent
  },
  {
    path: "/register",
    name: "register",
    component: SignUpComponent
  },
  {
    path: "/habbits",
    name: "habbits",
    component: Habbits
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;