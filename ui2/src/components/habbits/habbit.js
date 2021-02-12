import { useState, useCallback } from 'react'

import { useHistory } from 'react-router-dom'

import { Box, Card, CardActions, CardContent, Typography } from "@material-ui/core";
import { makeStyles } from '@material-ui/core/styles';

import CompleteHabbit from './completeHabbit'

const useStyles = makeStyles((theme) => ({
    habbitCard: {
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

export default function Habbit(props){
    const classes = useStyles();
    const [elevation, setElevation] = useState(5);
    const [showDialog, setShowDialog] = useState(false);

    const [nextDue, setNextDue] = useState(props.completedToday ?
        props.dueDates.next_due_on_completed
        : props.dueDates.next_due);

    const history = useHistory()
    const redirect = useCallback((path) => history.push(path), [history]);

    const changeNextDue = () => setNextDue(props.dueDates.next_due_on_completed);


    return (
        <Card
            variant="elevation"
            elevation={elevation}
            className={classes.habbitCard}
            onMouseEnter={() => setElevation(10)}
            onMouseLeave={() => setElevation(5)}
            onClick={(e) => {
                e.preventDefault();
                redirect(`/habbit/${props.habbitID}`);
            }}
        >
            <Box textAlign="left" fontFamily="verdana" fontWeight="fontWeightLight">
                <CardContent>
                    <Typography variant="overline">{props.name}</Typography>
                    <Typography variant="subtitle2">Next Due: {nextDue}</Typography>
                </CardContent>
            </Box>
            <CardActions>
                <CompleteHabbit
                    habbitID={props.habbitID}
                    completed={props.completedToday}
                    dueDates={props.dueDates}
                    onCompleteChange={() => changeNextDue()}
                />
            </CardActions>
        </Card>
    )
}