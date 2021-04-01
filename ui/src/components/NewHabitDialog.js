import React, { useState } from 'react'
import PropTypes from 'prop-types';

import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogTitle from '@material-ui/core/DialogTitle';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';

import TextField from '@material-ui/core/TextField';

import AddIcon from '@material-ui/icons/Add';
import { makeStyles } from '@material-ui/core/styles';

import ChipInput from './ChipInput';
import DayPicker from './DayPicker'
import { createHabit } from '../endpoints';

const useStyles = makeStyles((theme) => ({
  chips: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  chip: {
    margin: 2,
  },
  dialogContent: {
    marginTop: theme.spacing(1)
  }
}))

function NewHabitDialog(props){

  const classes = useStyles();
  const [open, setOpen] = useState(false);
  const [days, setDays] = useState([]);
  const [tags, setTags] = useState([]);
  const [habitName, setHabitName] = useState("");

  const hableSubmit = () => {

    const habit = {
      name: habitName,
      days: days,
      tags: tags
    }
    createHabit(props.token, habit)
      .then(response => {
        props.onClose(response.data);
        setOpen(false);
      })
      .catch(error => {
        console.log(error);
        alert('failed to create habit')
      })
  }

  return (
    <div>
      <Button onClick={() => setOpen(true)}>
        New Habit <AddIcon />
      </Button>
      <Dialog
        onClose={() => {setOpen(false)}}
        open={open}
        aria-labelledby="form-dialog-title"
      >
        <DialogTitle id="form-dialog-title">New Habit</DialogTitle>
        <DialogContent className={classes.dialogContent}>
          <DialogContentText>
            Use the form below to add a new habit to your schedule.
          </DialogContentText>
            <TextField
              autoFocus
              variant="outlined"
              color="primary"
              margin="dense"
              id="name"
              type="text"
              label="name"
              onChange={(e) => setHabitName(e.target.value)}
              fullWidth
            />
            <ChipInput onChange={setTags}/>
            <DayPicker updateDays={setDays}/>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => {setOpen(false)}} color="secondary" variant="outlined">
            Close
          </Button>
          <Button onClick={hableSubmit} color="primary" variant="outlined">
            Create
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  )
}

NewHabitDialog.propTypes = {
  onClose: PropTypes.func.isRequired,
};

export default NewHabitDialog;

