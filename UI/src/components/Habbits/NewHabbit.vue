<template>
    <v-container fluid class="habbit-menu">
        <v-dialog
            transition="dialog-bottom-transition"
            max-width="600"
        >
            <template v-slot:activator="{ on, attrs }">
                <v-btn
                    v-bind="attrs"
                    v-on="on"
                    target="_blank"
                    text
                >
                    <v-icon>mdi-plus</v-icon>
                    new habit
                </v-btn>
            </template>
            <template v-slot:default="dialog">
                <v-card>
                    <v-toolbar
                        color="primary"
                        dark
                    >
                        Create New Habbit
                    </v-toolbar>
                    <v-card-text>
                        <new-habbit-form :dialog="dialog" />
                    </v-card-text>
                    <!-- <v-card-actions class="justify-end">

                    </v-card-actions> -->
                </v-card>
            </template>
        </v-dialog>
    </v-container>
</template>

<script>
import newHabbitForm from '../Habbits/NewHabbitForm'

export default {
    name: 'newHabbit',
    components: {
        newHabbitForm
    },
    data: () => ({
        selectedDays: []
    }),
    methods: {
        close: function(dialog){
            this.$emit('closeNewHabbit');
            dialog.value = false;
        }
    }
}
</script>