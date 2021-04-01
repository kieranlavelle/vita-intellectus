import React from "react"
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

    if (props.token === "") {
        return <Redirect to="/login" />
    } else {
        return authorized;
    }
}

function Router() {

    const [token, setToken] = usePersistedState('token', '');

    return (
        <div>
            <BrowserRouter basename={process.env.PUBLIC_URL}>
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