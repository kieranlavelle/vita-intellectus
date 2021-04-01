import React, { useState, useEffect } from 'react';

import Box from '@material-ui/core/Box';

import Chip from '@material-ui/core/Chip';
import TextField from '@material-ui/core/TextField';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';


const useStyles = makeStyles((theme) => ({
  tagChip: {
    margin: '5px'
  },
  tagHeader: {
    color: 'rgb(99, 115, 129)'
  }
}));

const ChipInput = (props) => {

  const { onChange } = props;
  const classes = useStyles();
  const [value, setValue] = useState('');
  const [tags, setTags] = useState(props.tags ? props.tags : []);

  const handleKeyPress = (e) => {
    if (e.key === 'Enter') {
      setTags([...tags, value]);
      setValue('');
    }
  }

  useEffect(() => {
    onChange(tags);
  }, [tags, setTags]);

  return (
    <div>
      <Box display="flex" alignItems="center">
        {
          tags.length > 0 ? (<Typography className={classes.tagHeader}>Tags:</Typography>) : <span></span>
        }
        {
          tags.map((value, index) => {
            return <Chip
                    className={classes.tagChip}
                    key={index}
                    size="small"
                    color="primary"
                    variant="outlined"
                    label={value}
                  />
          })
        }
      </Box>
      <TextField
        variant="outlined"
        value={value}
        label="tag"
        color="primary"
        margin="dense"
        fullWidth
        onChangeCapture={(e) => setValue(e.target.value)}
        onKeyPress={handleKeyPress}
      />
    </div>
  )
}

export default ChipInput;