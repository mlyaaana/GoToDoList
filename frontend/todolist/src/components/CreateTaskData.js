import React, { useState } from 'react'

import { Button } from '@material-ui/core'
import TextField from '@material-ui/core/TextField'

function CreateTaskData (props) {
  const { onCreate } = props

  const [name, setName] = useState('')
  const [description, setDescription] = useState('')
  const [deadline, setDeadline] = useState(new Date())

  const createTask = (event) => {
    event.preventDefault()

    onCreate(name, description, deadline)
    setName('')
    setDescription('')
    setDeadline(new Date())
  }

  return (
        <form onSubmit={createTask}>
            <TextField
                name="name"
                placeholder="Name"
                value={name}
                required
                multiline
                classes={{ root: 'input' }}
                onChange={(e) => setName(e.target.value)}
            />
            <TextField
                name="description"
                placeholder="Description"
                value={description}
                multiline
                classes={{ root: 'input' }}
                onChange={(e) => setDescription(e.target.value)}
            />
            <TextField
                type="datetime-local"
                classes={{ root: 'input' }}
                onChange={(e) => setDeadline(new Date(e.target.value))}
            />
            <br />
            <Button variant="outlined" type="submit">Create</Button>
        </form>
  )
}

export default CreateTaskData
