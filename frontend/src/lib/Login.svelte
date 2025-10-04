<script lang="ts">
    import {authToken} from './stores';
    import api from "./api";

    let username = '';
    let password = '';
    let error: string | null = null;

    async function login() {
        error = null;
        try {
            const data = await api.login(username, password)
            authToken.set(data.token);
        } catch (e: any) {
            error = e.message;
        }
    }
</script>

<main>
    <h1>Login</h1>
    <form on:submit|preventDefault={login}>
        <div>
            <label for="username">Username</label>
            <input id="username" type="text" bind:value={username} required/>
        </div>
        <div>
            <label for="password">Password</label>
            <input id="password" type="password" bind:value={password} required/>
        </div>
        {#if error}
            <p style="color: red;">{error}</p>
        {/if}
        <button type="submit">Login</button>
    </form>
</main>

<style>
    main {
        max-width: 320px;
        margin: 0 auto;
        padding: 2rem;
    }
</style>