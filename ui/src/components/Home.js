import { useEffect, useState } from 'react';

import Grid from '@material-ui/core/Grid';
import { makeStyles } from '@material-ui/core/styles';

import usePersistedState from '../utilities';
import { getHabits } from '../endpoints'

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


function Home() {
  const [token, setToken] = usePersistedState('token', '');
  const classes = useStyles();

  const [habits, setHabits] = useState([]);


  useEffect(() => {
    getHabits(token).then(response => {
      response.data.habits.forEach(function (habit, index){
        setHabits([...habits, habit])
      })
    }).catch(error => {
      console.log(error);
    })
  }, [])
  

  return (
    <div className={classes.root}>
      <Nav />
      <HomeToolbar token={token}/>
      <Grid
        container
        direction="row"
        justify="flex-start"
        alignItems="flex-start"
        className={classes.habitContainer}
      >
        {habits.map(habit => <HabitCard key={habit.name} {...habit} />)}
      </Grid>
    </div>
  )
}

export default Home;