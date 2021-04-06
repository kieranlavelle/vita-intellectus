import React, { useState, useEffect } from 'react';

import Typography from '@material-ui/core/Typography';

import Chip from '@material-ui/core/Chip';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardHeader from '@material-ui/core/CardHeader';



import IconButton from '@material-ui/core/IconButton';
import DeleteIcon from '@material-ui/icons/Delete';
import DoneIcon from '@material-ui/icons/Done';
import EditIcon from '@material-ui/icons/Edit';

import { makeStyles } from '@material-ui/core/styles';

import ViewTaskDialog from './dialogs/ViewTaskDialog';
import EditHabitDialog from './dialogs/EditHabitDialog';
import { completeTask, deleteTask } from '../endpoints';


const useStyles = makeStyles((theme) => ({
  habitCard: {
    width: '95%',
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
    margin: '5px',
  }
}))


function TaskCard(props){

  const classes = useStyles();
  const [elevation, setElevation] = useState(5);

  const [hoverEdit, setHoverEdit] = useState(false);
  const [hoverTick, setHoverTick] = useState(false);
  const [hoverDelete, sethoverDelete] = useState(false);

  const [editTask, setEditTask] = useState(false);
  const [viewTask, setViewTask] = useState(false);

  const [name, setName] = useState(props.name);
  const [description, setDescription] = useState(props.description)
  const [state, setState] = useState(props.state);
  const [tags, setTags] = useState(props.tags ? props.tags : []);
  const [days, setDays] = useState(props.days);
  const [date, setDate] = useState(props.days)
  
  const {recurring} = props;

  const streak = 0;
  const percentage = 0

  const stateColor = (state) => {
    if (state == 'not-due') {
      return 'default';
    } else if (state == 'missed') {
      return 'secondary';
    }
    return 'primary'
  }

  const onEdit = (task) => {
    setName(task.name);
    setTags(task.tags);
    setDescription(task.description)
  }

  const onComplete = () => {
    completeTask(props.token, props.id).then(
      response => {
        setState('completed');
      }
    )
    .catch(error => {
      alert("failed to complete habit.")
    })
  }

  const onDelete = () => {
    deleteTask(props.token, props.id).then(
      response => {
        props.onDelete(props.id);
      }
    )
    .catch(error => {
      alert("failed to complete habit.")
    })
  }

  const onClick = (event) => {
    setViewTask(true);
    console.log(viewTask)
  }

  const Actions = () => {

    return (
        <CardActions className={classes.actions}>
          <Grid
            container
            width='100%'
            direction='column'
            alignItems='center'
          >
            <Grid item lg={12} md={12}>
              {tags.map((value, index) => {
                return <Chip
                        className={classes.tagChip}
                        key={index}
                        size="small"
                        color="primary"
                        variant="outlined"
                        label={value}
                      />
              })}
              {tags.length == 0 ? (<Chip
                      className={classes.tagChip}
                      size="small"
                      variant="outlined"
                      label="no-tags"
                    />) : <span></span>
              }
            </Grid>
            <Grid container item lg={12} md={12} sm={12}>
              <Box
                display="flex"
                flexDirection='row'
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
                      className={(hoverTick || state === 'completed') ? classes.hoverTick : classes.tick}
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
                    open={editTask}
                    setOpen={setEditTask}
                    currentDays={days}
                    currentTags={tags}
                    name={name}
                    onEdit={onEdit}
                    {...props}
                  />
                </div>
              </Box>
            </Grid>
          </Grid>
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
      <ViewTaskDialog
        open={viewTask}
        setOpen={setViewTask}
        name={name}
        description={description}
        days={days}
        recurring={recurring}
      />
      <CardHeader
        onClick={onClick}
        title={
          <Typography
            className={classes.habitTitle}
            variant="h6"
          >
            {name}
          </Typography>
        }
        action={
          <Chip
            className={classes.tagChip}
            size="small"
            variant="outlined"
            label={state}
            color={stateColor(state)}
          />
        }
      />
      <Box
        onClick={onClick}
        display="flex"
        justifyContent="space-evenly"
        className={classes.statistics}
      >
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

export default TaskCard;