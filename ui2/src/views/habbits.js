import React, { useState, useEffect } from 'react';
import { useHistory } from 'react-router-dom'

import Box from '@material-ui/core/Box'
import { makeStyles } from '@material-ui/core/styles';

import NewHabbit from '../components/habbits/newHabbit'
import Habbit from '../components/habbits/habbit'
import HabbitsFilter from '../components/habbits/habbitsFilter'

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
    habbitContainer: {
        width: '100%'
    }
  }));

export default function Habbits() {
    const classes = useStyles()
    const [token, setToken] = useStickyState('token', '');

    const [habbits, setHabbits] = useState([]);
    const [completedHabbits, setCompletedHabbits] = useState([]);
    const [notDueHabbits, setnotDueHabbits] = useState([]);
    const filters = useSynState(['due', 'not_due', 'completed']);

    const addNewHabbit = (habbit) => {
        setHabbits([...habbits, habbit]);
    }

    const filterHabbits = new_filters => filters.set(new_filters)

    const config = {
        headers: {
            Authorization: `Bearer ${token}`
        }
    }

    useEffect(() => {
        API.get("/habbits", config).then(response => {
            setHabbits(response.data.due);
            setCompletedHabbits(response.data.completed);
            setnotDueHabbits(response.data.not_due);
        })
    }, []);


    const dueHabbitsList = habbits.map((habbit) => (
        <Habbit
            className={classes.habbit}
            key={habbit.habbit_id}
            name={habbit.name}
            dueDates={habbit.due_dates}
            completedToday={habbit.completed_today}
            habbitID={habbit.habbit_id}
        />
    ));

    const completedHabbitsList = completedHabbits.map((habbit) => (
        <Habbit
            className={classes.habbit}
            key={habbit.habbit_id}
            name={habbit.name}
            dueDates={habbit.due_dates}
            completedToday={habbit.completed_today}
            habbitID={habbit.habbit_id}
        />
    ));

    const notDueHabbitsList = notDueHabbits.map((habbit) => (
        <Habbit
            className={classes.habbit}
            key={habbit.habbit_id}
            name={habbit.name}
            dueDates={habbit.due_dates}
            completedToday={habbit.completed_today}
            habbitID={habbit.habbit_id}
        />
    ));
    

    return (
        <Box className={classes.container}>
            <Box display="flex" alignItems="center" justifyContent="flex-end" width='100%' className={classes.subMenu}>
                <HabbitsFilter defaultFilter={filters.get()} onUpdate={filterHabbits}/>
                <NewHabbit onNewHabbit={addNewHabbit}/>
            </Box>
            <Box
                display="flex"
                flexDirection="row"
                flexWrap="wrap"
                justifyContent="flex-start"
                className={classes.habbitContainer}
            >
                {filters.get().includes('due') ? dueHabbitsList : ''}
                {filters.get().includes('completed') ? completedHabbitsList : ''}
                {filters.get().includes('not_due') ? notDueHabbitsList : ''}
            </Box>
        </Box>
    )
}