<script>
    import { navigate } from 'svelte-routing';
    import { token } from '../store';

    let src = 'images/logos.png'

    function logout(){
        token.reset();
    }

    $: if (!$token){
        navigate('/login', {replace: true});
    }

</script>

<nav 
    class="grid grid-cols-8 py-2 px-4 text-lg border-b border-gray-400 font-semibold"
>
    <img {src} alt="the vita logo" width="75" height="75" class="col-span-1"/>
    <button class="text-left col-span-1 border-none font-semibold">
        Habbits
    </button>
    <button
        class="text-right col-span-6 border-none font-semibold"
        on:click={logout}
    >
        Logout
    </button>
</nav>