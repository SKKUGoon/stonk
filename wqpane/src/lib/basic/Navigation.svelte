<script lang="ts">

</script>

<nav>
  <a href="/">Home</a>
  <a href="/oversea">Oversea</a>
  <a href="/crypto">Crypto</a>
</nav>

<style>
  nav {
    position: fixed;
    top: 0;
    left: 50%;
    transform: translateX(-50%);
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: transparent; /* Example background */
    border-bottom: 1px solid #ccc; /* Example border */
    padding: 10px;
    gap: 20px;
    z-index: 1000; /* Ensure it's above other content */
  }

  nav a {
    margin: 0 10px;
    text-decoration: none;
    color: #fff; /* Example text color */
    font-size: 30px; /* Example font size */
    font-weight: bold;
  }
  
  /* Hover state */
  nav a:hover {
    font-weight: bold; /* Change color on hover */
    transform: scale(1.1); /* Slightly increase size */
  }

  /* Active (pressed) state */
  nav a:active {
    color: #0056b3; /* Darker color for pressed state */
    transform: scale(0.9); /* Slightly decrease size */
  }
</style>