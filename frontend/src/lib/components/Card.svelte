<script lang="ts">
    export let title: string | null = null;
    export let width: string | null = null;
    export let loaded: boolean = true;
    export let classes: string[] = [];
</script>

<div class="card {width?'w-'+width:''} bg-base-100 shadow {classes.join(' ')}">
    <div class="card-body">
        {#if title }
            <div class="card-title flex justify-between">
                <h2>{title}</h2> <span>
                <slot name="tr-action"></slot>
            </span>
            </div>
        {/if}
        {#if loaded}
            <slot></slot>
        {:else}

            <div class="flex justify-center items-center h-32">
                <span class="loading loading-lg"></span>
            </div>
        {/if}
    </div>
</div>