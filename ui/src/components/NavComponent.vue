<template>
  <div id="nav-menu">
    <router-link class="link-style" to="/scores">
      <div class="nav-link secondary-halo-font">Scores</div>
    </router-link>
    <router-link class="link-style" to="/brackets">
      <div class="nav-link secondary-halo-font">Brackets (Experimental)</div>
    </router-link>
  </div>
</template>

<script>
export default {
  name: 'NavComponent'
}
</script>

<style>
  @import '../css/fonts.css';

  #nav-menu {
    display: grid;
    grid-auto-flow: column;
    grid-auto-columns: 1fr;
    margin: auto;
    margin-bottom: 10px;
    padding: 0px;
    width: 100%;
    height: auto;
    font-size: 1em;
    background-color: rgba(23, 32, 48, 0.75);
  }

  .nav-link {
    margin: auto;
    padding: 10px;
    border: solid rgba(128, 128, 128, 0.5);
    border-width: 0px 1px;
    font-size: 1.125em;
    text-align: center;
    color: white;
    text-decoration: none;
  }

  .nav-link:hover {
    background-color: rgba(290, 116, 61, 0.5);
    cursor: pointer;
  }
  
</style>
