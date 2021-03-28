import { makeStyles } from '@material-ui/core/styles';

import NewHabitDialog from './NewHabitDialog'

const useStyles = makeStyles((theme) => ({
  root: {
    color: "#e0e0e0",
    padding: theme.spacing(2),
  }
}))

function HomeToolbar(props){

  const classes = useStyles();

  const onClose = (value) => {
    console.log(value);
  }

  return (
    <div className={classes.root}>
      <NewHabitDialog onClose={onClose} token={props.token}/>
    </div>
  )
}

export default HomeToolbar;

