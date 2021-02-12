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


import { makeStyles } from '@material-ui/core/styles';

import DayPicker from '../dayPicker'
import { API } from '../../http'
import { GetAuthHeaders } from '../../auth'

const useStyles = makeStyles((theme) => ({

}));



export default function NewHabbit(props){

    const [open, setOpen] = useState(false);
    const [createHabbit, setCreateHabbit] = useState(false);
    const [selectedDays, setSelectedDays] = useState();
    const [habbitName, setHabbitName] = useState("");

    const classes = useStyles();
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
        setHabbitName(name);
    }

    const toggleHabbit = () => setCreateHabbit(!createHabbit);

    useEffect(() => {
        if (!createHabbit){
            return
        }

        API.post("/habbits", {
            name: habbitName,
            days: selectedDays
        }, config)
        .then(response => {
            toggleHabbit();
            props.onNewHabbit(response.data);
            handleClose();
        })
        .catch(error => {
            toggleHabbit()
            handleClose();
        })
    }, [habbitName, selectedDays, config, createHabbit, toggleHabbit, props])

    return (
        <div>
            <Button onClick={handleClickOpen}>
                <Add />
                New Habbit
            </Button>
            <Dialog open={open} onClose={handleClose} aria-labelledby="form-dialog-title">
                <DialogTitle id="form-dialog-title">Create a new habbit.</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                    This will create a new habbit that will be added to your habbit schedule. You can
                    alter this by clicking on the habbit in the habbits dashboard.
                    </DialogContentText>
                    <TextField
                        autoFocus
                        margin="dense"
                        id="name"
                        label="Habbit Name"
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
                    <Button onClick={toggleHabbit} color="primary">
                        Create
                    </Button>
                </DialogActions>
            </Dialog>
        </div>
    )
}