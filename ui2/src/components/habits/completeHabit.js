import { useState } from 'react'

import { green } from '@material-ui/core/colors';
import DoneRounded from '@material-ui/icons/DoneRounded';
import { Box, Button } from "@material-ui/core";
import { makeStyles } from '@material-ui/core/styles';

import { API } from '../../http'
import { GetAuthHeaders } from '../../auth'

const useStyles = makeStyles((theme) => ({
    checkButton: {
        transform: 'scale(1.2)'
    },
    colorCheckButton: {
        color: green[500],
        transform: 'scale(1.3)'
    },
    completeHabit: {
        width: '100%'
    }
}));


export default function CompleteHabit(props){
    const classes = useStyles()
    const [hovered, setHovered] = useState(false);
    const [completed, setCompleted] = useState(props.completed);

    const toggleHovered = () => setHovered(!hovered)

    const config = GetAuthHeaders()

    const completeHabit = () => {
        API.put("/habit/complete", {'habit_id': props.habitID}, config)
           .then(response => {
                setCompleted(true);
                props.onCompleteChange();
           })
    }

    return (
        <Button>
            <DoneRounded
                className={hovered || completed ? classes.colorCheckButton : classes.checkButton}
                onMouseEnter={toggleHovered}
                onMouseLeave={toggleHovered}
                onClick={completeHabit}
            />
        </Button>
    )
}