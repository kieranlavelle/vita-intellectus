import { Box, Card, CardContent, Typography } from "@material-ui/core";
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
    habbitCard: {
        width: "30%"
    }
  }));

export default function Habbit(props){
    const classes = useStyles();

    return (
        <Card variant="outlined" className={classes.habbitCard}>
            <Box textAlign="left">
                <CardContent>
                    <Typography>{props.name}</Typography>
                    <Typography>Next Due:</Typography>
                </CardContent>
            </Box>
        </Card>
    )
    return <h1>{props.name}</h1>
}