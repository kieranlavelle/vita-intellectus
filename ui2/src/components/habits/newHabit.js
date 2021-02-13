import React from 'react';
import { useState, useEffect } from 'react'
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import Button from '@material-ui/core/Button'
import Add from '@material-ui/icons/Add';

import DayPicker from '../dayPicker'
import { API } from '../../http'
import { GetAuthHeaders } from '../../auth'


export default function NewHabit(props){

    const [open, setOpen] = useState(false);
    const [createHabit, setCreateHabit] = useState(false);
    const [selectedDays, setSelectedDays] = useState();
    const [habitName, setHabitName] = useState("");

    const config = GetAuthHeaders();

    const handleClickOpen = () => {
      setOpen(true);
    };
  
    const handleClose = () => {
      setOpen(false);
    };

    const updateDays = (days) => {
        setSelectedDays(days);
    };

    const setName = (name) => {
        setHabitName(name);
    }

    const toggleHabit = () => setCreateHabit(!createHabit);

    useEffect(() => {
        if (!createHabit){
            return
        }

        API.post("/habit", {
            name: habitName,
            days: selectedDays
        }, config)
        .then(response => {
            toggleHabit();
            props.onNewHabit(response.data);
            handleClose();
        })
        .catch(error => {
            toggleHabit()
            handleClose();
        })
    }, [habitName, selectedDays, config, createHabit, toggleHabit, props])

    return (
        <div>
            <Button onClick={handleClickOpen}>
                <Add />
                New Habit
            </Button>
            <Dialog open={open} onClose={handleClose} aria-labelledby="form-dialog-title">
                <DialogTitle id="form-dialog-title">Create a new habit.</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                    This will create a new habit that will be added to your habit schedule. You can
                    alter this by clicking on the habit in the habits dashboard.
                    </DialogContentText>
                    <TextField
                        autoFocus
                        margin="dense"
                        id="name"
                        label="Habit Name"
                        type="text"
                        fullWidth
                        onChange={event => setName(event.target.value)}
                    />
                    <DayPicker updateDays={updateDays}/>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose} color="primary">
                        Cancel
                    </Button>
                    <Button onClick={toggleHabit} color="primary">
                        Create
                    </Button>
                </DialogActions>
            </Dialog>
        </div>
    )
}