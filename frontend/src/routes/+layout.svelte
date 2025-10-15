<script lang="ts">
    import favicon from '$lib/assets/favicon.svg';
    import list from '$lib/assets/icons/list.svg';
    import "../app.css";
    import {ERR_NEEDS_RE_LOGIN, ERR_NEEDS_REFRESH_TOKEN, ROUTE_LOGIN} from "$lib/constants";
    import {afterNavigate} from "$app/navigation";
    import {authToken} from "$lib/store/auth";
    import Menu from "./Menu.svelte";

    export const ssr = false;

    let {children} = $props();

    afterNavigate((navigation) => {
        authToken.subscribe(value => {
            if (!value?.token && navigation?.to?.route.id !== ROUTE_LOGIN) {
                window.location.href = ROUTE_LOGIN;
            }
        })

    })

    function handleError(event: ErrorEvent) {
        console.error('Global error caught:', event.error);
        if (event.error === ERR_NEEDS_REFRESH_TOKEN) {
            console.log("AIE COUP DUR")
        } else if (event.error === ERR_NEEDS_RE_LOGIN) {
            // For example, redirect to login page
            window.location.href = ROUTE_LOGIN;

        }
        // Optionally, you can display a user-friendly message or log the error to an external service
    }
</script>

<svelte:window on:error={handleError}/>

<svelte:head>
    <link rel="icon" href={favicon}/>
</svelte:head>


<div class="drawer lg:drawer-open">
    <input id="my-drawer-3" type="checkbox" class="drawer-toggle"/>
    <div class="drawer-content min-h-screen p-4 bg-base-200 flex flex-col items-center">
        {@render children?.()}

    </div>
    <Menu/>
</div>

