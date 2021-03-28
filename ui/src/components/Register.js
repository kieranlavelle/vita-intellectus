import React, { useState } from "react"

import { useHistory, Link } from "react-router-dom";

import { makeStyles } from '@material-ui/core/styles';
import Box from '@material-ui/core/Box';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

import { AUTH } from '../http'

const useStyles = makeStyles((theme) => ({
  root: {
    height: "100vh",
    backgroundColor: "#e0e0e0"
  },
  registerForm: {
    padding: "15px",
    textAlign: 'center',
    backgroundColor: 'white',
    borderRadius: '10px',
    [theme.breakpoints.up('md')]: {
      width: '33%'
    }
  },
  formInput: {
    width: '100%',
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
  },
  registerButton: {
    marginTop: '10px'
  },
  errorMessage: {
    color: 'red'
  }
}))

function Register() {

  const classes = useStyles();
  const history = useHistory();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const [errorMessage, setErrorMessage] = useState("");


  let register = (event) => {
    event.preventDefault();

    if (password !== confirmPassword) {
      setErrorMessage("password's don't match");
    }

    let formData = new FormData();
    formData.append('username', email);
    formData.append('email', email);
    formData.append('password', password);

    AUTH.post('register', formData)
        .then(response => {
          history.push('/login')
        })
        .catch(error => {
          setErrorMessage(error.response.data.detail);
        })
  };

  return (
    <Box
      display="flex"
      flexDirection="column"
      className={classes.root}
      justifyContent="center"
      alignItems="center"
    >
      <form 
        className={classes.registerForm}
      >
        <h1>Sign Up</h1>
        <p className={classes.errorMessage}>{errorMessage}</p>
        <TextField
          label="Email"
          className={classes.formInput}
          onChange={(e) => setEmail(e.target.value)}
        />
        <TextField
          label="Password"
          className={classes.formInput}
          type="password"
          onChange={(e) => setPassword(e.target.value)}
        />
        <TextField
          label="Confirm Password"
          className={classes.formInput}
          type="password"
          onChange={(e) => setConfirmPassword(e.target.value)}
        />
        <Link to="/login">
          Already have an account? Click here.
        </Link>
        <br/>
        <Button
          variant="contained"
          color="primary"
          type="submit"
          onClick={register}
          className={classes.registerButton}
        >
          Sign Up
        </Button>
      </form>
    </Box>
  )
}

export default Register;