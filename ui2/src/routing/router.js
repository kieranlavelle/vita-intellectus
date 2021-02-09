import React from "react"
import {
    BrowserRouter,
    Switch,
    Route,
    Link
} from 'react-router-dom';

import Login from '../views/login'
import Habbits from '../views/Habbits'


export default function Router(){
    return (
        <BrowserRouter>
            <Switch>
                <Route exact path="/login">
                    <Login />
                </Route>
                <Route exact path="/habbits">
                    <Habbits />
                </Route>
            </Switch>
        </BrowserRouter>
    )
}