<template>
    <v-container fluid full-height>
        <v-row class="nav-row">
            <v-btn
                color="deep-purple lighten-3"
            >
                New Habbit
            </v-btn>
        </v-row>
        
        <v-row class="habbit-holder">
            <v-col>
                <HabbitDrawer />
            </v-col>
        </v-row>

    </v-container>
</template>

<script>
import HabbitDrawer from '../components/Habbits/HabbitDrawer'

export default {
    name: "Habbits",
    components: {
        HabbitDrawer
    }
}
</script>

<style scoped>

.nav-row {
    padding-left: 120px;
    padding-top: 50px;
}

.habbit-holder {
    padding: 0px 15px 50px 50px;
}

</style>