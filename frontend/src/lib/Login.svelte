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

<div class="card w-96 bg-base-100 shadow-xl">
    <div class="card-body">
        <h2 class="card-title">Login</h2>
        <form on:submit|preventDefault={login}>
            <div class="form-control">
                <label class="label" for="username">
                    <span class="label-text">Username</span>
                </label>
                <input id="username" type="text" bind:value={username} required class="input input-bordered"/>
            </div>
            <div class="form-control">
                <label class="label" for="password">
                    <span class="label-text">Password</span>
                </label>
                <input id="password" type="password" bind:value={password} required class="input input-bordered"/>
            </div>
            {#if error}
                <div class="alert alert-error mt-4">
                    <div>
                        <span>{error}</span>
                    </div>
                </div>
            {/if}
            <div class="form-control mt-6">
                <button type="submit" class="btn btn-primary">Login</button>
            </div>
        </form>
    </div>
</div>