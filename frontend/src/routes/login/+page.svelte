<script lang="ts">
    import {enhance} from '$app/forms';

    /** @type {import('./$types').ActionData} */
    export let form;

</script>

<form method="POST" action="?/login" use:enhance
      class="w-96 bg-base-100 p-6 rounded-lg shadow-lg">

    <div class="form-control mb-4">
        <label class="label" for="username">
            <span class="label-text">Username</span>
        </label>
        <input id="username" type="text" name="username" required class="input input-bordered"/>
    </div>
    <div class="form-control mb-4">
        <label class="label" for="password">
            <span class="label-text">Password</span>
        </label>
        <input id="password" type="password" name="password" required class=" input input-bordered"/>
    </div>

    <button type="submit" class="btn btn-primary">Login</button>

</form>