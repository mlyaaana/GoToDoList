import { Button } from '@material-ui/core';
import TextField from '@material-ui/core/TextField';

function CreateTaskData(props) {
    const { onClick } = props

    return (
        <form onSubmit={onClick}>
            <TextField name="name" placeholder="Name" required multiline classes={{root: 'input'}} />
            <TextField name="description" placeholder="Description" multiline classes={{root: 'input'}} />
            <br/>
            <Button variant="outlined" type="submit">Create</Button>
        </form>
    )
}

export default CreateTaskData
