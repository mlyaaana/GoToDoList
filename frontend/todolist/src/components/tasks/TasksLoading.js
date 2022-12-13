import React from 'react'

function OnLoadingTasks (Component) {
  return function LoadingTasksData ({ isLoading, ...props }) {
    if (!isLoading) return <Component {...props} />

    else {
      return (
            <div>
                <p className="loading-data">Подождите, данные загружаются!</p>
            </div>
      )
    }
  }
}

export default OnLoadingTasks
