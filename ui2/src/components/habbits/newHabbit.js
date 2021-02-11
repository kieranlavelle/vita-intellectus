import React from 'react';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import Button from '@material-ui/core/Button'
import Add from '@material-ui/icons/Add';

import { makeStyles } from '@material-ui/core/styles';


const days = [
    'monday',
    'tusday',
    'wednesday',
    'thursday',
    'friday',
    'saturday',
    'sunday'
  ];

const useStyles = makeStyles((theme) => ({

}));



export default function NewHabbit(){

    const [open, setOpen] = React.useState(false);
    const [selectedDays, setSelectedDays] = React.useState([]);
    const classes = useStyles();

    const handleClickOpen = () => {
      setOpen(true);
    };
  
    const handleClose = () => {
      setOpen(false);
    };

    const handleChange = (event) => {
        setSelectedDays(event.target.value);
    };

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
                    />
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose} color="primary">
                    Cancel
                    </Button>
                    <Button onClick={handleClose} color="primary">
                        Create
                    </Button>
                </DialogActions>
            </Dialog>
        </div>
    )
}