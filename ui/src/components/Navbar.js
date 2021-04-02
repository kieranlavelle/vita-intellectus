import React, { useEffect } from 'react'

import { useHistory } from 'react-router-dom';

import { makeStyles } from '@material-ui/core/styles';
import IconButton from '@material-ui/core/IconButton';
import SettingsIcon from '@material-ui/icons/Settings';
import ClickAwayListener from '@material-ui/core/ClickAwayListener';
import Grow from '@material-ui/core/Grow';
import Paper from '@material-ui/core/Paper';
import Popper from '@material-ui/core/Popper';
import MenuItem from '@material-ui/core/MenuItem';
import MenuList from '@material-ui/core/MenuList';


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

  const [open, setOpen] = React.useState(false);
  const anchorRef = React.useRef(null);

  const classes = useStyles();
  let history = useHistory();
  const [token, setToken] = usePersistedState('token', '');

  useEffect(() => {
    if (token === ''){
      history.push('/login');
    };
  }, [token]);

  // return focus to the button when we transitioned from !open -> open
  React.useEffect(() => {
    if (prevOpen.current === true && open === false) {
      anchorRef.current.focus();
    }

    prevOpen.current = open;
  }, [open]);

  const logout = (event) => {
    handleClose(event);
    setToken('');
  }

  const handleToggle = () => {
    setOpen((prevOpen) => !prevOpen);
  };

  const handleClose = (event) => {
    if (anchorRef.current && anchorRef.current.contains(event.target)) {
      return;
    }

    setOpen(false);
  };

  function handleListKeyDown(event) {
    if (event.key === 'Tab') {
      event.preventDefault();
      setOpen(false);
    }
  }

  // return focus to the button when we transitioned from !open -> open
  const prevOpen = React.useRef(open);

  return (
        <IconButton
          className={classes.settings}
          style={{padding: '0px'}}
          ref={anchorRef}
          aria-controls={open ? 'menu-list-grow' : undefined}
          aria-haspopup="true"
          onClick={handleToggle}
        >
          <SettingsIcon />
          <Popper open={open} anchorEl={anchorRef.current} role={undefined} transition disablePortal>
            {({ TransitionProps, placement }) => (
              <Grow
                {...TransitionProps}
                style={{ transformOrigin: placement === 'bottom' ? 'center top' : 'center bottom' }}
              >
                <Paper>
                  <ClickAwayListener onClickAway={handleClose}>
                    <MenuList autoFocusItem={open} id="menu-list-grow" onKeyDown={handleListKeyDown}>
                      <MenuItem onClick={logout}>Logout</MenuItem>
                    </MenuList>
                  </ClickAwayListener>
                </Paper>
              </Grow>
            )}
          </Popper>
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