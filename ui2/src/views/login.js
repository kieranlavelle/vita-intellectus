import React, { useCallback } from 'react';

import { useHistory } from 'react-router-dom';

import Card from '@material-ui/core/Card'
import Grid from '@material-ui/core/Grid'
import { makeStyles } from '@material-ui/core/styles';

import background from '../img/login-bg.jpg'
import LoginFrom from '../components/forms/loginForm'
import { CardContent } from '@material-ui/core';


const useStyles = makeStyles((theme) => ({
    container: {
      height: '100%',
      backgroundImage: `url(${background})`,
      backgroundSize: 'cover',
      backgroundRepeat: 'no-repeat'
    },
    loginCard: {
        color: "white",
        opacity: 0.9,
        [theme.breakpoints.down('sm')]: {
            width: '90%',
        },
        [theme.breakpoints.up('md')]: {
            width: '50%'
        },
        [theme.breakpoints.up('lg')]: {
            width: '30%',
        },
    },
    text: {
        color: 'black'
    },
    link: {
        color: "blue"
    }
  }));

export default function Login() {
    const classes = useStyles()
    const history = useHistory()
    const redirect = useCallback((path) => history.push(path), [history]);

    return (
            <Grid
                container
                justify="center"
                alignContent="center"
                className={classes.container}
            >
                <Card 
                    className={classes.loginCard}
                    variant="elevation"
                    elevation={3}
                >
                    <CardContent>
                        <LoginFrom />
                        <p className={classes.text}>Not Registered? 
                        <a className={classes.link} onClick={() => redirect('/register')}> Sign Up!</a>
                        </p>
                    </CardContent>
                </Card>
            </Grid>
    )
}