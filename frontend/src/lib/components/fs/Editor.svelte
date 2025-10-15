<script lang="ts">

    import {onMount} from "svelte";
    import {basicSetup, EditorView} from "codemirror"
    import {ViewPlugin, ViewUpdate,keymap} from "@codemirror/view"
    import {indentWithTab} from "@codemirror/commands"
    import iconSave from "$lib/assets/icons/save.svg"


    let {
        value = $bindable("//write some code here!"),
        originalValue = "",
        onSave = (content: string) => {
            console.log("save pressed with content " + content)
        },
    } = $props();


    let editorContainer: Element;

    onMount(() => {
        const view = new EditorView({
            doc: value,
            parent: editorContainer,
            extensions: [basicSetup,
                keymap.of([indentWithTab]),
                ViewPlugin.fromClass(class {
                    constructor(view: EditorView) {
                    }

                    update(update: ViewUpdate) {
                        if (update.docChanged) {
                            value = update.state.doc.toString()
                        }

                    }
                })
            ]
        })

    })

    async function onKeyPress(evt: KeyboardEvent) {
        if (evt.ctrlKey && evt.key === 's') {
            evt.preventDefault()
            await onSave(value)
            return;
        }
    }


</script>

<svelte:window
        on:keydown={onKeyPress}
/>

<div class="bg-base-200 h-6 flex justify-end">
    {#if originalValue != value}
        <span class="bg-danger">
            Changes not saved
        </span>
        <button class="btn btn-xs btn-ghost" on:click={onSave(value)}>save now <img src={iconSave}
                                                                                                       class="h-6"/>
        </button>
    {/if}
</div>
<div bind:this={editorContainer} class="editor-wrapper"></div>
