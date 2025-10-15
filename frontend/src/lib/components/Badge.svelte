<script lang="ts">
    // This component represents a badge with primary styling.
    // It uses a slot to allow for flexible content insertion.
    export let level: 'primary' | 'secondary' | 'accent' | 'info' | 'success' | 'warning' | 'error' = 'primary';

    function getLevelClass(level: string): string {
        switch (level) {
            case 'primary':
                return 'badge-primary';
            case 'secondary':
                return 'badge-secondary';
            case 'accent':
                return 'badge-accent';
            case 'info':
                return 'badge-info';
            case 'success':
                return 'badge-success';
            case 'warning':
                return 'badge-warning';
            case 'error':
                return 'badge-error';
            default:
                return 'badge-primary';
        }

    }
</script>
<span class="badge badge-soft {getLevelClass(level)}">
    <slot></slot>
</span>
