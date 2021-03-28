import React, { useState } from 'react';

import Typography from '@material-ui/core/Typography';

import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import CardActions from '@material-ui/core/CardActions';

import IconButton from '@material-ui/core/IconButton';
import Button from '@material-ui/core/Button';
import DeleteIcon from '@material-ui/icons/Delete';
import DoneIcon from '@material-ui/icons/Done';
import EditIcon from '@material-ui/icons/Edit';

import { makeStyles } from '@material-ui/core/styles';

import usePersistedState from '../utilities';
import { completeHabit } from '../endpoints';

const useStyles = makeStyles((theme) => ({
  habitCard: {
    width: "32%",
    backgroundColor: "white",
    margin: theme.spacing(1)
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
  }
}))

function HabitCard(props){

  const classes = useStyles();
  const [elevation, setElevation] = useState(5);
  const [token, setToken] = usePersistedState('token', '');

  const [hoverEdit, setHoverEdit] = useState(false);
  const [hoverTick, setHoverTick] = useState(false);
  const [hoverDelete, sethoverDelete] = useState(false);

  return (
    <Card
      variant="elevation"
      className={classes.habitCard}
      elevation={elevation}
      onMouseEnter={() => setElevation(10)}
      onMouseLeave={() => setElevation(5)}
    >
      <CardHeader subheader={props.name} />
      <CardContent>
        <Typography gutterBottom>
          Streak: ...
        </Typography>
      </CardContent>
      <CardActions style={{float: 'right'}}>
        <IconButton
          className={classes['MuiIconButton-root']}
          onMouseEnter={() => sethoverDelete(true)}
          onMouseLeave={() => sethoverDelete(false)}
        >
          <DeleteIcon
            className={hoverDelete ? classes.hoverDelete : classes.delete}
          />
        </IconButton>
        <IconButton
          className={classes['MuiIconButton-root']}
          onMouseEnter={() => setHoverTick(true)}
          onMouseLeave={() => setHoverTick(false)}
        >
          <DoneIcon
            className={(hoverTick || props.completed) ? classes.hoverTick : classes.tick}
          />
        </IconButton>
        <IconButton
          className={classes['MuiIconButton-root']}
          onMouseEnter={() => setHoverEdit(true)}
          onMouseLeave={() => setHoverEdit(false)}
        >
          <EditIcon
            className={hoverEdit ? classes.hoverEdit : classes.edit}
          />
        </IconButton>
      </CardActions>
    </Card>
  )
}

export default HabitCard;