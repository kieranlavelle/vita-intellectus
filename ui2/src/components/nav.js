import React, { useCallback, useEffect } from 'react';
import { useHistory } from 'react-router-dom'

import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Box from '@material-ui/core/Box'
import Button from '@material-ui/core/Button'

import useStickyState from '../state/store'

const useStyles = makeStyles((theme) => ({
    root: {
      flexGrow: 1,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    appBar: {
      'background-color': 'white',
      'color': 'black'
    },
    button: {
      fontWeight: 'bold'
    },
    contentContainer: {
      height: '100%'
    }

  }));

  export default function Navigation(props) {
    const classes = useStyles();

    const history = useHistory()
    const redirect = useCallback((path) => history.push(path), [history]);
    const [token, setToken] = useStickyState('token', '');

    useEffect(() => {
      if (token == '') {
        redirect('/login');
      }
    }, [token, redirect]);

    const childrenWithProps = React.Children.map(props.children, child => {
      // checking isValidElement is the safe way and avoids a typescript error too
      if (React.isValidElement(child)) {
        return React.cloneElement(child, { id:  props.id});
      }
      return child;
    });

    return (
      <div className={classes.contentContainer}>
        <AppBar position="static" className={classes.appBar}>
            <Toolbar className={classes.toolbar}>
                <Button className={classes.button}>Habbits</Button>
                <Box textAlign="right" width='100%' fontWeight='bold'>
                  <Button onClick={() => setToken('')} className={classes.button}>logout</Button>
                </Box>
            </Toolbar>
        </AppBar>
        {childrenWithProps}
      </div>
    )

}