<template>
  <div id="main">
    <div id="main-view" v-if="!featured_scoreboard">
      <header>
          <h1 class="title">HALO</h1>
          <h2 class="title">Tournament of the Chosen</h2>
          <h4 class="title">Sponsored by Flet Inc.™</h4>
      </header>
      <div id="scoreboard-list">
        <Scoreboard class="scoreboard-item" v-for="[id, scoreboard] in scores" :key="id" :score-id="id" :data="scoreboard"/>
      </div>
    </div>
    <Scoreboard v-else :score-id="featured_scoreboard['score-id']" :data="featured_scoreboard" :focused="true"/>
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

      featured_scoreboard: null,
    }
  },

  methods: {
    load_scores() {
        this.handler.list_scoreboards()
        .then((scoreboards) => {
            for(const score of scoreboards) {
              this.scoreboards[score['score-id']] = score;
            }
        }).catch((err) => {
            console.error(err);
        });
    },

    load_scoreboard(id) {
      this.handler.get_scoreboard(id)
        .then((scoreboard) => {
          if(scoreboard) {
            this.featured_scoreboard = scoreboard;
            console.log(this.featured_scoreboard);
          }
        }).catch((err) => {
          console.error(err);
        });
    },

    load_featured_scoreboard() {
      this.handler.get_featured_scoreboard()
        .then((scoreboard) => {
          if(scoreboard) {
            this.featured_scoreboard = scoreboard;
            console.log(this.featured_scoreboard);
          }
        }).catch((err) => {
          console.error(err);
        });
    }
  },

  created() {
    const query_params = new URLSearchParams(window.location.search);
    if(query_params.get('featured') && query_params.get('featured').toLowerCase() === 'true') {
      this.load_featured_scoreboard();
    } else if(query_params.get('score-id')) {
      this.load_scoreboard(query_params.get('score-id'));
    } else {
      this.load_scores();
    }

    this.listener.on('message', (event) => {
      const score = JSON.parse(event.data);
      if(score.featured && this.featured_scoreboard) {
        this.featured_scoreboard = score;
      } else {
        this.scoreboards[score['score-id']] = score;
      }
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
  @import url('http://fonts.cdnfonts.com/css/halo');

  #main {
    position: fixed;
    height: 100vh;
    width: 100vw;
    margin: 0px;
    padding: 0px;
  }

  #main-view {
    position: fixed;
    padding: 0px;
    width: 100%;
    height: 100%;
    background-image: url('./assets/halo_bg_gs.webp');
    background-repeat: no-repeat;
    background-attachment: fixed;
    background-size: cover;
  }

  #scoreboard-list {
    margin: auto;
    display: flex;
    width: 100%;
    height:100%;
    flex-flow: row wrap;
    justify-content: space-evenly;
  }

  .scoreboard-item {
    margin: 10px 2px;
  }

  .title {
    margin: 0px;
    padding: 0px;
  }

  header {
    text-align: center;
    overflow: hidden;
  }

  html, body {
    margin: 0px;
    padding: 0px;
    font-family: 'Halo', sans-serif;
    /* font-family: 'Halo Outline', sans-serif; */
  }
</style>
