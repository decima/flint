<script lang="ts">
    import {onMount} from "svelte";
    import api from "../../store/api";

    interface Server {
        name: string;
        host: string;
        port: number;
        username: string;
        workdir: string;
    }


    let servers: Server[] = [];
    let loading = true;
    let error: string | null = null;

    onMount(async () => {
        try {
            const data = await api.listServers();
            servers = data.servers;
        } catch (e: any) {
            error = e.message;
        } finally {
            loading = false;
        }
    });
</script>

<h2 class="text-2xl font-bold mb-4">Servers</h2>

{#if loading}
    <div class="flex justify-center items-center">
        <span class="loading loading-lg"></span>
    </div>
{:else if error}
    <div class="alert alert-error">
        <div>
            <span>{error}</span>
        </div>
    </div>
{:else if servers.length === 0}
    <div class="alert alert-info">
        <div>
            <span>No servers found.</span>
        </div>
    </div>
{:else}
    <div class="overflow-x-auto">
        <table class="table">
            <thead>
            <tr>
                <th></th>
                <th>Name</th>
                <th>Host</th>
                <th>Port</th>
                <th>Username</th>
                <th>Workdir</th>
            </tr>
            </thead>
            <tbody>
            {#each servers as server, index}
                <tr>
                    <th>{index + 1}</th>
                    <td>{server.name}</td>
                    <td>{server.host}</td>
                    <td>{server.port}</td>
                    <td>{server.username}</td>
                    <td>{server.workdir}</td>
                </tr>
            {/each}
            </tbody>
        </table>
    </div>
{/if}
