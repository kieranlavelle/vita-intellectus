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
                <v-icon>mdi-dots-vertical</v-icon>
            </v-btn>
        </template>

        <v-card>
            <v-list>
                <v-list-item>
                    <v-list-item-content>
                        <v-list-item-title>Kieran Lavelle</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>

                <v-divider></v-divider>

                <v-list-item>
                    <v-list-item-content>
                        <v-list-item-action class="text-center">
                            <router-link to="/habbits">Habbits</router-link>
                        </v-list-item-action>
                    </v-list-item-content>
                </v-list-item>


            </v-list>

            <v-divider></v-divider>

            <v-card-actions>
                <v-spacer></v-spacer>
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
            window.localStorage.setItem('token', "null")
            this.$router.push('login');
        }
    }
}
</script>