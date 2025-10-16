<script lang="ts">

    import {getFiles} from "$lib/api/servers";
    import iconFileAdd from "$lib/assets/icons/file_add.svg"
    import iconFolderAdd from "$lib/assets/icons/folder_add.svg"
    import iconDelete from "$lib/assets/icons/delete.svg"

    import iconFolder from "$lib/assets/icons/folder.svg"
    import iconFolderOpen from "$lib/assets/icons/folder_open.svg"
    import {onMount} from "svelte";

    export let server = "";
    export let path = "";
    export let name = "";
    export let isDir = false;
    let expanded = false;
    export let onClickFile: (str: string | null) => void = (str) => {
        console.log("component not set click action: ", str)
    };
    export let onClickFolder: (str: string | null) => void = (str) => {
        console.log("component not set click on folder action: ", str)
    }

    export let onAddFile: (path: string | null) => void = (path) => {
        console.log("adding file ", path)
    }
    export let onAddFolder: (path: string | null) => void = (path) => {
        console.log("adding folder ", path)
    }
    export let onDeleteFile: (path: string | null) => void = (path) => {
        console.log("deleting folder ", path)
    }

    let files = [];

    async function loadContent(path: string) {
        files = await getFiles(server, path)
    }

    onMount(async () => {
        if (isDir && expanded) {
            loadContent(path)
        }
    })
    let loaded = false

    async function toggle() {
        expanded = !expanded;
        loaded = false;
        if (expanded) await loadContent(path)
        loaded = true;
    }
</script>
<li>
    {#if isDir}
        <div class="w-full flex justify-end on_hover">
            <span class="w-3/4 flex gap-2" on:click={()=>{toggle(); onClickFolder(path)}}>
            {#if expanded}
                <img src={iconFolderOpen} width="16px" height="16px"/>
            {:else}

                <img src={iconFolder} width="16px" height="16px"/>
            {/if}
                <div>
                {name}
            </div>
            </span>
            <div class="w-1/4 hidden">
                        <span on:click={function(e){e.preventDefault();if(confirm(`You are about to delete ${path} and its content? this action cannot be undone.`)){onDeleteFile(path)}}}>
                            <img class="inline" src="{iconDelete}" width="16px" height="16px"/>
                        </span>
                <span on:click={function(e){e.preventDefault();onAddFolder(prompt("new dir",path))}}>
                                <img class="inline" src="{iconFolderAdd}" width="16px" height="16px"/>
                        </span>

                <span on:click={function(e){e.preventDefault();onAddFile(prompt("new file",path))}}>
                                <img class="inline" src="{iconFileAdd}" width="16px" height="16px"/>
                        </span>
            </div>
        </div>
        {#if expanded}
            {#if loaded}
                <ul>
                    {#each files as file}
                        <svelte:self name="{file.name}" path="{path + '/'+file.name}" isDir="{file.is_dir}"
                                     server="{server}"
                                     onClickFile={onClickFile}
                                     onClickFolder={onClickFolder}
                                     onAddFile={onAddFile}
                                     onAddFolder={onAddFolder}
                                     onDeleteFile={onDeleteFile}
                        />
                    {/each}

                </ul>
            {:else}
                <span class="loading loading-xs"></span>
            {/if}

        {/if}
    {:else}

        <div class="flex gap-2 justify-between on_hover">
            <a on:click={function(){onClickFile(path)}} class="w-4/5">
                {name}
            </a>
            <span class="w-1/5 hidden"
                  on:click={function(e){e.preventDefault();if(confirm(`You are about to delete ${path}? this action cannot be undone.`)){onDeleteFile(path)}}}>
                        <img src="{iconDelete}" width="16px" height="16px"/>
                    </span>
        </div>
    {/if}
</li>

<style>
    .on_hover:hover .hidden {
        display: inline-block;
    }

</style>