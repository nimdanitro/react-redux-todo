import fetch from 'isomorphic-fetch'
import {
    todos as TODOS
} from '../constants/types';
import {
    createAction
} from './generator';
import {
    schema,
    normalize
} from 'normalizr';


const todos = new schema.Entity('todos');

/* Static */
//export const add            = createAction(, 'text');
export const edit = createAction(TODOS.EDIT, 'id', 'text');
export const complete = createAction(TODOS.COMPLETE, 'id');
export const completeAll = createAction(TODOS.COMPLETE_ALL);
export const clearCompleted = createAction(TODOS.CLEAR_COMPLETED);

// ADD
export const requestAdd = createAction(TODOS.REQUEST_ADD, 'text');
export const failureAdd = createAction(TODOS.FAILURE_ADD, 'error');
export function add(text) {

    return (dispatch) => {
        dispatch(requestAdd(text))

        return fetch(`/api/v1/todos`, {
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                method: "POST",
                body: JSON.stringify({
                    text: text
                })
            })
            .then(response => response.json())
            .then(json => dispatch(receiveAdd(json)))
            .catch(err => dispatch(failureAdd(err)))
    }
}
export function receiveAdd(json) {
    return {
        type: TODOS.RECEIVE_ADD,
        payload: normalize(json, todos),
        receivedAt: Date.now()
    }
}


// DELETE
export const requestDelete = createAction(TODOS.REQUEST_DELETE, 'id');
export const receiveDelete = createAction(TODOS.RECEIVE_DELETE, 'payload');
export const failureDelete = createAction(TODOS.FAILURE_DELETE, 'error');

export function deleteByID(id) {

    return (dispatch) => {
        dispatch(requestDelete(id))

        return fetch(`/api/v1/todo/${id}`, {
                headers: {
                    'Accept': 'application/json',
                },
                method: "DELETE",
            })
            .then(response => response.json())
            .then(json => dispatch(receiveDelete(json)))
            .catch(err => dispatch(failureDelete(err)))

    }
}

export function requestTodos() {
    return {
        type: TODOS.REQUEST_ALL,
        requstedAt: Date.now()
    }
}

export function receiveTodos(json) {
    return {
        type: TODOS.RECEIVE_ALL,
        payload: normalize(json, [todos]),
        receivedAt: Date.now()
    }
}

export function fetchAll() {
    return (dispatch) => {

        dispatch(requestTodos())

        return fetch(`/api/v1/todos`)
            .then(response => response.json())
            .then(json => dispatch(receiveTodos(json)))

    }

}
