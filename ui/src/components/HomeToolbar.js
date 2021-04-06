import React from 'react';

import { makeStyles } from '@material-ui/core/styles';
import FormGroup from '@material-ui/core/FormGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Switch from '@material-ui/core/Switch';
import Box from '@material-ui/core/Box';


import NewTaskDialog from './dialogs/NewTaskDialog';
import DateCycler from './sub_components/DateCycler';


const useStyles = makeStyles((theme) => ({
  root: {
    color: "#e0e0e0",
    padding: theme.spacing(2),
  }
}))

function HomeToolbar(props){

  const classes = useStyles();
  const {filter, setFilter, date, setDate} = props;

  const onClose = (value) => {
    props.onNew(value);
  }


  return (
    <Box
      className={classes.root}
      display="flex"
      flexDirection="col"
      alignContent="start"
    >
      <Box
        display="flex"
        flexDirection="row"
        justifyContent="space-between"
        width="100%"
      >
        <NewTaskDialog onClose={onClose} token={props.token}/>
        <FormGroup row>
          <FormControlLabel
            control={
              <Switch
                checked={filter == 'all'}
                onChange={() => setFilter(filter === 'due' ? 'all' : 'due')}
                name="filter"
                color="primary"
              />
            }
            label="Show not-due."
          />
        </FormGroup>
      </Box>
      <Box>
        <DateCycler date={date} setDate={setDate}/>
      </Box>
    </Box>
  )
}

export default HomeToolbar;

