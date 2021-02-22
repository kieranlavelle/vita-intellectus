<script>
	import { Router, Route } from 'svelte-routing';

	import Login from './views/Login.svelte';
	import Register from './views/Register.svelte';
	import Habits from './views/Habits.svelte';
	import ProtectedRoute from './components/ProtectedRoute.svelte'

	export let url;
</script>

<main>

	<Router url="{url}">
		<Route path="/login" component={Login} />
		<Route path="/register" component={Register} />
		<!-- <Route path="/habits" component={Habits} /> -->
		<ProtectedRoute path="/habits" componenet={Habits} redirect="/login" />
	</Router>

</main> 

<style global>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>