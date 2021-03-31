import React, { useState, useEffect } from 'react';

import Typography from '@material-ui/core/Typography';

import Grid from '@material-ui/core/Grid';

import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';

import IconButton from '@material-ui/core/IconButton';
import DeleteIcon from '@material-ui/icons/Delete';
import DoneIcon from '@material-ui/icons/Done';
import EditIcon from '@material-ui/icons/Edit';

import { makeStyles } from '@material-ui/core/styles';

import EditHabitDialog from './EditHabitDialog'
import { completeHabit, deleteHabit, habitInfo } from '../endpoints';


const useStyles = makeStyles((theme) => ({
  habitCard: {
    [theme.breakpoints.down('sm')]: {
      width: "98%"
    },
    [theme.breakpoints.up('sm')]: {
      width: "98%"
    },
    [theme.breakpoints.up('md')]: {
      width: "48%"
    },
    [theme.breakpoints.up('lg')]: {
      width: "30%"
    },
    backgroundColor: "white",
    margin: theme.spacing(1),
    padding: theme.spacing(1)
  },
  hoverDelete: {
    color: 'red',
    scale: 1.2
  },
  delete: {
    scale: 1
  },
  hoverTick: {
    color: 'green',
    scale: 1.2
  },
  tick: {
    scale: 1
  },
  hoverEdit: {
    color: 'orange',
    scale: 1.2
  },
  edit: {
    scale: 1
  },
  'MuiIconButton-root': {
    padding: '0px'
  },
  statsTitle: {
    color: "Gray",
  },
  statistics: {
    fontWeight: 'bold',
    color: 'DimGray'
  }
}))

function HabitCard(props){

  const classes = useStyles();
  const [elevation, setElevation] = useState(5);

  const [hoverEdit, setHoverEdit] = useState(false);
  const [hoverTick, setHoverTick] = useState(false);
  const [hoverDelete, sethoverDelete] = useState(false);

  const [editHabit, setEditHabit] = useState(false);

  const [name, setName] = useState(props.name);
  const [completed, setCompleted] = useState(props.completed);

  const [streak, setStreak] = useState(0);
  const [consecutive, setConsecutive] = useState(0);
  const [percentage, setPercentage] = useState(0);

  useEffect(() => {
    habitInfo(props.token, props.id)
      .then(response => {
        setStreak(response.data.info.streak);
        setConsecutive(response.data.info.consecutive);
        setPercentage(response.data.info['28_day_percent']);
      })
      .catch(error => {
        console.log(`Error for habit id: ${props.id}`)
        console.log(error);
      })
  }, []);


  const onEdit = (habit) => {
    setName(habit.name);
  }

  const onComplete = () => {
    completeHabit(props.token, props.id).then(
      response => {
        setCompleted(true);
      }
    )
    .catch(error => {
      alert("failed to complete habit.")
    })
  }

  const onDelete = () => {
    deleteHabit(props.token, props.id).then(
      response => {
        props.onDeleteHabit(props.id);
      }
    )
    .catch(error => {
      alert("failed to complete habit.")
    })
  }

  return (
    <Card
      variant="elevation"
      className={classes.habitCard}
      elevation={elevation}
      onMouseEnter={() => setElevation(10)}
      onMouseLeave={() => setElevation(5)}
    >
      <Typography variant="overline" style={{fontSize: 14, fontWeight: 'bold'}}>
        {name}
      </Typography>
        <Grid container spacing={3}>
          <Grid item lg={4}>
            <Typography style={{textAlign: 'center', fontSize: 12}}>
              <span className={classes.statsTitle}>Streak</span>
              <br/>
              <span className={classes.statistics}>{streak}</span>
            </Typography>
          </Grid>
          <Grid item lg={4}>
            <Typography style={{textAlign: 'center'}}>
              <span className={classes.statsTitle}>Consecutive</span>
              <br/>
              <span className={classes.statistics}>{consecutive}</span>
            </Typography>
          </Grid>
          <Grid item lg={4}>
            <Typography style={{textAlign: 'center'}}>
              <span className={classes.statsTitle}>1M Percent</span>
              <br/>
              <span className={classes.statistics}>{Math.round(percentage*100)}</span>
            </Typography>
          </Grid>
        </Grid>
      <CardActions style={{float: 'right'}}>
        <IconButton
          className={classes['MuiIconButton-root']}
          onMouseEnter={() => sethoverDelete(true)}
          onMouseLeave={() => sethoverDelete(false)}
          onClick={onDelete}
        >
          <DeleteIcon
            className={hoverDelete ? classes.hoverDelete : classes.delete}
          />
        </IconButton>
        <IconButton
          className={classes['MuiIconButton-root']}
          onMouseEnter={() => setHoverTick(true)}
          onMouseLeave={() => setHoverTick(false)}
          onClick={onComplete}
        >
          <DoneIcon
            className={(hoverTick || completed) ? classes.hoverTick : classes.tick}
          />
        </IconButton>
        <IconButton
          className={classes['MuiIconButton-root']}
          onMouseEnter={() => setHoverEdit(true)}
          onMouseLeave={() => setHoverEdit(false)}
          onClick={() => setEditHabit(true)}
        >
          <EditIcon
            className={hoverEdit ? classes.hoverEdit : classes.edit}
          />
        </IconButton>
        <EditHabitDialog
          open={editHabit}
          setOpen={setEditHabit}
          currentDays={['monday']}
          name={name}
          onEdit={onEdit}
          {...props}
        />
      </CardActions>
    </Card>
  )
}

export default HabitCard;