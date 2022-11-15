import './App.css';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import CreateTaskData from './components/CreateTaskData'
import TasksData from './components/TasksData'
import OnLoadingTasksData from './components/LoadingTasksData'

function App() {
  const DataLoading = OnLoadingTasksData(TasksData);

  const [appState, setAppState] = useState(
    {
      loading: false,
      tasks: [],
    }
  )

  const submit = (event) => {
    const apiUrl = '/task'

    const name = event.target.name.value
    const description = event.target.description.value

    axios.post(apiUrl, null, { params: {
      name,
      description
    }}).catch((reason) => console.error(reason))
  }

  useEffect(() => {
    setAppState({ loading: true })
    const apiUrl = '/tasks';
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
        <CreateTaskData onClick={submit} />
        <DataLoading isLoading={appState.loading} tasks={appState.tasks} />
      </div>
    </div>
  );
}

export default App;
