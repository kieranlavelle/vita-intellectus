import React, { useState, useEffect } from 'react';
import { useHistory } from 'react-router-dom'

import Grid from '@material-ui/core/Grid'
import Box from '@material-ui/core/Box'
import { makeStyles } from '@material-ui/core/styles';

import NewHabbit from '../components/habbits/newHabbit'
import Habbit from '../components/habbits/habbit'

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
    },
    habbit: {
        // minWidth: '30%'
    }
  }));

export default function Habbits() {
    const classes = useStyles()
    const [token, setToken] = useStickyState('token', '');
    const [habbits, setHabbits] = useState([]);

    const config = {
        headers: {
            Authorization: `Bearer ${token}`
        }
    }

    useEffect(() => {
        API.get("/habbits", config).then(response => {
            setHabbits(response.data);
            console.log(response.data)
        })
      }, []);

    return (
        <Box className={classes.container}>
            <Box textAlign="right" width='100%' className={classes.subMenu}>
                <NewHabbit />
            </Box>
            <Box
                display="flex"
                flexDirection="row"
                flexWrap="wrap"
                justifyContent="space-between" 
                // display="flex"
                // flexWrap="wrap"
                // justifyContent="space-between" 
                // flexDirection="row"
                className={classes.habbitContainer}
            >
                {habbits.map((habbit) => (
                    <Habbit
                        className={classes.habbit}
                        key={habbit.habbit_id}
                        name={habbit.name}
                        dueDates={habbit.due_dates}
                        completedToday={habbit.completed_today}
                    />
                ))}
            </Box>
        </Box>
    )
}