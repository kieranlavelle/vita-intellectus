<template>
  <v-app fill-height>
    <v-app-bar
      app
      color="white"
      flat
      v-if="this.$store.getters.loggedIn"
    >
      <v-container class="py-0 fill-height fluid">
        <v-avatar
          class="mr-10"
          color="primary"
          size="40"
        >
          {{this.$store.getters.username.charAt(0)}}
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

            if (this.$store.getters.token == null) {
                this.$store.commit('logout');
                if (this.$route.name != 'login' && this.$route.name != 'register'){
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
                        this.$store.commit('setToken', response.headers.token)
                    }
                })
                .catch(response => {
                    this.$store.commit('logout');
                    sessionStorage.clear();

                    this.$store.commit('setToken', null);
                    this.$store.commit('setUser', null);

                    if (this.$route.name != 'login' && this.$route.name != 'register'){
                        this.$router.push('login');
                    }
                })
        }
    },
    created: function(){
        let username = this.$store.getters.username;
        if (username != null) {
            this.initial = username.charAt(0).toUpperCase();
        }
    },
    beforeMount: function() {
        this.checkToken();
        window.setInterval(() => {
            this.checkToken();
        }, 1000*60*10)
    }
};
</script>

<style scoped>

.body-container {
  padding: 0px;
}

</style>