<script lang="ts">
    import type { PageProps } from './$types';
    import Card from "$lib/components/Card.svelte";

    let { data }: PageProps = $props();
</script>

<Card title={data.title}>

</Card>
