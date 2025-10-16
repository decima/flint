<script lang="ts">
    import type {PageProps} from './$types';
    import Card from "$lib/components/Card.svelte";
    import {onMount} from "svelte";
    import {getServerSummary} from "$lib/api/servers";
    import Badge from "$lib/components/Badge.svelte";
    import Terminal from "$lib/components/Terminal.svelte";
    import Fs from "$lib/components/fs/Fs.svelte";

    let {data}: PageProps = $props();
    let showFs = $state(true);

    let showTerminal = $state(false);
    let details = $state({
        server: {
            host: '',
            port: 0,
            username: '',
            work_dir: '',
        },
        docker: {
            containers: {total: 0, running: 0, stopped: 0, paused: 0},
            images: 0,
            client: {version: '', api_version: '', architecture: '', operating_system: ''},
            server: {operating_system: '', architecture: '', server_version: '', kernel_version: ''},
        },
    });
    let loaded = $state(false);

    onMount(async () => {
        details = await getServerSummary(data.server)
        loaded = true
    })
</script>
<div class="grid gap-4 grid-cols-1 md:grid-cols-1 lg:grid-cols-3 w-full">
    <Card classes={["col-span-1","md:col-span-2","lg:col-span-1"]} bind:loaded={loaded}
          title="Server details">
        <h1>{details.server?.host}:{details.server?.port}</h1>
        <p>User: {details.server?.username}</p>
        <p>Working directory: {details.server?.work_dir}</p>
    </Card>

    <Card bind:loaded={loaded}
          title="Server"
          width="full">
        <div>
            OS: {details.docker.server.operating_system} ({details.docker.server.architecture})<br>
            Docker version: {details.docker.server.server_version}<br>
            Kernel version: {details.docker.server.kernel_version}
        </div>
    </Card>
    <Card bind:loaded={loaded}
          title="Docker">
        <div>
            Docker version: {details.docker.client.version}<br>
            Docker min. api version: {details.docker.client.api_version}<br>
            OS: {details.docker.client.operating_system} ({details.docker.client.architecture})<br>
            <hr>
            Images: {details.docker.images}<br>
            Containers:
            <Badge>{details.docker.containers.total}</Badge>
            <br/>
            <Badge level="success">Running: {details.docker.containers.running}</Badge>
            <Badge level="warning">Paused: {details.docker.containers.paused}</Badge>
            <Badge level="error">Stopped: {details.docker.containers.stopped}</Badge>
        </div>
    </Card>
    <Card classes={["col-span-1","md:col-span-2","lg:col-span-3"]} title="Terminal">
        {#if showTerminal}
            <Terminal server="{data.server}" onClose="{function(){
            showTerminal = false
        }}"></Terminal>
        {:else}
            <div class="flex justify-center items-center h-8">
                <button class="btn btn-primary" on:click="{function(){showTerminal = true}}">
                    Open Terminal
                </button>
            </div>
        {/if}
    </Card>

    <Card classes={["col-span-1","md:col-span-2","lg:col-span-3"]} title="Files" loaded="{showFs}">
        <Fs server="{data.server}"/>

        <div slot="tr-action">
            <button class="btn btn-sm btn-ghost" on:click="{async function(){
                showFs=false
                await new Promise(r => setTimeout(r, 100));
                showFs=true
            }}">
                Refresh
            </button>
        </div>
    </Card>

</div>