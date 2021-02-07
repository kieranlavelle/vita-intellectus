<template>
  <v-app fill-height>
    <v-app-bar
      app
      color="white"
      flat
      v-if="loggedIn"
    >
      <v-container class="py-0 fill-height fluid">
        <v-avatar
          class="mr-10"
          color="primary"
          size="40"
        >
          {{initial}}
        </v-avatar>

        <v-btn
          v-for="link in links"
          :key="link"
          @click="navigate(link)"
          text
        >
          {{ link }}
        </v-btn>

        <v-spacer></v-spacer>
        <Menu />
      </v-container>
    </v-app-bar>

    <v-main class="grey lighten-3">
      <v-container fluid fill-height class="body-container">
        <v-row style="height: 100%">
          <router-view></router-view>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>

import { AUTH } from "./endpoints/http-config"
import Menu from './components/Menu'

export default {
    name: 'App',
    components: {
        Menu
    },
    data: () => ({
        loggedIn: false,
        links: [
            "Habbits"
        ],
        initial: "",
    }),
    methods: {
        navigate: function(destination) {
            this.$router.push(destination.toLowerCase());
        },
        checkToken: function() {

            if (this.$store.getters.token == "") {
                this.loggedIn = false;
                if (this.$route.name != 'login'){
                    this.$router.push('login');
                }
                return
            }

            const config = {
                headers: {
                    Authorization: `Bearer ${this.$store.getters.token}`
                }
            }

            AUTH.put("/refresh", {}, config)
                .then(response => {
                    if (response.status == 200) {
                        this.loggedIn = true;
                        this.$store.commit('setToken', response.headers.token)
                    }
                })
                .catch(response => {
                    this.loggedIn = false;
                    this.$store.commit('setToken', null);
                    this.$store.commit('setUser', null);

                    if (this.$route.name != 'login'){
                        this.$router.push('login');
                    }
                })
        }
    },
    created: function(){
        let username = window.localStorage.getItem("username");
        this.initial = username.charAt(0).toUpperCase();
    },
    beforeMount: function() {
        this.checkToken();
        window.setInterval(() => {
            this.checkToken();
        }, 1000*10*1)
    }
};
</script>

<style scoped>

.body-container {
  padding: 0px;
}

</style>