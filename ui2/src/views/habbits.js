import React, { Component } from 'react';

import TextField from '@material-ui/core/TextField'
import Card from '@material-ui/core/Card'
import Grid from '@material-ui/core/Grid'
import { makeStyles } from '@material-ui/core/styles';


const useStyles = makeStyles((theme) => ({
    container: {
      height: '100%',
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
    }
  }));

export default function Habbits() {
    const classes = useStyles()

    return (
            <Grid
                container
                justify="center"
                alignContent="center"
                className={classes.container}
            >
                <h1>Habbits</h1>
            </Grid>
    )
}