<template>
    <v-container
        fill-height
        class="d-flex justify-space-between"
    >
            <v-card
                v-for="habbit in habbits"
                :key="habbit.habbit_id"
                class="habbit"
                outlined
                hover
            >
                <v-card-title> {{habbit.name}} </v-card-title>
                <v-card-subtitle>
                    Completed: <b>False</b>
                </v-card-subtitle>
                <v-card-actions>
                    <v-btn
                        color="deep-purple accent-4"
                        class="habbit-done-button"
                        text
                    >
                        done
                    </v-btn>
                </v-card-actions>
            </v-card>
    </v-container>
</template>

<script>
import {getHabbits} from "../../endpoints/habbits"

export default {
    name: "HabbitDrawer",
    data: () => ({
        habbits: [],
    }),
    methods: {
        hasHabbits: function() {
            return this.habbits.length > 0;
        },
    },
    created: async function(){
        let innerThis = this;
        getHabbits().then((res) => {
            console.log(res);
            innerThis.habbits = res;
        })
        console.log(this.habbits);
    }
}
</script>

<style scoped>

.habbit {
    padding: 5px;
    margin-top: 30px;
    width: 30%;
    text-align: center;
}

.habbit-done-button {
    text-align: right;
}

</style>