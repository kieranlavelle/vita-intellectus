<template>
  <v-app fill-height>
    <v-app-bar
      app
      color="white"
      flat
      v-if="loggedIn()"
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

import Menu from './components/Menu'

export default {
  name: 'App',

  components: {
    Menu
  },

  data: () => ({
    links: [
      "Habbits"
    ],
    initial: "",
  }),
  methods: {
    loggedIn: function() {
      var loggedIn = window.localStorage.getItem('token') != "null"
      if (!loggedIn) {
        this.$router.push('login');
      }
      return loggedIn;
    },
    navigate: function(destination) {
      this.$router.push(destination.toLowerCase());
    }
  },
  created: function(){
    let username = window.localStorage.getItem("username");
    this.initial = username.charAt(0).toUpperCase();
  }
};
</script>

<style scoped>

.body-container {
  padding: 0px;
}

</style>