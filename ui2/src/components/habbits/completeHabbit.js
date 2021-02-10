import { useState } from 'react'

import { green } from '@material-ui/core/colors';
import DoneRounded from '@material-ui/icons/DoneRounded';
import { Box } from "@material-ui/core";
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
    checkButton: {
        transform: 'scale(1.2)'
    },
    colorCheckButton: {
        color: green[500],
        transform: 'scale(1.3)'
    },
    completeHabbit: {
        width: '100%'
    }
  }));

export default function CompleteHabbit(){
    const classes = useStyles()
    const [hovered, setHovered] = useState(false);
    const toggleHovered = () => setHovered(!hovered)

    return (
        <Box textAlign="right" width='100%' className={classes.completeHabbit}>
            <DoneRounded
                className={hovered ? classes.colorCheckButton : classes.checkButton}
                onMouseEnter={toggleHovered}
                onMouseLeave={toggleHovered}
            />
        </Box>
    )


}