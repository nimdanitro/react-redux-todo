import {
    merge,
    omit
} from "lodash";

import {
    combineReducers
} from 'redux';


import {
    todos as TODOS
} from '../constants/types';


function todosByID(state = {}, action) {
    const {payload} = action;

    switch (action.type) {

        case TODOS.COMPLETE:
            return completeTodo(state, action);

        case TODOS.RECEIVE_DELETE:
            if (payload && payload.id) {
                return omit(state, payload.id)
            }
            return state

        default:
            if (payload && payload.entities && payload.entities.todos) {
                return merge({}, state, payload.entities.todos);
            }
            return state;
    }
}

function allTodos(state = [], action) {
    const {payload} = action;

    switch (action.type) {
        case TODOS.RECEIVE_DELETE:

            if (payload && payload.id) {
                return state.filter(curID => curID !== payload.id);
            }
            return state

        default:
            if (payload && payload.result) {
                return state.concat(payload.result)
            }
            return state;
    }
}

function completeTodo(state, action) {
    const { id } = action;
    const todo = state[id];

    return {
        ...state,
        [id]: {
            ...todo,
            completed: !todo.completed
        }
    };
}


export const todosEntitiesReducer = combineReducers({
    byId: todosByID,
    allIds: allTodos
});
