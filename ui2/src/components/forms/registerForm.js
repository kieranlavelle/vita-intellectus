import React, { useCallback } from 'react'
import {useHistory} from 'react-router-dom';

import TextField from '@material-ui/core/TextField'
import Button from '@material-ui/core/Button'
import Box from '@material-ui/core/Box'

import { makeStyles } from '@material-ui/core/styles';
import { useForm } from "react-hook-form";

import { AUTH } from '../../http'
import useStickyState from '../../state/store'

const useStyles = makeStyles((theme) => ({
    root: {
      '& .MuiTextField-root': {
        margin: theme.spacing(1),
        width: '100%'
      },
      margin: theme.spacing(2),
    },
    header: {
        color: 'black'
    },
    button: {
        backgroundColor: theme.palette.success.main,
    }
}));


export default function RegisterForm(){

    const {register, handleSubmit, errors} = useForm();
    const classes = useStyles();

    const history = useHistory()
    const redirect = useCallback((path) => history.push(path), [history]);

    const onSubmit = (data) => {
        const formData = new FormData();
        formData.append('username', data.username);
        formData.append('password', data.password);
        formData.append('email', data.email);

        console.log(data)

        AUTH.post("/register", formData)
            .then(response => {
                redirect('/login')
            })
            .catch(error => console.log(error.response.data))
    }

    return (
        <form
            className={classes.root}
            onSubmit={handleSubmit(onSubmit)}
        >
            <h2 className={classes.header}>Sign Up</h2>
            <Box
                display="flex"
                flexDirection="column"
                alignItems="center"
            >
                <TextField
                    label="Username"
                    variant="outlined"
                    name="username"
                    type="text"
                    inputRef={register}
                />
                <TextField
                    label="Email"
                    type="text" 
                    variant="outlined"
                    name="email"
                    inputRef={register}
                />
                <TextField
                    label="Password"
                    type="password" 
                    variant="outlined"
                    name="password"
                    inputRef={register}
                />
                <Button
                    variant="contained"
                    className={classes.button}
                    type="submit"
                >
                    Sign Up
                </Button>
            </Box>
        </form>
    )
}