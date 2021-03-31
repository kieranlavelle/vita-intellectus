import React from 'react';

import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

const useStyles = makeStyles((theme) => ({
  drawerContainer: {
    border: '1px solid',
    borderColor: 'rgba(145, 158, 171, 0.24)',
    width: '20%',
    height: '100%'
  },
}));

const SideDrawer = () => {

  const classes = useStyles();

  return (
    <div
      className={classes.drawerContainer}
    >
      <List>
        <ListItem button>
          {/* <ListItemIcon>{index % 2 === 0 ? <InboxIcon /> : <MailIcon />}</ListItemIcon> */}
          <ListItemText primary="Dashboard" />
        </ListItem>
      </List>
    </div>
  )
}

export default SideDrawer