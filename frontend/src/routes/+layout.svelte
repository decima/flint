<script lang="ts">
    import Aside from './Aside.svelte';
    import '../app.css';

    let {children} = $props();
</script>

<div class="app flex">
    <Aside/>

    <main class="p-4">
        {@render children()}
    </main>


</div>