<script lang="ts">
  import { onMount } from 'svelte';
  
  let sy = 0;
  
  onMount(() => {
    const handleWheel = (event: WheelEvent) => {
      if (window.scrollY === 0 && event.deltaY < 0) {
        event.preventDefault();
      }
    }
  
    const handleKeydown = (event: KeyboardEvent) => {
      if ((window.scrollY === 0 && ['ArrowUp', 'Space', 'PageUp'].includes(event.code)) || (event.code === 'Space' && event.shiftKey)) {
        event.preventDefault();
      }
    }
  
    return (() => {
      window.removeEventListener('wheel', handleWheel);
      window.removeEventListener('keydown', handleKeydown);
    });
  });
</script>

<svelte:window bind:scrollY={sy}/>

<div class='parallax-container'>
  <img style='transform: translate(0,{-sy * 0.2}px)' alt="" src='parallax0.png'>
  <img style='transform: translate(0,{-sy * 0.3}px)' alt="" src='parallax1.png'>
  <img style='transform: translate(0,{-sy * 0.4}px)' alt="" src='parallax3.png'>
  <img style='transform: translate(0,{-sy * 0.5}px)' alt="" src='parallax5.png'>
  <img style='transform: translate(0,{-sy * 0.6}px)' alt="" src='parallax7.png'>
</div>

<div class='text'>
  <small style='transform: translate(0,{-sy * 1.5}px); opacity: {1 - Math.max( 0, sy / 80 )}'>
    Portfolio @Goonzard 
  </small>
  
  <span>LLM</span>
  <br />
  <span>Oversea Stock</span>
  <br />
  <span>Crypto</span>
</div>


<style>
  .parallax-container {
    position: fixed;
    width: 2400px;
    height: 712px;
    left: 50%;
    transform: translate(-50%,0);
  }

  .parallax-container img {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    will-change: transform;
  }

  .text {
    position: relative;
    width: 100%;
    min-height: 100vh;
    color: white;
    text-align: center;
    padding: 50vh 0.5em 0.5em 0.5em;
    box-sizing: border-box;
  }

  .text::before {
    content: '';
    position: absolute;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    background: linear-gradient(to bottom, rgba(45,10,13,0) 60vh,rgba(45,10,13,1) 712px);
  }

  small {
    display: block;
    font-size: 4vw;
    will-change: transform, opacity;
  }

  .text span {
    font-size: 15vw;
    position: relative;
    z-index: 2;
  }
</style>
  