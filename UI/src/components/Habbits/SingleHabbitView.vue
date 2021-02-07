<template>
    <v-dialog
      v-model="show"
      fullscreen
      hide-overlay
      transition="dialog-bottom-transition"
    >
      <v-card>
        <v-toolbar
          dark
          color="primary"
        >
          <v-btn
            icon
            dark
            @click.stop="show=false"
          >
            <v-icon>mdi-close</v-icon>
          </v-btn>
          <v-toolbar-title>{{name}}</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-toolbar-items>
            <v-btn
              dark
              text
              @click="editSave()"
            >
              {{editSaveText}}
            </v-btn>
          </v-toolbar-items>
        </v-toolbar>
        <v-list
        >
          <v-list-item>
            <v-select
                v-model="selectedDays"
                :items="allDays"
                attach
                chips
                label="Habbit Days"
                multiple
                :disabled="disableEditing"
            ></v-select>
          </v-list-item>
          <v-list-item class="justify-center">
            <v-btn
                depressed
                color="error"
                @click="remove"
            >
                Delete Habbit
            </v-btn>
          </v-list-item>
        </v-list>
      </v-card>
    </v-dialog>
</template>

<script>
import { API } from '../../endpoints/http-config'

export default {
    name: "singleHabbitView",
    props: {
        value: Boolean,
        name: String,
        days: Array,
        habbit_id: Number
    },
    data: () => ({
        disableEditing: true,
        editSaveText: "Edit",
        selectedDays: [],
        allDays: [
            "monday",
            "tuesday",
            "wednesday",
            "thursday",
            "friday",
            "saturday",
            "sunday"
        ]
    }),
    methods: {
        editSave() {
            this.editSaveText = this.disableEditing ? "Save" : "Edit"
            this.disableEditing = !this.disableEditing;
        },
        remove() {
            API.delete("/habbit/"+this.habbit_id, this.config)
               .then(response => {
                 this.show=false;
                 this.$emit('delete');
               })
               .catch(error => console.log(error))
        }
    },
    mounted(){
        this.selectedDays = this.days;
        this.config = {
          headers: {
            Authorization: `Bearer ${this.$store.getters.token}`
          }
        }
    },
    computed: {
        show: {
            get() {
                return this.value;
            },
            set(value) {
                this.$emit('input', value)
            }
        }
    }
}
</script>