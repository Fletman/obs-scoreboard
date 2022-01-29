<template>
  <div id="main">
    <div id="main-view">
      <NavComponent/>
      <header class="primary-halo-font">
        <h1 class="title text-glow">HALO</h1>
        <h3 class="title text-glow">Tournament of the Chosen</h3>
        <h6 class="title">Sponsored by Flet Inc.™</h6>
      </header>
      <div id="rounds">
        <div class="round-col" v-for="(round, index) in bracket.rounds" :key="index">
          <div v-for="match in round['match-ids']" :key="match">
            <img src="../assets/logo.png" width="50" height="50"/>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ScoreboardAPI from '../js/api';
import ScoreListener from '../js/listener';
//import ScoreboardItem from '../components/ScoreboardItem.vue';
import NavComponent from '../components/NavComponent.vue';

export default {
  name: 'BracketViewPage',
  components: {
    //ScoreboardItem,
    NavComponent
  },

  data() {
    return {
      listener: ScoreListener.get_socket_handler(),

      handler: new ScoreboardAPI(),

      scoreboards: {},

      bracket: {}
    }
  },

  methods: {
    load_bracket_scores() {
        const bracket_id = this.$route.params.id;
        this.handler.get_bracket(bracket_id)
          .then((bracket) => {
            this.bracket = bracket;
            Promise.all(bracket.rounds.map(async (round) => {
              return await Promise.all(round['match-ids'].map((id) => this.handler.get_scoreboard(id)));
            })).then((score_data) => {
                for(const score_set of score_data) {
                  for(const score of score_set) {
                    this.scoreboards[score['score-id']] = score;
                  }
                }
                console.log(this.scoreboards);
              }).catch((err) => {
                console.error(err);
              });
          }).catch((err) => {
              console.error(err);
          });
    }
  },

  beforeMount() {
    this.load_bracket_scores();

    this.listener.on('message', (event) => {
      const score = JSON.parse(event.data);
      console.log(score)
    });
  },

  computed: {
    
  }
}
</script>

<style>
  @import '../css/main.css';
  @import '../css/fonts.css';

  #rounds {
    display: grid;
    grid-auto-flow: column;
    grid-auto-columns: 1fr;
    margin: auto;
    height: 900px;
  }

  .round-col {
    background-color: white;
    display: flex;
    flex-flow: column;
    justify-content: space-evenly;
    margin-left: auto;
    margin-right: auto;
  }
</style>
