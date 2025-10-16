<script lang="ts">
    import {deleteFile, getFiles, writeFile} from "$lib/api/servers";
    import {onMount} from "svelte";
    import FileTreeItem from "$lib/components/fs/FileTreeItem.svelte";
    import Editor from "$lib/components/fs/Editor.svelte";

    export let server: string = "";
    let content: string | null = "";
    let originalContent: string | null = "";
    let currentPath: string | null = null;


    async function onClickFile(path: string) {
        content = null
        currentPath = null
        const result = await getFiles(server, path, true)
        content = result
        originalContent = result
        currentPath = path
    }

    async function onAddFolder(path: string | null) {
        if (!path) return;
        onAddFile(path + "/.flintkeep")
    }

    async function onAddFile(path: string | null) {
        if (!path) return;
        if (!path.startsWith("/")) {
            path = "/" + path;
        }
        await writeFile(server, path, "")
        await onClickFile(path)
    }

    async function onSave(newContent: string) {
        if (!currentPath) return;
        await writeFile(server, currentPath, newContent)
        originalContent = newContent
    }

    async function onDeleteFile(path: string | null) {
        if (!path) return;
        if (!path.startsWith("/")) {
            path = "/" + path;
        }
        if (path == currentPath) {
            currentPath = null
        }
        await deleteFile(server, path)

    }
</script>

<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
    <div class="col-span-1 md:col-span-1 lg:col-span-1">
        <ul class="menu menu-sm bg-base-200 rounded-box max-w-xs w-full">
            <FileTreeItem name="/" path="/" isDir="true" server={server}
                          onClickFile="{onClickFile}"
                          onAddFile={onAddFile}
                          onAddFolder={onAddFolder}
                          onClickFolder={()=>{currentPath=null}}
                          onDeleteFile={onDeleteFile}/>

        </ul>
    </div>
    <div class="col-span-1 md:col-span-1 lg:col-span-3">
        {#if currentPath != null }
            <Editor bind:value={content} originalValue={originalContent} onSave="{onSave}"/>
        {/if}
    </div>
</div>