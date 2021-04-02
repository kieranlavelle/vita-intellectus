import React, { useState } from "react"

import { useHistory, Link } from "react-router-dom";

import useMediaQuery from '@material-ui/core/useMediaQuery';
import { makeStyles } from '@material-ui/core/styles';
import Box from '@material-ui/core/Box';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import Alert from '@material-ui/lab/Alert';

import { createMuiTheme, useTheme, ThemeProvider  } from '@material-ui/core/styles';

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
    [theme.breakpoints.up('lg')]: {
      marginRight: '100px'
    }
  },
  left: {
    [theme.breakpoints.up('lg')]: {
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
  loginForm: {
    padding: "15px",
    backgroundColor: 'white',
    borderRadius: '10px',
    [theme.breakpoints.up('lg')]: {
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
  greenText: {
    color: 'rgb(0, 171, 85)'
  },
  signUpPrompt: {
    fontWeight: 600,
    marginTop: '20px',
    [theme.breakpoints.up('md')]: {
      paddingRight: '100px',
      textAlign: 'right',
    },
    [theme.breakpoints.down('md')]: {
      paddingRight: '100px',
      textAlign: 'right',
    }
  }
}))

function Login() {

  const classes = useStyles();
  const history = useHistory();

  const [loggedIn, setLoggedIn] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const [token, setToken] = usePersistedState('token', '');

  const theme = useTheme();
  const showLeft = useMediaQuery(theme.breakpoints.up('md'));


  function Message() {
    if (errorMessage.length > 0) {
      return <Alert severity="error">{errorMessage}</Alert>
    } else if (loggedIn) {
      return <Alert severity="success">Success! Logging you in.</Alert>;
    } else {
      return "";
    }
  }

  const LeftPanel = () => {
      return (
        <div className={classes.left}>
          <Typography gutterBottom variant="h4" className={classes.leftHeader}>
            Welcome Back!
          </Typography>
          <img src={loginLogo} />
        </div>
      )
  }

  const RightPanel = (props) => {

    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const {onSubmit} = props;

    const submitForm = (e, email, password) => {
      e.preventDefault();
      onSubmit(email, password);
    }

    return (
        <form 
          className={classes.loginForm}
        >
          <ThemeProvider theme={theme}>
            <Typography gutterBottom variant="h5" className={classes.headers}>
              Sign In To <span className={classes.greenText}>Vita</span>
            </Typography>
            <Typography gutterBottom className={classes.subHeaders}>
              Enter your details below
            </Typography>
          </ThemeProvider>
          <Message />
          <TextField
            label="Email"
            className={classes.formInput}
            onChangeCapture={(e) => setEmail(e.target.value)}
            value={email}
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
          <Button
            variant="contained"
            color="primary"
            type="submit"
            onClickCapture={(e) => submitForm(e, email, password)}
            className={classes.loginButton}
          >
            Login
          </Button>
        </form>
    )
  }

  let login = (email, password) => {
    let formData = new FormData();
    formData.append('username', email);
    formData.append('password', password);

    AUTH.post('login', formData)
        .then(response => {
          setToken(response.data.access_token);
          setLoggedIn(true);
          setTimeout(function(){history.push('/home')}, 1000);
        })
        .catch(error => {
          setErrorMessage(error.response.data.detail);
        })
  };

  return (
    <div className={classes.container}>
        <Typography
          className={classes.signUpPrompt}
        >
          Don&apos;t have an account? <Link to="/register"
            className={classes.greenText}
            style={{textDecoration: 'none'}}
            onClick={() => history.push('/register')}
          >Get started.</Link>
        </Typography>
      <Box
        display="flex"
        flexDirection="row"
        className={classes.root}
        justifyContent={showLeft ? 'space-between' : 'center'}
        alignItems="center"
      >
        {showLeft ? <LeftPanel /> : <span />}
        <RightPanel onSubmit={login}/>
      </Box>
    </div>
  )
}

export default Login;