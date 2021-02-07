<template>
    <v-menu
        v-model="menu"
        :close-on-content-click="true"
        :nudge-width="200"
        offset-y
        class="text-center">
        <template v-slot:activator="{ on, attrs }">
            <v-btn
                target="_blank"
                v-bind="attrs"
                v-on="on"
                text
            >
                <v-icon>mdi-cog-outline</v-icon>
            </v-btn>
        </template>

        <v-card>
            <v-card-actions class="justify-center">
                <v-btn
                    text
                    @click="logout"
                >
                    Logout
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-menu>
</template>

<script>
export default {
    name: "Menu",
    data: () => ({
        menu: false
    }),
    methods: {
        logout: function() {
            this.menu = false;
            
            this.$store.commit('setToken', null);
            this.$store.commit('setUser', null);

            this.$store.commit('logout');
            sessionStorage.clear();
            this.$router.push('login');
        }
    }
}
</script>