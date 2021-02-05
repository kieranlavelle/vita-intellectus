<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
      v-if="loggedIn()"
    >
      <div class="d-flex align-center">
        <h1>Vita</h1>
      </div>

      <v-spacer></v-spacer>
      <Menu></Menu>

      <!-- <v-btn
        @click="logout"
        target="_blank"
        text
      >
        <span class="mr-2">Logout</span>
        <v-icon>mdi-dots-vertical</v-icon>
      </v-btn> -->
    </v-app-bar>

    <v-main>
      <router-view></router-view>
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
  }),
  methods: {
    loggedIn: function() {
      return window.localStorage.getItem('token') != "null"
    }
  }
};
</script>
