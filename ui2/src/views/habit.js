import { useEffect, useState, useCallback } from 'react'

import { useHistory } from 'react-router-dom'

import { API } from '../http'
import { GetAuthHeaders } from '../auth'
import { Box, Button } from '@material-ui/core'


export default function Habit(props){

    const history = useHistory()
    const redirect = useCallback((path) => history.push(path), [history]);

    const [habit, setHabit] = useState({});
    const config = GetAuthHeaders();

    const deleteHabit = () => {
        API.delete(`/habit/${habit.habit_id}`, config)
           .then(response => redirect('/habits'));
    }

    useEffect(() => {
        API.get("/habits", config).then(response => {

            const allHabits = [
                ...response.data.due,
                ...response.data.completed,
                ...response.data.not_due
            ]

            setHabit(allHabits.filter(
                habit => habit.habit_id == props.id
            )[0]);
        })
    }, [])

    return (
        <Box>
            <h1>{habit.name}</h1>
            <Button onClick={deleteHabit}>
                Delete
            </Button>
        </Box>
    )
}