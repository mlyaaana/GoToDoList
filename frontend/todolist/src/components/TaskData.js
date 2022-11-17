import { Checkbox, IconButton, ListItem, ListItemSecondaryAction, ListItemText } from '@material-ui/core';
import DeleteIcon from '@material-ui/icons/Delete';

function TaskData(props) {
    const { task, index, onDelete } = props

    return (
        <ListItem divider={true}>
            <ListItemText
                primary={index + 1}
                classes={{
                    root: 'item-id'
                }}
            />
            <ListItemText
                primary={task.Name}
                secondary={task.Description}
                classes={{
                    root: 'item-content'
                }}
            />
            <ListItemSecondaryAction>
                <IconButton aria-label="delete" onClick={() => onDelete(task.Id)}>
                    <DeleteIcon />
                </IconButton>
                <Checkbox />
            </ListItemSecondaryAction>
        </ListItem>
    )
}

export default TaskData;
