import List from '@material-ui/core/List';
import TaskData from './TaskData';

function TasksData(props) {
    const { tasks } = props

    if (!tasks || tasks.length === 0) return <p className="no-data">Нет данных.</p>

    return (
        <List>
            {
                tasks.map((task, index) => <TaskData key={task.Id} task={task} index={index} />)
            }
        </List>
    )
}

export default TasksData
