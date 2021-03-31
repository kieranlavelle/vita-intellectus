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
import SideDrawer from './Drawer';

const useStyles = makeStyles((theme) => ({
  root: {
    height: "100vh",
    // backgroundColor: "#e0e0e0",
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
      {/* <Nav /> */}
      <Box
        display="flex"
        flex="row"
        className={classes.root}
      >
        <SideDrawer className={classes.drawer}/>
        <div style={{width: '100%'}} className={classes.habits}>
          <Nav />
          <Grid
            container
            direction="row"
            justify="flex-start"
            alignItems="flex-start"
            // className={classes.habits}
          >
            {habits.map(habit => 
              <HabitCard
                key={habit.id}
                token={token}
                onDeleteHabit={onDeleteHabit}
                {...habit}
              />
            )}
          </Grid>
        </div>
      </Box>
      {/* <SideDrawer className={classes.drawer}/> */}
      {/* <HomeToolbar token={token} onNewHabit={onNewHabit}/> */}
    </div>
  )
}

export default Home;