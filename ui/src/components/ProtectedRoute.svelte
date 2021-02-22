<script>
    import { navigate, Route } from 'svelte-routing';
    import { token } from '../store'

    export let path
    export let redirect
    export let component

    let pathname = window.location.pathname;
    let paths = ['/login', '/register']
    
    $: if (!$token) {
        if (!paths.includes(pathname)){
            navigate(redirect, {replace: true});
        }
    }

</script>

<Route {path}>
    <svelte:component this={component} />
</Route>