<template>
  <div id="main">
    <ScoreboardStatic
      v-if="active_scoreboard"
      class="scoreboard"
      :layout="score_layout"
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
      listener: ScoreListener.get_socket_handler(),

      handler: ScoreboardAPI.get_api_handler(),

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

  computed: {
    score_layout() {
      const query_strs = window.location.href.substring(window.location.href.indexOf('?'));
      const query_params = new URLSearchParams(query_strs);
      const layout = query_params.get('layout');
      console.log(layout);
      return layout ? layout.toLowerCase() : 'default';
    }
  },

  beforeMount() {
    this.load_featured_scoreboard();

    this.listener.on('message', (event) => {
      const score = JSON.parse(event.data);
      if(score.featured) {
        this.active_scoreboard = score;
      }
    });
  }
}
</script>

<style>
  @import '../css/main.css';
  @import '../css/fonts.css';
</style>
