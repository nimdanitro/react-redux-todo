import React, { Component } from 'react';
import PropTypes from 'prop-types';
import classnames from 'classnames';
import { filters } from '../constants/types';

const FILTER_TITLES = {
  [filters.SHOW_ALL]: 'All',
  [filters.SHOW_ACTIVE]: 'Active',
  [filters.SHOW_COMPLETED]: 'Completed'
}

export default class Footer extends Component {
  static propTypes = {
    completedCount: PropTypes.number.isRequired,
    activeCount: PropTypes.number.isRequired,
    filter: PropTypes.string.isRequired,
    onClearCompleted: PropTypes.func.isRequired,
    onShow: PropTypes.func.isRequired
  }

  renderTodoCount() {
    const { activeCount } = this.props
    const itemWord = activeCount === 1 ? 'item' : 'items'

    return (
      <span className="todo-count">
        <strong>{activeCount || 'No'}</strong> {itemWord} left
      </span>
    )
  }

  renderFilterLink(filter) {
    const title = FILTER_TITLES[filter]
    const { filter: selectedFilter, onShow } = this.props

    return (
      <a className={classnames({ selected: filter === selectedFilter })}
         style={{ cursor: 'pointer' }}
         onClick={() => onShow(filter)}>
        {title}
      </a>
    )
  }

  renderClearButton() {
    const { completedCount, onClearCompleted } = this.props
    if (completedCount > 0) {
      return (
        <button className="clear-completed"
                onClick={onClearCompleted} >
          Clear completed
        </button>
      )
    }
  }

  render() {
    return (
      <footer className="footer">
        {this.renderTodoCount()}
        <ul className="filters">
          {[ filters.SHOW_ALL, filters.SHOW_ACTIVE, filters.SHOW_COMPLETED ].map(filter =>
            <li key={filter}>
              {this.renderFilterLink(filter)}
            </li>
          )}
        </ul>
        {this.renderClearButton()}
      </footer>
    )
  }
}
