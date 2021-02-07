<template>
    <v-container fluid class="habbits-view">
        <v-row>
            <new-habbit 
                @add="reloadHabits"
            />
        </v-row>
        <v-row>
            <habbit
                v-for="habbit in habbits"
                :key="habbit.habbit_id"
                v-bind="habbit"
                @delete="deleteHabbit(habbit.habbit_id)"
                class="habbit"
                :class="{mobile: isMobile, tablet: isTablet, computer: isComputer}"
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
        reloadHabits: function() {
            API.get("habbits", this.config)
            .then(response => (this.habbits = response.data))
            .catch(error => (this.errorGettingHabbits = true))
        },
        deleteHabbit: function(habbit_id) {
            this.habbits = this.habbits.filter(h => h.habbit_id != habbit_id)
        }
    },
    beforeMount: async function(){
        this.config = {
            headers: {
                Authorization: `Bearer ${this.$store.getters.token}`
            }
        }

        API.get("habbits", this.config)
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