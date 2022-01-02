<template>
  <div>
    <header>
      <h1>Scoreboard</h1>
      <h4>Sponsored by Flet Inc.™</h4>
    </header>
    <div id="scoreboard-list" v-if="!active_scoreboard">
      <Scoreboard class="scoreboard" v-for="[id, scoreboard] in scores" :key="id" :score-id="id" :data="scoreboard"/>
    </div>
    <Scoreboard v-else :score-id="id" :data="active_scoreboard" :focused="true"/>
  </div>
</template>

<script>
import ScoreboardAPI from './js/api';
import ScoreListener from './js/listener';
import Scoreboard from './components/Scoreboard.vue';

export default {
  name: 'App',
  components: {
    Scoreboard
  },

  data() {
    return {
      listener: new ScoreListener(),

      handler: new ScoreboardAPI(),

      scoreboards: {},

      active_scoreboard: null,
    }
  },

  methods: {
    load_scores() {
        this.handler.list_scoreboards()
        .then((scoreboards) => {
            console.log(scoreboards)
            this.scoreboards = scoreboards;
        }).catch((err) => {
            console.error(err);
        });
    }
  },

  created() {
    this.load_scores();

    this.listener.on('message', (event) => {
      const score = JSON.parse(event.data);
      this.scoreboards[score['score-id']] = score.scoreboard;
    });
    this.listener.connect();
  },

  computed: {
      scores() {
          return Object.entries(this.scoreboards);
      }
  }
}
</script>

<style>
  #main {
    position: fixed;
    height:100vh;
    width: 100vw;
    margin: 0px;
    padding: 5px;
    background-image: url('./assets/halo_bg_gs.webp');
    background-repeat: no-repeat;
    background-attachment: fixed;
    background-size: cover;
  }

  #scoreboard-list {
    margin: auto;
    display: flex;
    flex-flow: row wrap;
    justify-content: space-evenly;
  }

  .scoreboard {
    margin: 10px 5px;
  }

  header {
    margin: auto;
    text-align: center;
  }

  html, body {
    margin: 0px;
    padding: 0px;
  }
</style>
