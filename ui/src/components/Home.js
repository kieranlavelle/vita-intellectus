import { useEffect, useState } from 'react';
import { useHistory } from "react-router-dom";

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
    backgroundColor: "#e0e0e0",
  },
  habitContainer: {
    padding: theme.spacing(2)
  }
}))


function getFullHabits(){}

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

  useEffect(() => {
    console.log(habits)
    setHabits(habits.map((habit, index) => {
      habitInfo(token, habit.id)
        .then(response => {
          return Object.assign({}, habit, response.data.info);
        })
        .catch(error => {
          console.log(error);
        })
    }));
  }, []);

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
      <Nav />
      <HomeToolbar token={token} onNewHabit={onNewHabit}/>
      <Grid
        container
        direction="row"
        justify="flex-start"
        alignItems="flex-start"
        className={classes.habitContainer}
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
  )
}

export default Home;