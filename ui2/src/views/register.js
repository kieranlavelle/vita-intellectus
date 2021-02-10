import React from 'react';

import Card from '@material-ui/core/Card'
import Grid from '@material-ui/core/Grid'
import { makeStyles } from '@material-ui/core/styles';

import background from '../img/login-bg.jpg'
import RegisterForm from '../components/forms/registerForm'


const useStyles = makeStyles((theme) => ({
    container: {
      height: '100%',
      backgroundImage: `url(${background})`,
      backgroundSize: 'cover',
      backgroundRepeat: 'no-repeat'
    },
    registerCard: {
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
    }
  }));

export default function Register() {
    const classes = useStyles()

    return (
            <Grid
                container
                justify="center"
                alignContent="center"
                className={classes.container}
            >
                <Card 
                    className={classes.registerCard}
                    variant="elevation"
                    elevation={3}
                >
                    <RegisterForm />
                </Card>
            </Grid>
    )
}