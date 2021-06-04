import React from 'react';
import { BrowserRouter as Router, Switch, Route, Redirect } from "react-router-dom";
import RegisterPage from './pages/Register';
import SignIn from './pages/Login';
import Home from './pages/Home';
import AskQuestion from './pages/AskQuestion';
import Footer from '../src/components/layout/Footer';
import Dashboard from './components/dashboard/Dashboard';
import mainListItems from './components/dashboard/listItems';
import Title from './components/dashboard/Title';
import QuestionDetail from './pages/QuestionDetail';
import Search from './components/Search';
import NotFound from './pages/Notfound';

function App() {
  let search = window.location.search;
  let params = new URLSearchParams(search);
  let foo = params.get("search");
  let url = "/questions/?search" + foo
  return (
    <div>
      <div className="App">
        <Router>
          <Switch>
            <Route path="/signup" exact component={RegisterPage} />
            <Route path="/login" exact component={SignIn} />
            <Route path="/admin" exact component={Home} />
            <Route path="/questions/ask" exact component={AskQuestion} />
            <Route path="/questions/:id" exact component={QuestionDetail} />
            {/* <Route path="/questions/result/:query" component={Search} /> */}
            <Route path="/" exact component={Home} />
            <Route component={NotFound} />
          </Switch>
        </Router>
        {/* <header>
        </header>
        <body>
          <Dashboard />
          <mainListItems />
          <Title />
        </body>
        <Footer /> */}
      </div>
    </div>
  );
}

export default App;