import React, { useState, useEffect } from 'react';

import Typography from '@material-ui/core/Typography';

import Chip from '@material-ui/core/Chip';
import Box from '@material-ui/core/Box';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardHeader from '@material-ui/core/CardHeader';

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
      width: "95%"
    },
    backgroundColor: "white",
    margin: theme.spacing(1),
    padding: theme.spacing(0.5)
  },
  hoverDelete: {
    color: 'red',
    scale: 1.5
  },
  delete: {
    scale: 1
  },
  hoverTick: {
    color: 'green',
    scale: 1.5
  },
  tick: {
    scale: 1
  },
  hoverEdit: {
    color: 'orange',
    scale: 1.5
  },
  edit: {
    scale: 1
  },
  'MuiIconButton-root': {
    padding: '0px'
  },
  statistics: {
    width: '100%',
    marginBottom: '10px'
  },
  statsTitle: {
    color: "rgb(99, 115, 129)",
    fontWeight: 200
  },
  habitTitle: {
    fontWeight: 400,
    color: "rgb(99, 115, 129)",
    marginBottom: '10px'
  },
  tagChip: {
    margin: '5px'
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
  const [tags, setTags] = useState(props.tags);

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
    setTags(habit.tags);
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

  const Actions = () => {

    return (
        <CardActions className={classes.actions}>
          <Box
            display="flex"
            justifyContent="space-evenly"
            width='100%'
          >
            <div>
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
            </div>
            <div>
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
            </div>
            <div>
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
                currentDays={props.days}
                currentTags={tags}
                name={name}
                onEdit={onEdit}
                {...props}
              />
            </div>
          </Box>
      </CardActions>
    )
  }

  return (
    <Card
      variant="elevation"
      className={classes.habitCard}
      elevation={elevation}
      onMouseEnter={() => setElevation(10)}
      onMouseLeave={() => setElevation(5)}
    >
      <CardHeader
        title={
          <Typography
            className={classes.habitTitle}
            variant="h6"
          >
            {name}
          </Typography>
        }
        action={
          tags.slice(0, 2).map((value, index) => {
            return <Chip
                    className={classes.tagChip}
                    key={index}
                    size="small"
                    color="primary"
                    variant="outlined"
                    label={value}
                  />
          })
        }
      />
      <Box display="flex" justifyContent="space-evenly" className={classes.statistics}>
        <div>
          <Typography
            style={{fontWeight: 600}}
            variant="h4"
            color="primary"
          >
            {streak}
          </Typography>
          <Typography className={classes.statsTitle}>Streak</Typography>
        </div>
        <div>
          <Typography
            style={{fontWeight: 600}}
            variant="h4"
            color="primary"
          >
            {Math.round(percentage*100)}%
          </Typography>
          <Typography className={classes.statsTitle}>28 day</Typography>
        </div>
      </Box>
      <Actions/>
    </Card>
  )
}

export default HabitCard;