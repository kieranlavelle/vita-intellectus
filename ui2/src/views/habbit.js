import { useEffect, useState, useCallback } from 'react'

import { useHistory } from 'react-router-dom'

import { API } from '../http'
import { GetAuthHeaders } from '../auth'
import { Box, Button } from '@material-ui/core'


export default function Habbit(props){

    const history = useHistory()
    const redirect = useCallback((path) => history.push(path), [history]);

    const [habbit, setHabbit] = useState({});
    const config = GetAuthHeaders();

    const deleteHabbit = () => {
        API.delete(`/habbit/${habbit.habbit_id}`, config)
           .then(response => redirect('/habbits'));
    }

    useEffect(() => {
        API.get("/habbits", config).then(response => {
            setHabbit(response.data.filter(
                habbit => habbit.habbit_id == props.id
            )[0]);
        })
    }, [])

    return (
        <Box>
            <h1>{habbit.name}</h1>
            <Button onClick={deleteHabbit}>
                Delete
            </Button>
        </Box>
    )
}