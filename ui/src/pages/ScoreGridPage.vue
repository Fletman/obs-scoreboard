<template>
  <div id="main">
    <div id="main-view">
      <NavComponent/>
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
import ScoreboardAPI from '../js/api';
import ScoreListener from '../js/listener';
import ScoreboardItem from '../components/ScoreboardItem.vue';
import NavComponent from '../components/NavComponent.vue';

export default {
  name: 'ScoreGridPage',
  components: {
    ScoreboardItem,
    NavComponent
  },

  data() {
    return {
      listener: ScoreListener.get_socket_handler(),

      handler: ScoreboardAPI.get_api_handler(),

      scoreboards: {}
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
              this.$router.push(`/scores/${score_id}/edit`);
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
    this.load_scores();

    this.listener.on('message', (event) => {
      const score = JSON.parse(event.data);
      this.scoreboards[score['score-id']] = score;
    });
  },

  computed: {
      scores() {
          return Object.entries(this.scoreboards);
      }
  }
}
</script>

<style>
  @import '../css/main.css';
  @import '../css/fonts.css';

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
</style>
