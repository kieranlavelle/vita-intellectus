import React from "react"
import { makeStyles } from '@material-ui/core/styles';

import {
    BrowserRouter,
    Switch,
    Route,
    Redirect
} from 'react-router-dom';

import Register from '../views/register'
import Login from '../views/login'
import Habbits from '../views/habbits'
import Navigation from '../components/nav'

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

function GuardedRoute(props) {
    const classes = useStyles();
    const [token, setToken] = useStickyState("token", '');

    if (token != "") {
        return (
            props.children
        )
    } else {
        return <Redirect to={{ pathname: '/login' }} />
    }
}

export default function Router(){

    return (
        <BrowserRouter>
            <Switch>
                <Route exact path="/login" component={Login} />
                <Route exact path="/register" component={Register} />
                <GuardedRoute exact path="/habbits">
                    <Navigation>
                        <Habbits />
                    </Navigation>
                </GuardedRoute>
            </Switch>
        </BrowserRouter>
    )
}