import React from 'react';
<<<<<<< HEAD
import './App.css';
import Footer from '../src/components/layout/Footer';
import Dashboard from './components/dashboard/Dashboard';
import mainListItems from './components/dashboard/listItems';
import Title from './components/dashboard/Title';


function App() {
  return (
    <div className="App">
      <header>
      </header>
      <body>
        <Dashboard />
        <mainListItems />
        <Title />
      </body>
      <Footer />
    </div>
=======
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import RegisterPage from './pages/Register';
import SignIn from './pages/Login';
import Home from './pages/Home';
import AskQuestion from './pages/AskQuestion';

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
              <Route path="/questions/ask" exact component={AskQuestion} />
              <Route path="/" exact component={Home} />
            </Switch>
          </Router>
        </body>
      </div>
    </>
>>>>>>> fe12eac0d29901b9e7bfc8555c9490be1bf810fb
  );
}

export default App;
