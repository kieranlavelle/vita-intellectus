import React, { useState, useEffect } from 'react';

import Box from '@material-ui/core/Box'
import { makeStyles } from '@material-ui/core/styles';

import NewHabit from '../components/habits/newHabit'
import Habit from '../components/habits/habit'
import HabitsFilter from '../components/habits/habitsFilter'

import useSynState from '../state/synState'
import useStickyState from '../state/store'
import { API } from '../http'


const useStyles = makeStyles((theme) => ({
    container: {
      height: '100%',
      backgroundColor: 'rgb(216, 216, 216);',
      padding: '10px 25px 10px 25px',
    },
    subMenu: {
        width: '100%',
        paddingBottom: '10px'
    },
    habitContainer: {
        width: '100%'
    }
  }));

export default function Habits() {
    const classes = useStyles()
    const [token, setToken] = useStickyState('token', '');

    const [habits, setHabits] = useState([]);
    const [completedHabits, setCompletedHabits] = useState([]);
    const [notDueHabits, setnotDueHabits] = useState([]);
    const [triggerFetchHabits, settriggerFetchHabits] = useState(0);
    const filters = useSynState(['due', 'not_due', 'completed']);

    const addNewHabit = (habit) => {
        setHabits([...habits, habit]);
    }
    const filterHabits = new_filters => filters.set(new_filters)
    const fetchHabits = () => settriggerFetchHabits(triggerFetchHabits+1)

    const config = {
        headers: {
            Authorization: `Bearer ${token}`
        }
    }

    useEffect(() => {
        API.get("/habits", config).then(response => {
            setHabits(response.data.due);
            setCompletedHabits(response.data.completed);
            setnotDueHabits(response.data.not_due);
        });
        console.log("in effect")
    }, [triggerFetchHabits]);


    const dueHabitsList = habits.map((habit) => (
        <Habit
            className={classes.habit}
            key={habit.id}
            name={habit.name}
            dueDates={habit.due_dates}
            completedToday={habit.completed}
            habitID={habit.id}
            onComplete={fetchHabits}
        />
    ));

    const completedHabitsList = completedHabits.map((habit) => (
        <Habit
            className={classes.habit}
            key={habit.id}
            name={habit.name}
            dueDates={habit.due_dates}
            completedToday={habit.completed}
            habitID={habit.id}
            onComplete={fetchHabits}
        />
    ));

    const notDueHabitsList = notDueHabits.map((habit) => (
        <Habit
            className={classes.habit}
            key={habit.id}
            name={habit.name}
            dueDates={habit.due_dates}
            completedToday={habit.completed}
            habitID={habit.id}
            onComplete={fetchHabits}
        />
    ));
    

    return (
        <Box className={classes.container}>
            <Box display="flex" alignItems="center" justifyContent="flex-end" width='100%' className={classes.subMenu}>
                <HabitsFilter defaultFilter={filters.get()} onUpdate={filterHabits}/>
                <NewHabit onNewHabit={addNewHabit}/>
            </Box>
            <Box
                display="flex"
                flexDirection="row"
                flexWrap="wrap"
                justifyContent="flex-start"
                className={classes.habitContainer}
            >
                {filters.get().includes('due') ? dueHabitsList : ''}
                {filters.get().includes('completed') ? completedHabitsList : ''}
                {filters.get().includes('not_due') ? notDueHabitsList : ''}
            </Box>
        </Box>
    )
}