<template>
    <v-form>
        <v-text-field
            label="Habbit Name"
            v-model="habbitName"
            required
        ></v-text-field>
        <v-combobox
            outlined
            multiple
            dense
            label="day's"
            :items="days"
            v-model="selectedDays"
        >
        </v-combobox>
        <div class="text-right">
            <v-btn
                text
                @click="close"
            >
                Close
            </v-btn>
            <v-btn
                text
                color="green"
                @click="create"
            >
                Create
            </v-btn>
        </div>
    </v-form>
</template>

<script>
import { API } from '../../endpoints/http-config'

export default {
    name: "newHabbitForm",
    props: [
        "dialog"
    ],
    data: () => ({
        days: [
            "monday",
            "tuesday",
            "wednesday",
            "thursday",
            "friday",
            "saturday",
            "sunday"
        ],
        selectedDays: [],
        habbitName: "",
    }),
    methods: {
        close: function(){
            this.selectedDays = [];
            this.dialog.value = false;
        },
        create: function(){
            this.config = {
                headers: {
                    Authorization: `Bearer ${this.$store.getters.token}`
                }
            }

            API.post("/habbits", {name: this.habbitName, days: this.selectedDays}, this.config)
               .then(response => {
                   this.$emit("add");
               })
            
            this.habbitName = "";
            this.selectedDays = [];
            this.dialog.value = false;
        }
    }
}
</script>