<script>
    import { onMount } from 'svelte';

    import NavBar from '../components/NavBar.svelte';
    import { API, getHeaders } from '../http'
    import Habit from '../components/Habit.svelte';

    let habits = [];

    onMount(() => {
        API.get('/habits', getHeaders())
           .then(response => {
               habits = response.data.habits;
        });
    })
</script>

<div class="min-w-screen min-h-screen">
    <NavBar />
    <div class="min-w-full mt-8 grid grid-cols-3 gap-6 px-6 grid-flow-row-dense">
        {#each habits as habit (habit.id)}
            <Habit {...habit} />
        {/each}
    </div>
</div>