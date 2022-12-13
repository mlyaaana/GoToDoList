import '../../App.css'
import React, { useEffect, useState } from 'react'

import axios from 'axios'
import humps from 'humps'

import url from '../../util'
import CreateTask from './CreateTask'
import TasksList from './TasksList'
import OnLoadingTasks from './TasksLoading'

function Tasks () {
  const DataLoading = OnLoadingTasks(TasksList)

  const [appState, setAppState] = useState(
    {
      loading: false,
      tasks: []
    }
  )

  const createTask = (name, description, deadline) => {
    const apiUrl = url('task')

    axios.post(apiUrl, {
      name,
      description,
      deadline
    })
      .then((resp) => {
        const task = humps.camelizeKeys(resp.data)
        setAppState({
          loading: appState.loading,
          tasks: [task, ...appState.tasks]
        })
      })
      .catch((reason) => console.error(reason))
  }

  const deleteTask = (id) => {
    const apiUrl = url('task')

    axios.delete(apiUrl, {
      params: {
        id
      }
    })
      .then(_ => {
        setAppState({
          loading: appState.loading,
          tasks: appState.tasks.filter(task => task.id !== id)
        })
      })
      .catch((reason) => console.error(reason))
  }

  const doneTask = (id) => {
    const apiUrl = url('task')

    axios.put(apiUrl, null, {
      params: {
        id
      }
    })
      .then((_) => {
        const tasks = appState.tasks
        const index = tasks.findIndex(t => t.id === id)
        const task = tasks[index]
        task.done = !task.done

        setAppState({
          loading: appState.loading,
          tasks
        })
      })
      .catch((reason) => console.error(reason))
  }

  useEffect(() => {
    setAppState({ loading: true, tasks: [] })
    const apiUrl = url('tasks')
    axios.get(apiUrl).then((resp) => {
      const allTasks = humps.camelizeKeys(resp.data)
      setAppState({
        loading: false,
        tasks: allTasks
      })
    }).catch((reason) => console.error(reason))
  }, [setAppState])

  return (
        <div className="App">
            <header className="app-header">
                <p>Todo List</p>
            </header>
            <div>
                <CreateTask onCreate={createTask} />
                <DataLoading isLoading={appState.loading} tasks={appState.tasks} onDelete={deleteTask}
                    onDone={doneTask} />
            </div>
        </div>
  )
}

export default Tasks
