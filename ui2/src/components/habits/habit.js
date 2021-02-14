import { useState, useCallback } from 'react'

import { useHistory } from 'react-router-dom'

import { Box, Card, CardActionArea, CardActions, CardContent, Typography } from "@material-ui/core";
import { makeStyles } from '@material-ui/core/styles';

import CompleteHabit from './completeHabit'

const useStyles = makeStyles((theme) => ({
    habitCard: {
        margin: '1%',
        [theme.breakpoints.down('sm')]: {
            width: '100%',
        },
        [theme.breakpoints.up('md')]: {
            width: '47%'
        },
        [theme.breakpoints.up('lg')]: {
            width: '31%',
        }
    },

  }));

export default function Habit(props){
    const classes = useStyles();
    const [elevation, setElevation] = useState(5);
    const [showDialog, setShowDialog] = useState(false);

    const [nextDue, setNextDue] = useState(props.completedToday ?
        props.dueDates.next_due_on_completed
        : props.dueDates.next_due);

    const history = useHistory()
    const redirect = useCallback((path) => history.push(path), [history]);

    const onCompleted = () => {
        setNextDue(props.dueDates.next_due_on_completed);
        props.onComplete();
        console.log("In habit onCompleted")
    };


    return (
        <Card
            variant="elevation"
            elevation={elevation}
            className={classes.habitCard}
            onMouseEnter={() => setElevation(10)}
            onMouseLeave={() => setElevation(5)}
        >
            <CardActionArea onClick={() => redirect(`/habit/${props.habitID}`)}>
                <Box textAlign="left" fontFamily="verdana" fontWeight="fontWeightLight">
                    <CardContent>
                        <Typography variant="overline">{props.name}</Typography>
                        <Typography variant="subtitle2">Next Due: {nextDue}</Typography>
                    </CardContent>
                </Box>
            </CardActionArea>

            <Box display="flex" justifyContent="flex-end">
                <CardActions>
                    <CompleteHabit
                        habitID={props.habitID}
                        completed={props.completedToday}
                        dueDates={props.dueDates}
                        onCompleted={onCompleted}
                    />
                </CardActions>
            </Box>
        </Card>
    )
}