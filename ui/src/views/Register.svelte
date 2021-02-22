<script>
  import { navigate } from "svelte-routing";
  import { AUTH } from '../http'

  let email = '';
  let password = '';
  let errorMessage = '';

  function register(){
      let formData = new FormData();
      formData.append('username', email);
      formData.append('email', email);
      formData.append('password', password);

      AUTH.post('/register', formData)
          .then(response => {
              navigate('/login', {replace: true});
          })
          .catch(error => {
              errorMessage = error.response.data.detail;
          })
  }

</script>

<main>
  <div class="flex flex-col justify-center min-h-screen min-w-full items-center bg-gray-50">
      <div class="max-w-md w-full space-y-8"></div>
          <div class="text-center">
              <h1 class="text-3xl font-extrabold text-gray-900">Register an account</h1>
              <p class="mt-2">
                  Or <a class="text-md text-indigo-600 hover:text-indigo-700">log in to your account</a>
              </p>
          </div>
          {#if errorMessage != ''}
              <div class="p-2 m-5 border border-red-400 rounded-md text-center">
                  <p>{errorMessage}</p>
              </div>
          {/if}
          <form on:submit|preventDefault={register} class="flex flex-col space-y-6 mt-8">
              <div class="-space-y-px">
                  <input bind:value={email} type="text" placeholder="Email address" class="border border-gray-300 placeholder-gray-500 w-full rounded-t-md"/>
                  <input bind:value={password} type="password" placeholder="Password" class="border border-gray-300 placeholder-gray-500 w-full rounded-b-md"/>
              </div>
              <div class="flex items-center justify-between">
                  <div class="flex items-center">
                      <input id="remember_me" name="remember_me" type="checkbox" class="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded">
                      <label for="remember_me" class="ml-2 block text-sm text-gray-900">
                          Remember me
                      </label>
                  </div>
                  <div class="text-sm">
                      <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500">
                        Forgot your password?
                      </a>
                  </div>
              </div>
              <div>
                  <button 
                      class="w-full bg-indigo-600 hover:bg-indigo-700 text-white rounded-md py-2 text-sm font-medium"
                      type="submit"
                  >
                      Register
                  </button>
              </div>
          </form>
      </div>
</main>