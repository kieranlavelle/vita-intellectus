import './App.css';

import Router from './routing/router' 
import { Link } from 'react-router-dom'

import Navigation from './components/nav'


function App() {
  return (
    <div className="App">
      <Navigation loggedIn={false} />
      <Router>
          <Link to="/login">Login</Link>
      </Router>
    </div>
  );
}

export default App;
