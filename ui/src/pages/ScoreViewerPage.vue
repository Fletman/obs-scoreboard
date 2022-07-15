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
  name: 'ScoreViewerPage',
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
  },

  computed: {
    score_id() {
      return this.$route.params.id;
    },

    score_layout() {
      const query_strs = window.location.href.substring(window.location.href.indexOf('?'));
      const query_params = new URLSearchParams(query_strs);
      const layout = query_params.get('layout');
      console.log(layout);
      return layout ? layout.toLowerCase() : 'default';
    }
  }
}
</script>

<style>
  @import '../css/main.css';
  @import '../css/fonts.css';
</style>
