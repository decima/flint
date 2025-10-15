<script lang="ts">

    import {getFiles} from "$lib/api/servers";
    import iconFileAdd from "$lib/assets/icons/file_add.svg"
    import iconFolderAdd from "$lib/assets/icons/folder_add.svg"
    import iconDelete from "$lib/assets/icons/delete.svg"
    import {onMount} from "svelte";

    export let server = "";
    export let path = "";
    export let name = "";
    export let isDir = false;
    export let onClickFile: (str: string | null) => void = (str) => {
        console.log("component not set click action: ", str)
    };
    export let onClickFolder:(str:string|null)=>void=(str)=>{
        console.log("component not set click on folder action: ",str)
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
        if (isDir) {
            loadContent(path)
        }
    })
</script>
<li>
    {#if isDir}
        <div class="w-full flex justify-between">
            <span class="w-full" on:click={()=>{onClickFolder(path)}}>{name}</span>
            <div class="flex gap-2">
                    <span on:click={function(e){e.preventDefault();if(confirm(`You are about to delete ${path} and its content? this action cannot be undone.`)){onDeleteFile(path)}}}>
                        <img src="{iconDelete}" class="h-4"/>
                    </span>
                <span on:click={function(e){e.preventDefault();onAddFolder(prompt("new dir",path))}}>
                            <img src="{iconFolderAdd}" class="h-4"/>
                    </span>

                <span on:click={function(e){e.preventDefault();onAddFile(prompt("new file",path))}}>
                            <img src="{iconFileAdd}" class="h-4"/>
                    </span>
            </div>
        </div>
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

        <div class="flex gap-2 justify-between">
            <a on:click={function(){onClickFile(path)}}>
                {name}
            </a>
            <span on:click={function(e){e.preventDefault();if(confirm(`You are about to delete ${path}? this action cannot be undone.`)){onDeleteFile(path)}}}>
                        <img src="{iconDelete}" class="h-4"/>
                    </span>
        </div>
    {/if}
</li>