import 'babel-polyfill'
import React                            from 'react'
import { render }                       from 'react-dom'
import { Provider }                     from 'react-redux'
import { createStore, combineReducers, applyMiddleware } from 'redux'
import { composeWithDevTools }          from 'redux-devtools-extension';
import thunk                            from 'redux-thunk';

import createHistory from 'history/createBrowserHistory'
import { Route } from 'react-router'

import { ConnectedRouter, routerMiddleware } from 'react-router-redux'


import App     from './containers/App'
import { default as reducers } from './reducers'
import 'todomvc-app-css/index.css'


// Create a history of your choosing (we're using a browser history in this case)
const history = createHistory()

// Build the middleware for intercepting and dispatching navigation actions
const router = routerMiddleware(history)


const composeEnhancers = composeWithDevTools({});
const store = createStore(
    reducers,
    composeEnhancers(
        applyMiddleware(thunk,router)
    )
);

render(
  <Provider store={store}>
    <ConnectedRouter history={history}>
        <div>
            <Route exact path="/" component={App}/>
            <Route path="/about" component={App}/>
            <Route path="/topics" component={App}/>
        </div>
    </ConnectedRouter>
  </Provider>,
  document.getElementById('root')
)
