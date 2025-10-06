<script>
    import {onMount} from "svelte";
    import {authToken} from "../../store/stores.js";

    onMount(async () => {
        // Clear the authentication token from local storage
        authToken.set(false);
        // Redirect to home page after logout
        window.location.href = "/login";
    })
    ;
</script>