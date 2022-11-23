import List from '@material-ui/core/List';
import TaskData from './TaskData';

function TasksData(props) {
    const { tasks, onDelete, onDone } = props

    if (!tasks || tasks.length === 0) return <p className="no-data">Нет данных.</p>

    return (
        <List>
            {
                tasks.map((task, index) => <TaskData key={task.id} task={task} index={index} onDelete={onDelete} onDone={onDone} />)
            }
        </List>
    )
}

export default TasksData
