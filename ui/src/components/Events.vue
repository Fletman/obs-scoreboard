<template>
  <div>
      <textarea readonly v-model="score"></textarea>
  </div>
</template>

<script>
import ScoreListener from '../js/listener';

export default {
  name: 'Events',
  components: {
    
  },
  data() {
    return {
      listener: new ScoreListener(),
      msg: ""
    }
  },
  created() {
    this.listener.on('message', (event) => {
      console.log(event);
      this.msg = event.data;
    });

    this.listener.connect();
  },
  computed: {
    score() {
      return this.msg;
    }
  }
}
</script>

<style>

</style>
