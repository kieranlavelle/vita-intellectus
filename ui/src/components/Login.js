import React, { useState } from "react"

import { useHistory, Link } from "react-router-dom";

import { makeStyles } from '@material-ui/core/styles';
import Box from '@material-ui/core/Box';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

import { AUTH } from '../http'
import usePersistedState from '../utilities'

const useStyles = makeStyles((theme) => ({
  root: {
    height: "100vh",
    backgroundColor: "#e0e0e0"
  },
  loginForm: {
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
  loginButton: {
    marginTop: '10px'
  },
  errorMessage: {
    color: 'red'
  }
}))

function Login() {

  const classes = useStyles();
  const history = useHistory();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [token, setToken] = usePersistedState('token', '');

  let login = (event) => {
    event.preventDefault();

    let formData = new FormData();
    formData.append('username', email);
    formData.append('password', password);

    AUTH.post('login', formData)
        .then(response => {
          setToken(response.data.access_token);
          history.push('/home')
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
        className={classes.loginForm}
      >
        <h1>Login</h1>
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
          onChange={(e) => setPassword(e.target.value)}/>
        <Link to="/register">
          Don't have an account? Sign up here.
        </Link>
        <br/>
        <Button
          variant="contained"
          color="primary"
          type="submit"
          onClick={login}
          className={classes.loginButton}
        >
          Login
        </Button>
      </form>
    </Box>
  )
}

export default Login;