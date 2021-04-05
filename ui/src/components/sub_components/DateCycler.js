import 'date-fns';
import { format } from 'date-fns';
import React from 'react';
import Grid from '@material-ui/core/Grid';
import DateFnsUtils from '@date-io/date-fns';
import {
  MuiPickersUtilsProvider,
  KeyboardDatePicker,
} from '@material-ui/pickers';

export default function DateCycler(props) {

  const {date, setDate} = props;

  const handleDateChange = (date) => {
    setDate(format(date, 'yyyy-MM-dd'));
  };

  return (
    <MuiPickersUtilsProvider utils={DateFnsUtils}>
      <Grid container justify="flex-start">
        <KeyboardDatePicker
          disableToolbar
          variant="inline"
          inputVariant="outlined"
          format="yyyy-MM-dd"
          margin="dense"
          id="date-picker-inline"
          label="Date picker inline"
          value={date}
          onChange={(date) => {handleDateChange(date)}}
          KeyboardButtonProps={{
            'aria-label': 'change date',
          }}
        />
      </Grid>
    </MuiPickersUtilsProvider>
    )
}