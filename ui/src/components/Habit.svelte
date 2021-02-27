<script>
// import { getHeaders } from "../http";

    import { API, getHeaders } from '../http'

    export let id;
    export let name;
    export let days;
    export let completed;
    export let user_id;
    export let tags;


    $: info = {
        streak: 0
    };
    API.get(`/habit/${id}/info`, getHeaders())
       .then(response => {
           console.log(response);
           info = response.data.info;
       })


    // $: nextDue = completed ? due_dates.next_due : due_dates.next_due_on_completed;
    // $: state = 'Not Due';

    // if (completed){
    //     state = 'Completed';
    // } else if (nextDue == 'Today') {
    //     state = 'Due';
    // }
</script>

<div class="rounded shadow-xl hover:shadow-2xl border border-gray-300">
    <div class="grid grid-cols-1 p-4">
        <span class="font-bold py-1">{name}</span>
        <span class="text-sm py-1">
            Streak: <span class="text-green-800 font-bold">
                {info.streak}
            </span>
        </span>
        {#if tags.length > 0}
            <span class="pb-1">Tags:
                {#each tags as tag}
                    <span class="bg-indigo-400 px-2 rounded-3xl text-white">
                        {tag}
                    </span>
                {/each}
            </span>
        {/if}
        <span class="text-right">
            <button class="border-2 border-green-600 border-opacity-50 rounded-md p-1 hover:border-3 hover:border-green-500 hover:border-opacity-100">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" class="h-4 w-4 text-green-600">
                    <path d="M4 12l6 6L20 6" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
            </button>

            <button class="border-2 border-indigo-200 rounded-md p-1 hover:border-3 hover:border-indigo-400">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" class="h-4 w-4 text-indigo-500">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                </svg>
            </button>

            <button class="border-2 border-red-200 rounded-md p-1 hover:border-3 hover:border-red-400">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" class="h-4 w-4 text-red-500">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                </svg>
            </button>
        </span>
    </div>
</div>