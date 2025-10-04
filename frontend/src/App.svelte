<script lang="ts">
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

<main>
  <header>
    <h1>Flint</h1>
    {#if token}
      <button on:click={logout}>Logout</button>
    {/if}
  </header>

  {#if token}
    <ServerList />
  {:else}
    <Login />
  {/if}
</main>

<style>
  main {
    padding: 1em;
    max-width: 800px;
    margin: 0 auto;
  }
  header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
  }
  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 2em;
    font-weight: 100;
  }
</style>