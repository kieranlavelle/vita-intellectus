import { useState, useEffect } from 'react'

import { makeStyles } from '@material-ui/core/styles';
import Chip from '@material-ui/core/Chip';
import { Button } from '@material-ui/core';

import useSynState from '../../state/synState'

const useStyles = makeStyles((theme) => ({
  chip: {
    margin: theme.spacing(0.5),
  },
}));

const filterNames = [
    'due',
    'completed',
    'not_due'
]

export default function HabbitFilter(props){
    const classes = useStyles();
    // const [filters, setFilters] = useState(
    //   props.defaultFilter ? props.defaultFilter : []
    // );
    const filters = useSynState(
      props.defaultFilter ? props.defaultFilter : []
    );

    const ToggleFilter = (filter) => {
      if (filters.get().includes(filter)) {
          filters.set(filters.get().filter(ifilter => ifilter != filter));
          props.onUpdate(filters.get());
      } else {
        filters.set([...filters.get(), filter]);
        props.onUpdate(filters.get());
      }
    }

    return (

      <Button>
        <Chip
            variant={filters.get().includes('due') ? 'default' : 'outlined'}
            className={classes.chip}
            label='Due'
            size='small'
            color='primary'
            onClick={() => ToggleFilter('due')}
        />
        <Chip
            variant={filters.get().includes('completed') ? 'default' : 'outlined'}
            className={classes.chip}
            label='Completed'
            size='small'
            color='primary'
            onClick={() => ToggleFilter('completed')}
        />
        <Chip
            variant={filters.get().includes('not_due') ? 'default' : 'outlined'}
            className={classes.chip}
            label='Not Due'
            size='small'
            color='primary'
            onClick={() => ToggleFilter('not_due')}
        />
      </Button>
    )
}
