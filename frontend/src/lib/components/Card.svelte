<script>
    export let title = null;
    export let width = null;
    export let loaded = true;
</script>

<div class="card {width?'w-'+width:''} bg-base-100 shadow">
    <div class="card-body">
        {#if title }
            <h2 class="card-title">{title}</h2>
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