<template>
  <div>
    <HelloWorld msg="Welcome to Your Vue.js App"/>
    <Events/>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import Events from './components/Events.vue'

export default {
  name: 'App',
  components: {
    HelloWorld,
    Events
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
