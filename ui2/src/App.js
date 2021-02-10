import './App.css';

import Router from './routing/router' 
import { Link } from 'react-router-dom'


export default function App() {
  return (
      <div className="App">
          <Router>
              <Link to="/login">Login</Link>
          </Router>
      </div>
  );
}