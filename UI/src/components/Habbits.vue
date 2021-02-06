<template>
    <v-container fluid class="habbits-view">
        <v-row
            @newhabbit="reloadHabits"
        >
            <new-habbit />
        </v-row>
        <v-row>
            <habbit
                v-for="habbit in habbits"
                :key="habbit.habbit_id"
                v-bind="habbit"
                
                class="habbit"
                :class="{isMobile: mobile, isTablet: tablet, isComputer: computer}"
            />
        </v-row>
    </v-container>
</template>

<script>
import { API } from '../endpoints/http-config'

import habbit from '../components/Habbits/Habbit'
import newHabbit from '../components/Habbits/NewHabbit'

export default {
    name: "Habbits",
    components: {
        habbit,
        newHabbit
    },
    data: () => ({
        habbits: [],
        errorGettingHabbits: false,
    }),
    methods: {
        hasHabbits: function() {
            return this.habbits.length > 0;
        },
        reloadHabits: function() {
            API.get("habbits")
            .then(response => (this.habbits = response.data))
            .catch(error => (this.errorGettingHabbits = true))
        }
    },
    mounted: async function(){
        API.get("habbits")
           .then(response => (this.habbits = response.data))
           .catch(error => (this.errorGettingHabbits = true))
    },
    computed: {
        isMobile() {
            return this.$vuetify.breakpoint.name == "xs";
        },
        isTablet() {
            let breakpoint = this.$vuetify.breakpoint.name;
            let tablets = ["sm", "md"]

            return tablets.includes(breakpoint);
        },
        isComputer() {
            let breakpoint = this.$vuetify.breakpoint.name;
            let computers = ["lg", "xl"]

            return computers.includes(breakpoint);
        }
    }
}
</script>

<style scoped>

.habbits-view {
    background-color: rgb(216, 216, 216);
    padding: 25px 25px 50px 50px;
}

.habbit {
    margin: 0px;
}

.computer {
    max-width: 32%;
}

.mobile {
    max-width: 80%;
}

.tablet {
    max-width: 50%;
}


</style>