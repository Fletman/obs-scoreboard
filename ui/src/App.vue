<template>
  <div id="main">
    <div id="main-view" v-if="edit === true && active_scoreboard">
      <ScoreboardEdit
        class="scoreboard"
        :score-id="active_scoreboard['score-id']"
        :data="active_scoreboard"
        :handler="handler"/>
    </div>
    <div v-else-if="active_scoreboard">
      <ScoreboardStatic
        class="scoreboard"
        :score-id="active_scoreboard['score-id']"
        :data="active_scoreboard"/>
    </div>
    <div id="main-view" v-else>
      <header class="primary-halo-font">
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
          :data="scoreboard"/>
        <div id="add-scoreboard" v-on:click="create_scoreboard">
          <span>+</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import './styling/fonts.css';
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
          } else {
            this.active_scoreboard = {};
          }
        }).catch((err) => {
          console.error(err);
        });
    },

    create_scoreboard() {
      const score_id = window.prompt("Enter a scoreboard name");
      if(score_id) {
        this.handler.get_scoreboard(score_id)
        .then((scoreboard) => {
          if(scoreboard) {
            console.error(`Scoreboard ${score_id} already exists`);
          } else {
            this.handler.set_scoreboard(
              score_id,
              {
                teams: [],
                completed: false,
                featured: false
              }
            ).then(() => {
              window.location.href = `${window.location.origin}/${score_id}/edit`;
            }).catch((err) => {
              console.error(err);
            });
          }
        }).catch((err) => {
          console.error(err);
        });
      }
    }
  },

  beforeMount() {
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
    background: url('./assets/ring_bg.webp');
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

  #add-scoreboard {
    display: flex;
    position: relative;
    margin: 10px 2px;
    width: 25%;
    height: 8.75vh;
    justify-content: center;
    align-items: center;
    background-color: #172030;
    color: white;
    font-size: calc(1.25vh + 1.25vw);
    text-align: center;
    filter: drop-shadow(1px 1px 5px gray);
  }

  #add-scoreboard:hover {
    cursor: pointer;
    background-color: #283141;
  }

  .scoreboard {
    padding: 2px;
    text-align: center;
  }

  .title {
    margin: 0px;
    padding: 0px;
    color: #ddd;
  }

  header {
    text-align: center;
    overflow: hidden;
    font-size: calc(1vh + 1vw);
    height: 20%;
  }

  html, body {
    margin: 0px;
    padding: 0px;
  }
</style>
