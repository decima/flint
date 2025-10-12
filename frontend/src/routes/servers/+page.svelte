<script>
    import Card from "$lib/components/Card.svelte";
    import {onMount} from "svelte";
    import {getServers} from "$lib/api/servers.js";

    let servers = [];
    let loaded = false;
    onMount(async () => {
        servers = await getServers()
        loaded = true;
    })
</script>
<Card title="Servers" width="full" {loaded}>
    {#if servers.length === 0}
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
                    <th>Name</th>
                    <th>Host</th>
                    <th>Port</th>
                    <th>Username</th>
                    <th>Workdir</th>
                    <th></th>
                </tr>
                </thead>
                <tbody>
                {#each servers as server, index}
                    <tr>
                        <td>{server.name}</td>
                        <td>{server.host}</td>
                        <td>{server.port}</td>
                        <td>{server.username}</td>
                        <td>{server.workdir}</td>
                        <td>
                            <a class="btn btn-sm btn-primary" href={`/servers/${server.name}`}>View</a>
                            <a class="btn btn-sm btn-secondary ml-2" href={`/servers/${server.name}/edit`}>Edit</a>
                            <a class="btn btn-sm btn-error ml-2" href={`/servers/${server.name}/delete`}>Delete</a>
                        </td>
                    </tr>
                {/each}
                </tbody>
            </table>
        </div>
    {/if}
</Card>
