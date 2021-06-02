import React from 'react';
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import RegisterPage from './pages/Register';
import SignIn from './pages/Login';
import Home from './pages/Home';
import Header from "./components/Header";
import Sidebar from './components/Sidebar';

function App() {
  return (
    <>
      <div className="App">
        <body>
          <Router>
            <Switch>
              <Route path="/signup" exact component={RegisterPage} />
              <Route path="/login" exact component={SignIn} />
              {/* <Route path="/admin" exact component={HomePage} /> */}
              <Route path="/" exact component={Home} />
            </Switch>
          </Router>
        </body>
      </div>
    </>
  );
}

export default App;
