import React, { useCallback } from 'react';
import { useHistory } from 'react-router-dom'

import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import Button from '@material-ui/core/Button'

import useStickyState from '../state/store'

const useStyles = makeStyles((theme) => ({
    root: {
      flexGrow: 1,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
    },
    appBar: {
      'background-color': 'white',
      'color': 'black'
    }
  }));

  export default function Navigation(props) {
    const classes = useStyles();

    const history = useHistory()
    const redirect = useCallback(() => history.push('/login'), [history]);
    const [token, setToken] = useStickyState('token', '');

    const logout = () => {
      setToken('');
      redirect()
    }

    return (
      <div>
        <AppBar position="static" className={classes.appBar}>
            <Toolbar>
                <Typography variant="h6" className={classes.title}>
                    Habbits
                </Typography>
                <Button
                  onClick={logout}
                >
                  Logout
                </Button>
            </Toolbar>
        </AppBar>
        {props.children}
      </div>
    )

}

// export default function Navigation(props){

//     console.log(`LoggedIn2: ${props.loggedin}`)

//     if (props.loggedin == true){
//         return <LoggedInNavigation/>
//     }
//     return <Redirect to="/login"/>
// }