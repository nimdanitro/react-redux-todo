import React, { Component } from 'react'
import PropTypes from 'prop-types'
import TodoItem from './TodoItem'
import Footer from './Footer'
import { filters } from '../constants/types'

const TODO_FILTERS = {
  [filters.SHOW_ALL]: () => true,
  [filters.SHOW_ACTIVE]: todo => !todo.completed,
  [filters.SHOW_COMPLETED]: todo => todo.completed
}

export default class MainSection extends Component {
  static propTypes = {
    todos: {
      byID: PropTypes.object.isRequired,
      allIds: PropTypes.array.isRequired,
    },
    actions: PropTypes.object.isRequired
  }

  state = { filter: filters.SHOW_ALL }

  handleClearCompleted = () => {
    this.props.actions.clearCompleted()
  }

  handleShow = filter => {
    this.setState({ filter })
  }

  componentDidMount() {
    this.props.actions.fetchAll()
  }


  renderToggleAll(completedCount) {
    const { todos, actions } = this.props
    if (todos.allIds.length > 0) {
      return (
        <input className="toggle-all"
               type="checkbox"
               checked={completedCount === todos.length}
               onChange={actions.completeAll} />
      )
    }
  }

  renderFooter(completedCount) {
    const { todos } = this.props
    const { filter } = this.state
    const activeCount = todos.length - completedCount

    if (todos.allIds.length) {
      return (
        <Footer completedCount={completedCount}
                activeCount={activeCount}
                filter={filter}
                onClearCompleted={this.handleClearCompleted}
                onShow={this.handleShow} />
      )
    }
  }

  render() {
    const { todos, actions } = this.props
    const { filter } = this.state
    const allTodos = todos.allIds.map( id => todos.byId[id])

    const filteredTodos = allTodos.filter(TODO_FILTERS[filter])
    const completedCount = allTodos.reduce((count, todo) =>
      todo.completed ? count + 1 : count,
      0
    )

    return (
      <section className="main">
        {this.renderToggleAll(completedCount)}
        <ul className="todo-list">
          {filteredTodos.map(todo =>
            <TodoItem key={todo.id} todo={todo} deleteTodo={actions.deleteByID} editTodo={actions.edit} completeTodo={actions.complete} />
          )}
        </ul>
        {this.renderFooter(completedCount)}
      </section>
    )
  }
}
