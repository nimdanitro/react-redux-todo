import { combineReducers } from 'redux'
import { entities } from './entities'
import { routerReducer } from 'react-router-redux'


const rootReducer = combineReducers({
  entities,
  router: routerReducer
})

export default rootReducer
