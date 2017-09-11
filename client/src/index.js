import 'babel-polyfill'
import React                            from 'react'
import { render }                       from 'react-dom'
import { Provider }                     from 'react-redux'
import { createStore, applyMiddleware } from 'redux'
import { composeWithDevTools }          from 'redux-devtools-extension';
import thunk                            from 'redux-thunk';


import App     from './containers/App'
import { default as reducer } from './reducers'
import 'todomvc-app-css/index.css'


const composeEnhancers = composeWithDevTools({});
const store = createStore(reducer,
  composeEnhancers(
    applyMiddleware(thunk)
  )
);

render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
)
