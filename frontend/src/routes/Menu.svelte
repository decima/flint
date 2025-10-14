<script lang="ts">
    import {isAuthenticated} from "$lib/store/auth.js";
    import logo from '$lib/assets/logo.svg';
    import {page} from "$app/state";

    function activeClass(path: string) {
        return page.url.pathname === path ? 'menu-active dock-active' : '';
    }
</script>

<div class="drawer-side shadow">
    <label for="my-drawer-3" aria-label="close sidebar" class="drawer-overlay"></label>

    <div class="min-h-full bg-base-100 p-4">
        <img src={logo} class="w-64 "/>
        {#if isAuthenticated()}
            <ul class="menu bg-base-100 w-full p-4">
                <!-- Sidebar content here -->
                <li class="{activeClass('/')}"><a href="/">Dashboard</a></li>
                <li class="{activeClass('/servers')}"><a href="/servers">Servers</a></li>
                <li><a href="/logout">Logout</a></li>
            </ul>
        {/if}
    </div>
</div>

{#if isAuthenticated()}
    <div class="dock lg:hidden">
        <a href="/" class="{activeClass('/')}">
            <span class="dock-label">Dashboard</span>
        </a>

        <a href="/servers" class="{activeClass('/servers')}">
            <span class="dock-label">Servers</span>
        </a>

        <a href="/logout">
            <span class="dock-label">Logout</span>
        </a>
    </div>
{/if}
