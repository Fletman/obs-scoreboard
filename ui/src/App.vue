<template>
  <div id="main">
    <ScoreboardEdit
      class="scoreboard"
      v-if="edit === true && active_scoreboard"
      :score-id="active_scoreboard['score-id']"
      :data="active_scoreboard"/>
    <ScoreboardStatic
      class="scoreboard"
      v-else-if="active_scoreboard"
      :score-id="active_scoreboard['score-id']"
      :data="active_scoreboard"/>
    <div id="main-view" v-else>
      <header>
          <h1 class="title text-glow">HALO</h1>
          <h3 class="title text-glow">Tournament of the Chosen</h3>
          <h6 class="title">Sponsored by Flet Inc.™</h6>
      </header>
      <div id="scoreboard-list">
        <ScoreboardItem
          class="scoreboard"
          v-for="[id, scoreboard] in scores"
          :key="id"
          :score-id="id"
          :data="scoreboard"
          v-on:click="select_scoreboard(id)"/>
      </div>
    </div>
  </div>
</template>

<script>
import ScoreboardAPI from './js/api';
import ScoreListener from './js/listener';
import ScoreboardItem from './components/ScoreboardItem.vue';
import ScoreboardStatic from './components/ScoreboardStatic.vue';
import ScoreboardEdit from './components/ScoreboardEdit.vue'

export default {
  name: 'App',
  components: {
    ScoreboardItem,
    ScoreboardStatic,
    ScoreboardEdit
  },

  data() {
    return {
      listener: new ScoreListener(),

      handler: new ScoreboardAPI(),

      scoreboards: {},

      active_scoreboard: null,

      edit: false
    }
  },

  methods: {
    route() {
      const path = window.location.pathname;
      if(path === '/') {
        const query_params = new URLSearchParams(window.location.search);
        if(query_params.get('featured') && query_params.get('featured').toLowerCase() === 'true') {
          this.load_featured_scoreboard();
        } else {
          this.load_scores();
        }
        return;
      }

      const path_param = '/[^/]+';
      const score_id_regex = new RegExp(path_param);
      const score_id = path.match(score_id_regex)[0].substring(1);
      const editable = path.match(new RegExp(`${path_param}/edit`)) != null;
      this.load_scoreboard(score_id, editable);
    },

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

    load_scoreboard(id, editable) {
      this.handler.get_scoreboard(id)
        .then((scoreboard) => {
          if(scoreboard) {
            this.edit = editable;
            this.active_scoreboard = scoreboard;
          }
        }).catch((err) => {
          console.error(err);
        });
    },

    load_featured_scoreboard() {
      this.handler.get_featured_scoreboard()
        .then((scoreboard) => {
          if(scoreboard) {
            this.active_scoreboard = scoreboard;
          }
        }).catch((err) => {
          console.error(err);
        });
    },

    select_scoreboard(id) {
      window.location.href = `${window.location.origin}/${id}/edit`;
    }
  },

  created() {
    this.route();

    this.listener.on('message', (event) => {
      const score = JSON.parse(event.data);
      if(score.featured && this.active_scoreboard) {
        this.active_scoreboard = score;
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
    background-image: url('./assets/ring_bg.webp');
    background-repeat: no-repeat;
    background-attachment: fixed;
    background-size: cover;
    overflow-y: auto;
  }

  #scoreboard-list {
    display: flex;
    width: 100%;
    max-height: 80%;
    flex-flow: row wrap;
    justify-content: space-evenly;
  }

  .scoreboard {
    padding: 2px;
    text-align: center;
  }

  .text-glow {
    filter: drop-shadow(1px 1px 10px white);
  }

  .title {
    margin: 0px;
    padding: 0px;
    color: #ddd;
  }

  header {
    text-align: center;
    overflow: hidden;
    font-family: 'Halo', sans-serif;
    font-size: calc(1vh + 1vw);
    height: 20%;
  }

  html, body {
    margin: 0px;
    padding: 0px;
  }
</style>
