import { Checkbox, ListItem, ListItemSecondaryAction, ListItemText } from '@material-ui/core';

function TaskData(props) {
    const { task, index } = props

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
            />
            <ListItemSecondaryAction>
                <Checkbox />
            </ListItemSecondaryAction>
        </ListItem>
    )
}

export default TaskData;
