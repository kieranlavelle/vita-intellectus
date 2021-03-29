import { useState, useEffect } from 'react'

import { makeStyles } from '@material-ui/core/styles';
import Chip from '@material-ui/core/Chip';
import { Box } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
    chipHolder: {
      
    },

    picker: {
        marginTop: theme.spacing(1),
      },
      chip: {
        margin: theme.spacing(0.5),
      },
}));

export default function DayPicker(props){
    const [days, setDays] = useState(props.selectedDays ? props.selectedDays : []);
    const classes = useStyles();

    const ToggleDay = async day => {
        if (days.includes(day)) {
            setDays(days.filter(iDay => iDay !== day));
        } else {
            setDays([...days, day]);
        }
    }

    // pass the data out of the compoent
    useEffect(() => {props.updateDays(days)}, [props, days]);

    return (
        <Box display="flex" flexWrap="wrap" width="100%" className={classes.picker}>
            <Chip
                variant={days.includes('monday') ? 'default' : 'outlined'}
                className={classes.chip}
                label='Monday'
                size='small'
                color='primary'
                onClick={() => ToggleDay('monday')}
                disabled={true ? props.disabled : false}
            />
            <Chip
                variant={days.includes('tuesday') ? 'default' : 'outlined'}
                className={classes.chip}
                label='Tuesday'
                size='small'
                color='primary'
                onClick={() => ToggleDay('tuesday')}
                disabled={true ? props.disabled : false}
            />
            <Chip
                variant={days.includes('wednesday') ? 'default' : 'outlined'}
                className={classes.chip}
                label='Wednesday'
                size='small'
                color='primary'
                onClick={() => ToggleDay('wednesday')}
                disabled={true ? props.disabled : false}
            />
            <Chip
                variant={days.includes('thursday') ? 'default' : 'outlined'}
                className={classes.chip}
                label='Thursday'
                size='small'
                color='primary'
                onClick={() => ToggleDay('thursday')}
                disabled={true ? props.disabled : false}
            />
            <Chip
                variant={days.includes('friday') ? 'default' : 'outlined'}
                className={classes.chip}
                label='Friday'
                size='small'
                color='primary'
                onClick={() => ToggleDay('friday')}
                disabled={true ? props.disabled : false}
            />
            <Chip
                variant={days.includes('saturday') ? 'default' : 'outlined'}
                className={classes.chip}
                label='Saturday'
                size='small'
                color='primary'
                onClick={() => ToggleDay('saturday')}
                disabled={true ? props.disabled : false}
            />
            <Chip
                variant={days.includes('sunday') ? 'default' : 'outlined'}
                className={classes.chip}
                label='Sunday'
                size='small'
                color='primary'
                onClick={() => ToggleDay('sunday')}
                disabled={true ? props.disabled : false}
            />
        </Box>
    )

}