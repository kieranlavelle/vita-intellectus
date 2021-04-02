import React, { useEffect } from "react"
import {
    BrowserRouter,
    Switch,
    Route,
    Redirect
} from "react-router-dom"

import usePersistedState from '../utilities'
import Home from './Home'
import Login from './Login'
import Register from './Register'

function ProtectedRoute(props){
    let authorized = (
        <Route {...props}>
            {props.children}
        </Route>
    )

    let show = false;
    useEffect(() => {
        if (props.token === "") {
            return <Redirect to="/login" />
        }
    }, [props.token])

    return authorized
}

function Router() {

    const [token, setToken] = usePersistedState('token', '');

    return (
        <div>
            <BrowserRouter>
                <Switch>
                    <ProtectedRoute path="/" exact token={token}>
                        <Home />
                    </ProtectedRoute>
                    <ProtectedRoute path="/home" token={token}>
                        <Home />
                    </ProtectedRoute>
                    <Route path="/login">
                        <Login />
                    </Route>
                    <Route path="/register">
                        <Register />
                    </Route>
                </Switch>
            </BrowserRouter>
        </div>
    )
}

export default Router;