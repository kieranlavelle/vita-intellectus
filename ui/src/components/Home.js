import React from 'react';
import { useHistory } from "react-router-dom";

import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import { makeStyles } from '@material-ui/core/styles';

import usePersistedState from '../utilities';
import { getTasks } from '../endpoints'

import Nav from './Navbar'
import HomeToolbar from './HomeToolbar'
import HabitCard from './TaskCard'

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
  const classes = useStyles();
  const history = useHistory();

  const [token, setToken] = usePersistedState('token', '');
  const [tasks, setTasks] = React.useState([]);
  const [filter, setFilter] = React.useState('due');


  React.useEffect(() => {
    getTasks(token, filter)
      .then(response => {
        setTasks(response.data.tasks);
      })
      .catch(error => {
        if (error.response.status === 401) {
          setToken('');
          history.push('/login');
        }
      })
  }, [filter])

  const onDelete = (id) => {
    const newTasks = [];
    tasks.forEach((val, index) => {
      if (val.id !== id) {
        newTasks.push(val);
      }
    })

    setTasks(newTasks);
  }
  
  const onNew = (task) => {
    let newTasks = tasks;
    newTasks.push(task);

    setTasks(newTasks);
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
          <HomeToolbar token={token} onNew={onNew} filter={filter} setFilter={setFilter}/>
          <Grid
            container
          >
            {
              tasks.map(task => {
                return (
                  <Grid item lg={4} md={6} sm={6} xs={12} key={task.id}>
                    <HabitCard
                      key={task.id}
                      token={token}
                      onDelete={onDelete}
                      {...task}
                    />
                  </Grid>
                )
              })
            }
          </Grid>
        </div>
      </Box>
    </div>
  )
}

export default Home;