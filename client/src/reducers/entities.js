import { combineReducers }      from 'redux';


import { todosEntitiesReducer } from './todos';

export const entities = combineReducers({
    todos: todosEntitiesReducer,
});
