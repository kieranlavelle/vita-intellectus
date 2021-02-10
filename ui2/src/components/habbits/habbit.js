import { useState } from 'react'

import { Box, Card, CardActions, CardContent, Typography } from "@material-ui/core";
import { makeStyles } from '@material-ui/core/styles';

import CompleteHabbit from './completeHabbit'

const useStyles = makeStyles((theme) => ({
    habbitCard: {
        // width: '30%',
        margin: '10px',
        [theme.breakpoints.down('sm')]: {
            width: '90%',
        },
        [theme.breakpoints.up('md')]: {
            width: '40%'
        },
        [theme.breakpoints.up('lg')]: {
            width: '30%',
        }
    },

  }));

export default function Habbit(props){
    const classes = useStyles();
    const [elevation, setElevation] = useState(5);

    const nextDue = props.completedToday ?
                        props.dueDates.next_due_on_completed
                        : props.dueDates.next_due

    return (
        <Card
            variant="elevation"
            elevation={elevation}
            className={classes.habbitCard}
            onMouseEnter={() => setElevation(10)}
            onMouseLeave={() => setElevation(5)}
        >
            <Box textAlign="left" fontFamily="verdana" fontWeight="fontWeightLight">
                <CardContent>
                    <Typography variant="overline">{props.name}</Typography>
                    <Typography variant="subtitle2">Next Due: {nextDue}</Typography>
                </CardContent>
            </Box>
            <CardActions>
                <CompleteHabbit />
            </CardActions>
        </Card>
    )
}