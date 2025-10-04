<script lang="ts">
    import {onMount} from 'svelte';
    import api from './api';

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

<main>
    <h2>Servers</h2>
    {#if loading}
        <p>Loading servers...</p>
    {:else if error}
        <p style="color: red;">{error}</p>
    {:else if servers.length === 0}
        <p>No servers found.</p>
    {:else}
        <table>
            <thead>
            <tr>
                <th>Name</th>
                <th>Host</th>
                <th>Port</th>
                <th>Username</th>
                <th>Workdir</th>
            </tr>
            </thead>
            <tbody>
            {#each servers as server}
                <tr>
                    <td>{server.name}</td>
                    <td>{server.host}</td>
                    <td>{server.port}</td>
                    <td>{server.username}</td>
                    <td>{server.workdir}</td>
                </tr>
            {/each}
            </tbody>
        </table>
    {/if}
</main>

<style>
    table {
        width: 100%;
        border-collapse: collapse;
    }

    th, td {
        border: 1px solid #ddd;
        padding: 8px;
        text-align: left;
    }

    th {
        background-color: #f2f2f2;
    }
</style>