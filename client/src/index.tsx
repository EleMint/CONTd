import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';

import './index.css';

import { App } from './App';
import { Account } from './contexts/Account';
import { API } from './contexts/API';

ReactDOM.render(
  <React.StrictMode>
    <BrowserRouter>
      <Account>
        <API>
          <App />
        </API>
      </Account>
    </BrowserRouter>
  </React.StrictMode>,
  document.getElementById('root')
);
