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
import Habbit from '../views/habbit'
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

        const childrenWithProps = React.Children.map(props.children, child => {
            // checking isValidElement is the safe way and avoids a typescript error too
            if (React.isValidElement(child)) {
              return React.cloneElement(child, { id:  props.computedMatch.params.id});
            }
            return child;
          });

        return (
            childrenWithProps
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
                <GuardedRoute exact path="/habbit/:id" children={Habbit}>
                    <Navigation>
                        <Habbit />
                    </Navigation>
                </GuardedRoute>
            </Switch>
        </BrowserRouter>
    )
}