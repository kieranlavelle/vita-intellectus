import React, { useState } from 'react'
import PropTypes from 'prop-types';

import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogTitle from '@material-ui/core/DialogTitle';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';

import FormGroup from '@material-ui/core/FormGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';

import TextField from '@material-ui/core/TextField';
import Switch from '@material-ui/core/Switch';
import AddIcon from '@material-ui/icons/Add';
import { makeStyles } from '@material-ui/core/styles';

import ChipInput from '../sub_components/ChipInput';
import DayPicker from '../sub_components/DayPicker';
import { createTask } from '../../endpoints';

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

const getDate = () => {
  var today = new Date();
  var dd = String(today.getDate()).padStart(2, '0');
  var mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
  var yyyy = today.getFullYear();
  
  today = yyyy + '-' + mm + '-' + dd;
  return today
}

function NewTaskDialog(props){

  const classes = useStyles();

  const [open, setOpen] = useState(false);
  const [days, setDays] = useState([]);
  const [tags, setTags] = useState([]);
  const [name, setName] = useState("");
  const [date, setDate] = useState(getDate())
  const [description, setDescription] = useState("");
  const [recurring, setRecurring] = useState(true);

  const hableSubmit = () => {

    const task = {
      name: name,
      description: description,
      tags: tags,
      days: days,
      date: `${date}T00:00:00.000Z`,
      recurring: recurring,
    }
    createTask(props.token, task)
      .then(response => {
        props.onClose(task);
        setOpen(false);
      })
      .catch(error => {
        console.log(error);
        alert('failed to create task')
      })
  }

  return (
    <div>
      <Button onClick={() => setOpen(true)}>
        New Task <AddIcon />
      </Button>
      <Dialog
        onClose={() => {setOpen(false)}}
        open={open}
        aria-labelledby="form-dialog-title"
      >
        <DialogTitle id="form-dialog-title">New Task</DialogTitle>
        <DialogContent className={classes.dialogContent}>
          <DialogContentText>
            Use the form below to add a new task.
          </DialogContentText>
            <TextField
              autoFocus
              variant="outlined"
              color="primary"
              margin="dense"
              id="name"
              type="text"
              label="name"
              onChange={(e) => setName(e.target.value)}
              fullWidth
            />
            <TextField
              variant="outlined"
              color="primary"
              margin="dense"
              id="description"
              type="text"
              label="description"
              onChange={(e) => setDescription(e.target.value)}
              fullWidth
            />
            <ChipInput onChange={setTags}/>
            <FormGroup row>
            <FormControlLabel
              control={
                <Switch
                  checked={recurring}
                  onChange={(e) => setRecurring(e.target.checked)}
                  name="recurring"
                  color="primary"
                />
              }
              label="Recurring Task"
            />
            </FormGroup>
            {recurring ?
                <DayPicker updateDays={setDays}/> :
                <TextField
                  fullWidth
                  id="date"
                  label="Task Date"
                  type="date"
                  defaultValue={date}
                  onChange={(e) => setDate(e.target.value)}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
            }
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

export default NewTaskDialog;

