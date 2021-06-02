import React from 'react';
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
  );
}

export default App;
