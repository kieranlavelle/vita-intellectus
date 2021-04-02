import React, { useEffect, useState } from 'react';
import { useHistory } from "react-router-dom";

import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import { makeStyles } from '@material-ui/core/styles';

import usePersistedState from '../utilities';
import { getHabits, habitInfo } from '../endpoints'

import Nav from './Navbar'
import HomeToolbar from './HomeToolbar'
import HabitCard from './HabitCard'

const useStyles = makeStyles((theme) => ({
  root: {
    height: "100vh",
  },
  habits: {
    padding: theme.spacing(2)
  },
  drawer: {
  }
}))

function Home() {
  const [token, setToken] = usePersistedState('token', '');
  const classes = useStyles();
  const history = useHistory();

  const [habits, setHabits] = useState([]);


  useEffect(() => {
    getHabits(token)
      .then(response => {
        setHabits(response.data.habits);
      })
      .catch(error => {
        console.log(error);
        if (error.response.status === 401) {
          setToken('');
          history.push('/login');
        }
      })
  }, [])

  const onDeleteHabit = (id) => {
    const newHabits = [];
    habits.forEach((val, index) => {
      if (val.id !== id) {
        newHabits.push(val);
      }
    })

    setHabits(newHabits);
  }
  
  const onNewHabit = (habit) => {
    setHabits([...habits, habit]);
  }

  return (
    <div className={classes.root}>
      <Box
        display="flex"
        flex="row"
        className={classes.root}
      >
        <div style={{width: '100%'}} className={classes.habits}>
          <Nav />
          <HomeToolbar token={token} onNewHabit={onNewHabit}/>
          <Grid
            container
          >
            {habits.map(habit => 
              <Grid item lg={4} md={6} sm={6} xs={12} key={habit.id}>
                <HabitCard
                  key={habit.id}
                  token={token}
                  onDeleteHabit={onDeleteHabit}
                  {...habit}
                />
              </Grid>
            )}
          </Grid>
        </div>
      </Box>
    </div>
  )
}

export default Home;