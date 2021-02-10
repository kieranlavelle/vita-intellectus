import { useState } from 'react'

import { Box, Card, CardActions, CardContent, Icon, Typography } from "@material-ui/core";
import { green } from '@material-ui/core/colors';
import DoneRounded from '@material-ui/icons/DoneRounded';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
    habbitCard: {
        width: "30%"
    },
    completeHabbit: {
        width: '100%'
    },
    checkButton: {
        transform: 'scale(1.2)'
    },
    colorCheckButton: {
        color: green[500],
        transform: 'scale(1.3)'
    }
  }));

export default function Habbit(props){
    const classes = useStyles();
    const [elevation, setElevation] = useState(5);
    const [hover, setHover] = useState(false);
    const toggleHover = () => setHover(!hover);

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
                <Box textAlign="right" width='100%' className={classes.completeHabbit}>
                    <DoneRounded
                        className={hover ? classes.colorCheckButton : classes.checkButton}
                        onMouseEnter={toggleHover}
                        onMouseLeave={toggleHover}
                    />
                </Box>
            </CardActions>
        </Card>
    )
    return <h1>{props.name}</h1>
}