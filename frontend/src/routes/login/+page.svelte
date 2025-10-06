<script lang="ts">

    import {authToken, login} from "$lib/store/auth.js";
    import Card from "$lib/components/Card.svelte";

    authToken.subscribe((value) => {
        console.log("TODO: redirect if already logged in")
    });

    let username = "";
    let password = "";
    let error: string | null = null;

    async function loginSubmit(evt: SubmitEvent) {
        evt.preventDefault()
        try {
            await login(username, password)
            window.location.href = "/";
        } catch (e: any) {
            error = e
            password = ""
        }
    }
</script>

<Card title="Login">

    {#if error}
        <div class="alert alert-error alert-soft">
            <div>
                <span>{error}</span>
            </div>
        </div>
    {/if}
    <form on:submit={loginSubmit}>
        <div class="form-control w-full max-w-xs">
            <label class="label">
                <span class="label-text">Username</span>
            </label>
            <input type="text" bind:value={username} name="username" placeholder=""
                   class="input input-bordered w-full max-w-xs"
                   required/>
        </div>

        <div class="form-control w-full max-w-xs mt-4">
            <label class="label">
                <span class="label-text">Password</span>
            </label>
            <input type="password" bind:value={password} name="password" placeholder=""
                   class="input input-bordered w-full max-w-xs" required/>
        </div>
        <div class="form-control mt-6">
            <button type="submit" class="btn btn-primary w-full max-w-xs">Login</button>
        </div>
</Card>