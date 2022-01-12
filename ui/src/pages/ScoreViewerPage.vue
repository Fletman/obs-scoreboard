<template>
  <div id="main">
    <ScoreboardStatic
      class="scoreboard"
      :score-id="score_id"
      :data="active_scoreboard"/>
  </div>
</template>

<script>
import ScoreboardAPI from '../js/api';
import ScoreListener from '../js/listener';
import ScoreboardStatic from '../components/ScoreboardStatic.vue';

export default {
  name: 'ScoreViewerPage',
  components: {
    ScoreboardStatic
  },

  data() {
    return {
      listener: new ScoreListener(),

      handler: new ScoreboardAPI(),

      active_scoreboard: null
    }
  },

  methods: {
    load_scoreboard(id) {
      this.handler.get_scoreboard(id)
        .then((scoreboard) => {
          if(scoreboard) {
            this.active_scoreboard = scoreboard;
          }
        }).catch((err) => {
          console.error(err);
        });
    }
  },

  beforeMount() {
    this.load_scoreboard(this.score_id);

    this.listener.on('message', (event) => {
      const score = JSON.parse(event.data);
      if(score['score-id'] === this.score_id) {
        this.active_scoreboard = score;
      }
    });
    this.listener.connect();
  },

  computed: {
    score_id() {
      return this.$route.params.id;
    }
  }
}
</script>

<style>
  @import '../css/main.css';
  @import '../css/fonts.css';
</style>
