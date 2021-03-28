import React, { useEffect } from 'react'
import { makeStyles } from '@material-ui/core/styles';

import { useHistory } from 'react-router-dom';

import Button from '@material-ui/core/Button';

import usePersistedState from '../utilities'


const useStyles = makeStyles((theme) => ({
  navContainer: {
    marginBottom: theme.spacing(1),
    padding: theme.spacing(1),
    backgroundColor: 'white'
  }
}));

function Nav(){
  let history = useHistory();
  let classes = useStyles();

  const [token, setToken] = usePersistedState('token', '');

  useEffect(() => {
    if (token === ''){
      history.push('/login');
    };
  }, [token]);

  return (
    <div className={classes.navContainer}>
      <Button
        onClick={history.push('/home')}
      >
        <b>Home</b>
      </Button>

      <Button
        onClick={() => setToken('')}
        style={{float: 'right'}}
      >
        <b>Logout</b>
      </Button>
    </div>
  )
}

export default Nav;