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
  name: 'ScoreFeaturedPage',
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
    load_featured_scoreboard() {
      this.handler.get_featured_scoreboard()
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
    this.load_featured_scoreboard();

    this.listener.on('message', (event) => {
      const score = JSON.parse(event.data);
      if(score.featured && this.active_scoreboard) {
        this.active_scoreboard = score;
      }
    });
    this.listener.connect();
  }
}
</script>

<style>
  @import '../css/main.css';
  @import '../css/fonts.css';
</style>
