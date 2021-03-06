import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles';

import Router from './components/Router'

const theme = createMuiTheme({
  typography: {
    fontFamily: ['"Be Vietnam"', 'sans-serif'].join(',')
  },
  palette: {
    primary: {
      main: 'rgb(0, 171, 85)'
    },
    secondary: {
      // main: '#ff1744',
      main: '#f44336'
    },
    error: {
      main: '#f44336'
    }
  },
})

ReactDOM.render(
  <ThemeProvider theme={theme}>
    <Router>
    </Router>
  </ThemeProvider>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
