<template>
    <v-container>
        <single-habbit-view
            v-model="showDialog"
            :name="name"
            :days="days"
            :habbit_id="habbit_id"
            @delete="$emit('delete')"
        />
        <v-card
            outlined
            hover
            @click="showDialog=true"
        >   
            <v-list>
                <v-list-item three-line>
                    <v-list-item-content>
                        <div class="overline mb-4">
                            {{name}}
                        </div>
                    </v-list-item-content>
                    <v-list-item-icon>
                        <v-icon
                            medium
                            :color="completed ? 'light-green' : ''"
                            @click.stop="complete"
                        >
                            mdi-check-bold
                        </v-icon>
                    </v-list-item-icon>
                </v-list-item>
                <v-list-item>
                    <v-list-item-content class="caption">
                        Next Due: {{nextDue()}}
                    </v-list-item-content>
                </v-list-item>
            </v-list>
        </v-card>
    </v-container>
</template>

<script>
import { API } from '../../endpoints/http-config'
import { range } from '../../helpers'

import singleHabbitView from './SingleHabbitView'

export default {
    name: "habbit",
    props: ["name", "habbit_id", "completed_today", "days"],
    components: {
        singleHabbitView
    },
    data: () => ({
        completed: false,
        showDialog: false,
        dayIndex: new Date().getDay(),
        allDays: {
            "monday": 1,
            "tuesday": 2,
            "wednesday": 3,
            "thursday": 4,
            "friday": 5,
            "saturday": 6,
            "sunday": 7
        },
        indexToDay: {
            1: "Monday",
            2: "Tuesday",
            3: "Wednesday",
            4: "Thursday",
            5: "Friday",
            6: "Saturday",
            7: "Sunday"
        }
    }),
    mounted: function(){
        this.completed = this.completed_today;
    },
    methods: {
        complete: function(){

            this.config = {
                headers: {
                    Authorization: `Bearer ${this.$store.getters.token}`
                }
            }

            if (!this.completed) {
                let innerThis = this;
                API.put("/habbits/complete", {"habbit_id": this.habbit_id}, this.config)
                .then(response => innerThis.completed = true)
            }
        },
        nextDue: function(){

            var daysThisWeek = range(Math.max(this.dayIndex+1, 7), 7);
            var pickedDay = ""

            if (this.days == null) {
                return this.completed ? "Tomorrow" : "Today" 
            }

            // Convert the schedule days to indexes
            var scheduleIndexs = this.days.map(x => this.allDays[x])
            scheduleIndexs.sort();

            if (scheduleIndexs.includes(this.dayIndex)){
                if (!this.completed){
                   return "Today";
                }
            }

            daysThisWeek.forEach((day) => {
                if (scheduleIndexs.includes(day)){
                    pickedDay = this.indexToDay[day];
                }
            })

            return pickedDay.length > 0 ? pickedDay : "Next " + this.indexToDay[scheduleIndexs[0]];
        }
    }
}
</script>

<style scoped>

.black {
    color: black;
}

</style>