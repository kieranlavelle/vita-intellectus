import { useState } from 'react'
import PropTypes from 'prop-types';

import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogTitle from '@material-ui/core/DialogTitle';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';

import TextField from '@material-ui/core/TextField';

import { makeStyles } from '@material-ui/core/styles';

import DayPicker from './DayPicker'
import { editHabit } from '../endpoints';

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

function EditHabitDialog(props){

  const classes = useStyles();
  const {open, setOpen, currentDays, name, onEdit} = props;

  const [days, setDays] = useState(currentDays ? currentDays : []);
  const [habitName, setHabitName] = useState(name);

  const onSubmit = () => {

    const habit = {
      name: habitName
    }
    editHabit(props.token, habit, props.id)
      .then(response => {
        onEdit(habit);
        setOpen(false);
      })
      .catch(error => {
        console.log(error);
        alert('failed to edit habit')
      })
  }

  return (
    <div>
      <Dialog
        open={open}
        aria-labelledby="form-dialog-title"
      >
        <DialogTitle id="form-dialog-title">Edit Habit</DialogTitle>
        <DialogContent className={classes.dialogContent}>
          <DialogContentText>
            Use the form below to alter a pre-existing habit.
          </DialogContentText>
            <TextField
              autoFocus
              margin="dense"
              id="name"
              type="text"
              label="name"
              value={habitName}
              onChange={(e) => setHabitName(e.target.value)}
              fullWidth
            />
            <DayPicker disabled={true} updateDays={setDays} selectedDays={currentDays}/>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => {setOpen(false)}} color="secondary" variant="outlined">
            Close
          </Button>
          <Button onClick={onSubmit} color="primary" variant="outlined">
            Edit
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  )
}

EditHabitDialog.propTypes = {
  onEdit: PropTypes.func.isRequired,
};

export default EditHabitDialog;

