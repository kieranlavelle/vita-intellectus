import React, { useEffect } from 'react'

import { useHistory } from 'react-router-dom';

import { makeStyles } from '@material-ui/core/styles';
import IconButton from '@material-ui/core/IconButton';
import SettingsIcon from '@material-ui/icons/Settings';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';


import Button from '@material-ui/core/Button';

import usePersistedState from '../utilities'


const useStyles = makeStyles((theme) => ({
  navContainer: {
    marginBottom: theme.spacing(1),
    padding: theme.spacing(1),
    backgroundColor: 'white',
  },
  settings: {
    float: 'right'
  }
}));

const SettingsMenu = () => {

  const ITEM_HEIGHT = 48;
  const [anchorEl, setAnchorEl] = React.useState(null);
  const open = Boolean(anchorEl);

  const classes = useStyles();
  let history = useHistory();
  const [token, setToken] = usePersistedState('token', '');

  useEffect(() => {
    if (token === ''){
      history.push('/login');
    };
  }, [token]);

  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const logout = () => {
    setToken('');
  }

  return (
      <IconButton
        className={classes.settings}
        onClick={handleClick}
        style={{padding: '0px'}}
      >
        <SettingsIcon />
        <Menu
          id="long-menu"
          anchorEl={anchorEl}
          keepMounted
          open={open}
          onClose={handleClose}
          PaperProps={{
            style: {
              maxHeight: ITEM_HEIGHT * 4.5,
              width: '20ch',
            },
          }}
        >
        <MenuItem onClick={logout}>
          Logout
        </MenuItem>
      </Menu>
      </IconButton>
  )
}

function Nav(){
  let classes = useStyles();
  let history = useHistory();

  return (
    <div className={classes.navContainer}>
      <Button
        onClick={history.push('/home')}
      >
        <b>Home</b>
      </Button>
      {/* <Button
        onClick={() => setToken('')}
        style={{float: 'right'}}
      >
        <b>Logout</b>
      </Button> */}
      <SettingsMenu />
    </div>
  )
}

export default Nav;