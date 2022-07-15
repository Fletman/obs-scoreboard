<template>
  <div id="main">
    <!--
    <div id="bracket-round-row" class="bracket-view-row primary-halo-font">
      <div v-for="(round, index) in bracket.rounds" :key="index">
        <div style="width:200px">
          Round
        </div>
      </div>
    </div>
    -->
    <div class="bracket-view-row secondary-halo-font">
      <div class="bracket-view-col" v-for="(round, index) in bracket.rounds" :key="index">
        <div v-for="match in round['match-ids']" :key="match">
          <div v-if="Object.keys(scoreboards).length > 0" class="bracket-matchup-block">
            <div class="bracket-matchup-team" v-for="team in scoreboards[match].teams" :key="team.name">
              <div>
                <span>{{ team.name }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ScoreboardAPI from '../js/api';
import ScoreListener from '../js/listener';

export default {
  name: 'BracketViewPage',
  components: {
      
  },

  data() {
    return {
      listener: ScoreListener.get_socket_handler(),

      handler: ScoreboardAPI.get_api_handler(),

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
            let score_ids = [];
            for(const round of bracket.rounds) {
              score_ids = score_ids.concat(round['match-ids']);
            }
            this.handler.list_scoreboards(score_ids)
              .then((scoreboards) => {
                for(const score of scoreboards) {
                  this.scoreboards[score['score-id']] = score;
                }
              }).catch((err) => console.error(err));
          }).catch((err) => {
              console.error(err);
          });
    }
  },

  beforeMount() {
    this.listener.on('message', (event) => {
        console.log(event);
    });

    this.load_bracket_scores();
  },

  computed: {
    
  }
}
</script>

<style>
  @import '../css/main.css';
  @import '../css/fonts.css';

  #bracket-round-row {
    height: 5vh;
    font-size: 1.25em;
  }
  
  .bracket-view-row {
    width: 100%;
    display: flex;
    flex-flow: row;
    justify-content: space-evenly;
  }

  .bracket-view-col {
    height: 100vh;
    display: flex;
    flex-flow: column;
    justify-content: space-evenly;
    overflow: auto;
  }

  .bracket-matchup-team {
    height: 50%;
    display: flex;
    flex-flow: column;
    justify-content: space-evenly;
    font-size: 1.75em;
  }
</style>
