import React from "react"

import {
    BrowserRouter,
    Switch,
    Route,
    Redirect
} from 'react-router-dom';

import Register from '../views/register'
import Login from '../views/login'
import Habits from '../views/habits'
import Habit from '../views/habit'
import Navigation from '../components/nav'

import useStickyState from '../state/store'


function GuardedRoute(props) {
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

    const [token, setToken] = useStickyState("token", '');

    return (
        <BrowserRouter basename={'/vita'}>
            <Switch>
                <Route
                    exact
                    path="/"
                    render={() => {
                        return (
                            token != "" ? <Redirect to='/habits'/> : <Redirect to='/login'/>
                        )
                    }}
                />
                <Route exact path="/login" component={Login} />
                <Route exact path="/register" component={Register} />
                <GuardedRoute exact path="/habits">
                    <Navigation>
                        <Habits />
                    </Navigation>
                </GuardedRoute>
                <GuardedRoute exact path="/habit/:id" children={Habit}>
                    <Navigation>
                        <Habit />
                    </Navigation>
                </GuardedRoute>
            </Switch>
        </BrowserRouter>
    )
}