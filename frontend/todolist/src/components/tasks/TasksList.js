import React from 'react'

import List from '@material-ui/core/List'

import TaskItem from './TaskItem'

function TasksList (props) {
  const { tasks, onDelete, onDone } = props

  if (!tasks || tasks.length === 0) return <p className="no-data">Нет данных.</p>

  return (
        <List>
            {
                tasks.sort((a, b) => new Date(a.createdAt) - new Date(b.createdAt)).map((task, index) => <TaskItem
                    key={task.id}
                    task={task}
                    index={index}
                    onDelete={onDelete}
                    onDone={onDone}/>)
            }
        </List>
  )
}

export default TasksList
