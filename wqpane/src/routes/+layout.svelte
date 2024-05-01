<script lang="ts">
  import Navigation from '$lib/basic/Navigation.svelte';
  import '../app.scss'; // Ensure the path to your global CSS file is correct
</script>

<Navigation />
<slot></slot>

<style>
  :global(body) {
    padding-top: 60px;
    margin: 0; 
    padding: 0; 
  }
</style>