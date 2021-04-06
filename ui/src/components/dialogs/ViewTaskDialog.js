import React, { useState } from 'react'

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
import { makeStyles } from '@material-ui/core/styles';

import ChipInput from '../sub_components/ChipInput';
import DayPicker from '../sub_components/DayPicker';

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

function ViewTaskDialog(props){

  const classes = useStyles();

  const {
    open,
    setOpen,
    name,
    description,
    days,
    recurring
  } = props;

  const [tags, setTags] = useState([]);
  const [date, setDate] = useState(getDate())

  return (
    <div>
      <Dialog
        onClose={() => {setOpen(false)}}
        open={props.open}
        aria-labelledby="form-dialog-title"
      >
        <DialogTitle id="form-dialog-title">{name}</DialogTitle>
        <DialogContent className={classes.dialogContent}>
          <DialogContentText>
            Use the form below to add a new task.
          </DialogContentText>
            <TextField
              variant="outlined"
              color="primary"
              margin="dense"
              id="name"
              type="text"
              label="name"
              value={name}
              disabled
              fullWidth
            />
            <TextField
              variant="outlined"
              color="primary"
              margin="dense"
              id="description"
              type="text"
              label="description"
              value={description}
              disabled
              fullWidth
            />
            <ChipInput onChange={setTags} disabled/>
            <FormGroup row>
              <FormControlLabel
                control={
                  <Switch
                    checked={recurring}
                    disabled
                    name="recurring"
                    color="primary"
                  />
                }
                label="Recurring Task"
              />
            </FormGroup>
            {recurring ?
                <DayPicker updateDays={()=>{}} selectedDays={days} disabled/> :
                <TextField
                  fullWidth
                  id="date"
                  label="Task Date"
                  type="date"
                  defaultValue={date}
                  disabled
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
        </DialogActions>
      </Dialog>
    </div>
  )
}

export default ViewTaskDialog;

