import React, { useState } from "react"

import { useHistory, Link } from "react-router-dom";

import useMediaQuery from '@material-ui/core/useMediaQuery';
import { makeStyles } from '@material-ui/core/styles';
import Box from '@material-ui/core/Box';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import Alert from '@material-ui/lab/Alert';

import { createMuiTheme, useTheme, ThemeProvider } from '@material-ui/core/styles';

import { AUTH } from '../http'
import usePersistedState from '../utilities'
import loginLogo from '../login.svg';

const theme = createMuiTheme({
  typography: {
    fontFamily: ['"Be Vietnam"', 'sans-serif'].join(',')
  }
});

const useStyles = makeStyles((theme) => ({
  root: {
    height: "100vh",
    backgroundColor: 'white',
    [theme.breakpoints.up('md')]: {
      marginRight: '100px'
    }
  },
  left: {
    [theme.breakpoints.up('md')]: {
      marginLeft: '200px'
    },
    textAlign: 'center'
  },
  leftHeader: {
    fontWeight: 700,
    marginBottom: '100px'
  },
  headers: {
    fontWeight: 500,
    lineHeight: 1.5,
  },
  subHeaders: {
    fontWeight: 500,
    color: 'rgb(99, 115, 129)'
  },
  container: {
  },
  loginForm: {
    padding: "15px",
    backgroundColor: 'white',
    borderRadius: '10px',
    [theme.breakpoints.up('md')]: {
      width: '40%'
    }
  },
  formInput: {
    width: '100%',
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  loginButton: {
    marginTop: '10px',
    width: '100%',
    padding: '10px',
    fontWeight: 700
  },
  errorMessage: {
    color: 'red'
  },

  greenText: {
    color: 'rgb(0, 171, 85)'
  },
  signUpPrompt: {
    fontWeight: 600,
    marginTop: '20px',
    textAlign: 'right',
    [theme.breakpoints.up('md')]: {
      paddingRight: '100px'
    }
  }

}))

function Register() {

  const classes = useStyles();
  const history = useHistory();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [token, setToken] = usePersistedState('token', '');

  const theme = useTheme();
  const showLeft = useMediaQuery(theme.breakpoints.up('md'));

  let register = (e) => {
    e.preventDefault();

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

  function Message() {
    if (errorMessage.length > 0) {
      return <Alert severity="error">{errorMessage}</Alert>
    } else if (loggedIn) {
      console.log('login message')
      return <Alert severity="success">Success! Logging you in.</Alert>;
    } else {
      return "";
    }
  }

  const LeftPanel = () => {
      return (
        <div className={classes.left}>
          <Typography gutterBottom variant="h4" className={classes.leftHeader}>
            Come on in!
          </Typography>
          <img src={loginLogo} />
        </div>
      )
  }

  const RightPanel = () => {
    return (
      <form 
        className={classes.loginForm}
      >
        <Typography gutterBottom variant="h5" className={classes.headers}>
          Sign Up To Vita
        </Typography>
        <Typography gutterBottom variant="p" className={classes.subHeaders}>
          Enter your details below
        </Typography>
        <p className={classes.errorMessage}>{errorMessage}</p>
        <TextField
          label="Email"
          className={classes.formInput}
          onChange={(e) => setEmail(e.target.value)}
          variant="outlined"
          color="primary"
          required
        />
        <TextField
          label="Password"
          className={classes.formInput}
          type="password"
          onChange={(e) => setPassword(e.target.value)}
          variant="outlined"
          color="primary"
          required
        />
        <TextField
          label="Confirm password"
          className={classes.formInput}
          type="password"
          onChange={(e) => setPassword(e.target.value)}
          variant="outlined"
          color="primary"
          required
        />
        <br/>
        <Button
          variant="contained"
          color="primary"
          type="submit"
          onClick={register}
          className={classes.loginButton}
        >
          Sign Up
        </Button>
      </form>
    )
  }

  return (
    <div className={classes.container}>
        <Typography
          className={classes.signUpPrompt}
        >
          Already have an account? <Link to="/login"
            className={classes.greenText}
            style={{textDecoration: 'none'}}
            onClick={() => history.push('/login')}
          >Login.</Link>
        </Typography>
      <Box
        display="flex"
        flexDirection="row"
        className={classes.root}
        justifyContent="space-between"
        alignItems="center"
      >
        {showLeft ? <LeftPanel /> : <span />}
        <RightPanel />
      </Box>
    </div>
  )
}

export default Register;