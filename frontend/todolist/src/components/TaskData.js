import { Checkbox, IconButton, ListItem, ListItemSecondaryAction, ListItemText } from '@material-ui/core';
import DeleteIcon from '@material-ui/icons/Delete';
import axios from 'axios';
import url from '../util'

function TaskData(props) {
    const { task, index, onDelete, onDone } = props

    return (
        <ListItem divider={true}>
            <ListItemText
                primary={index + 1}
                classes={{
                    root: 'item-id'
                }}
            />
            <ListItemText
                primary={task.name}
                secondary={task.description}
                classes={{
                    root: 'item-content'
                }}
            />
            <ListItemSecondaryAction>
                <IconButton aria-label="delete" onClick={() => onDelete(task.id)}>
                    <DeleteIcon />
                </IconButton>
                <Checkbox checked={task.done} onClick={() => onDone(task.id)} />
            </ListItemSecondaryAction>
        </ListItem>
    )
}

export default TaskData;
