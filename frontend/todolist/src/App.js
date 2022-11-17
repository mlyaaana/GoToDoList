import './App.css';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import CreateTaskData from './components/CreateTaskData'
import TasksData from './components/TasksData'
import OnLoadingTasksData from './components/LoadingTasksData'
import url from './util'

function App() {
  const DataLoading = OnLoadingTasksData(TasksData);

  const [appState, setAppState] = useState(
    {
      loading: false,
      tasks: [],
    }
  )

  const createTask = (name, description) => {
    const apiUrl = url('task')

    axios.post(apiUrl, null, {
      params: {
        name,
        description
      }
    })
      .then((resp) => {
        const task = resp.data
        setAppState({
          loading: appState.loading,
          tasks: [task, ...appState.tasks]
        })
      })
      .catch((reason) => console.error(reason))
  }

  const deleteTask = (id) => {
    const apiUrl = url('task');

    axios.delete(apiUrl, {
      params: {
        id
      }
    })
      .then(_ => {
        setAppState({
          loading: appState.loading,
          tasks: appState.tasks.filter(task => task.Id !== id)
        })
      })
      .catch((reason) => console.error(reason))
  }

  useEffect(() => {
    setAppState({ loading: true })
    const apiUrl = url('tasks');
    axios.get(apiUrl).then((resp) => {
      const allTasks = JSON.parse(resp.data);
      setAppState({
        loading: false,
        tasks: allTasks
      });
    }).catch((reason) => console.error(reason));
  }, [setAppState]);

  return (
    <div className="App">
      <header className="app-header">
        <p>Todo List</p>
      </header>
      <div>
        <CreateTaskData onCreate={createTask} />
        <DataLoading isLoading={appState.loading} tasks={appState.tasks} onDelete={deleteTask} />
      </div>
    </div>
  );
}

export default App;
