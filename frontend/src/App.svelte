<script lang="ts">
  import "./app.css";
  import { authToken } from './lib/stores';
  import Login from './lib/Login.svelte';
  import ServerList from './lib/ServerList.svelte';

  let token: string | null;

  authToken.subscribe(value => {
    token = value;
  });

  function logout() {
    authToken.set(null);
  }
</script>

<div class="navbar bg-base-100 shadow-lg">
  <div class="flex-1">
    <a href="/" class="btn btn-ghost normal-case text-xl">Flint</a>
  </div>
  <div class="flex-none">
    {#if token}
      <button class="btn btn-primary" on:click={logout}>Logout</button>
    {/if}
  </div>
</div>

<main class="container mx-auto p-4">
  {#if token}
    <ServerList />
  {:else}
    <div class="flex justify-center items-center h-[80vh]">
        <Login />
    </div>
  {/if}
</main>