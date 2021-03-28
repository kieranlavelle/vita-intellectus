import React, { useEffect } from 'react'


export default function usePersistedState(key, defaultValue) {
    const [state, setState] = React.useState(
        localStorage.getItem(key) || defaultValue
    );
    useEffect(() => {
        localStorage.setItem(key, state);
    }, [state, setState]);
    return [state, setState]
}