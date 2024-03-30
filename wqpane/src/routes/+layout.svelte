<script>
  import Navigation from '$lib/basic/Navigation.svelte';
import '../app.scss'; // Ensure the path to your global CSS file is correct
</script>
  
<!-- Your layout and slot -->
<Navigation />
<slot></slot>

<style>
  :global(body) {
    padding-top: 60px; 
  }
</style>