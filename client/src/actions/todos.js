import fetch                  from 'isomorphic-fetch'
import { todos as TODOS }     from '../constants/types';
import { createAction }       from './generator';
import { schema, normalize }  from 'normalizr';
import { v4 as uuid }         from 'uuid';



const todos = new schema.Entity('todos');

/* Static */
//export const add            = createAction(, 'text');
export const deleteByID     = createAction(TODOS.DELETE, 'id');
export const edit           = createAction(TODOS.EDIT, 'id', 'text');
export const complete       = createAction(TODOS.COMPLETE, 'id');
export const completeAll    = createAction(TODOS.COMPLETE_ALL);
export const clearCompleted = createAction(TODOS.CLEAR_COMPLETED);


export function add(text) {
  let id = uuid()
  let normalized = normalize({
    id: id,
    text: text,
  }, todos)
  console.log(normalized)
  return {
    type: TODOS.ADD,
    payload: normalized
  }

}

// /* API */
// export function fetch(id) {
//   return {
//     [RSAA]: {
//       endpoint: `/api/v1/todos/${id}`,
//       method: 'GET',
//       types: [TODOS.REQUEST, TODOS.RECEIVE, TODOS.FAILURE]
//     }
//   }
// }
//
// export function publish(id) {
//   return {
//     [RSAA]: {
//       endpoint: '/api/v1/todos',
//       method: 'POST',
//       types: [TODOS.REQUEST, TODOS.RECEIVE, TODOS.FAILURE]
//     }
//   }
// }
//
// export function fetchAll() {
//   console.log("FetchAll called")
//   return {
//     [RSAA]: {
//       endpoint: '/api/v1/todos',
//       method: 'GET',
//       types: [
//         TODOS.REQUEST_ALL, {
//           type: TODOS.RECEIVE_ALL,
//           payload: (action, state, res) => getJSON(res).then((json) => normalize(json, [todos]))
//         }, {
//           type: TODOS.FAILURE_ALL,
//           payload: (action, state, res) => getJSON(res).then((json) => normalize(json, [todos]))
//         }
//       ]
//     }
//   }
// }

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
  console.log("fetchAll Dispatch")

  return (dispatch) => {

    dispatch(requestTodos())

    return fetch(`/api/v1/todos`)
      .then(response => response.json())
      .then(json => dispatch(receiveTodos(json)))
  }

}
