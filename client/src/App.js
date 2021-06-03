import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import RegisterPage from "./pages/Register";
import SignIn from "./pages/Login";
import Home from "./pages/Home";
import AskQuestion from "./pages/AskQuestion";
import Footer from "../src/components/layout/Footer";
import Dashboard from "./components/dashboard/Dashboard";
import mainListItems from "./components/dashboard/listItems";
import Title from "./components/dashboard/Title";
import QuestionDetail from "./pages/QuestionDetail";

function App() {
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
            <Route path="/" exact component={Home} />
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
