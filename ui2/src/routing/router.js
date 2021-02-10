import React from "react"
import {
    BrowserRouter,
    Switch,
    Route,
    Redirect
} from 'react-router-dom';

import Login from '../views/login'
import Habbits from '../views/habbits'
import Navigation from '../components/nav'

import useStickyState from '../state/store'

function GuardedRoute(props) {

    const [token, setToken] = useStickyState("token", '');
    console.log(token)
    if (token != "") {
        return (
            <div>{props.children}</div>
        )
    } else {
        return <Redirect to={{ pathname: '/login' }} />
    }
}

export default function Router(){

    return (
        <BrowserRouter>
            <Switch>
                <Route exact path="/login">
                    <Login />
                </Route>
                <GuardedRoute exact path="/habbits">
                    <Navigation>
                        <Habbits />
                    </Navigation>
                </GuardedRoute>
            </Switch>
        </BrowserRouter>
    )
}