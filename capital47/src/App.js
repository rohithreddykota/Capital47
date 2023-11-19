// App.js
import React from 'react';
import './App.css';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LoginPage from './LoginPage';

import CurrencySendingPage from './CurrencySendingPage'; // Import your CurrencySendingPage component


function App() {
  console.log('I am app')

  return (
    <Router>
      <div className="App">
        <header className="App-header">
          <h1>
            Capital<span>47</span>
          </h1>
        </header>

        <main>
          <Routes>
            <Route path="/currency-sending" component={CurrencySendingPage} />
            <Route path="/login" component={LoginPage} />
            {/* Add more routes as needed */}
          </Routes>
        </main>

        <footer>
          <p>&copy; {new Date().getFullYear()} Capital47</p>
        </footer>
      </div>
    </Router>
  );
}

export default App;
