<template>
  <div id="main">
    <div id="main-view">
      <ScoreboardEdit
        v-if="active_scoreboard"
        class="scoreboard"
        :score-id="active_scoreboard['score-id']"
        :data="active_scoreboard"
        :handler="handler"/>
    </div>
  </div>
</template>

<script>
import ScoreboardAPI from '../js/api';
import ScoreListener from '../js/listener';
import ScoreboardEdit from '../components/ScoreboardEdit.vue'

export default {
  name: 'ScoreEditorPage',
  components: {
    ScoreboardEdit
  },

  data() {
    return {
      listener: ScoreListener.get_socket_handler(),

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
