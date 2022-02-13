<template>
  <div id="main">
    <div id="main-view">
      <NavComponent/>
      <div id="rounds" v-if="bracket && bracket.rounds">
        <div class="round-col" v-for="(round, index) in bracket.rounds" :key="index">
          <ScoreboardItem
            class="scoreboard"
            v-for="id in round['match-ids']"
            v-show="scoreboards[id]"
            :key="id"
            :score-id="id"
            :data="scoreboards[id] || {}"/>
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
  name: 'BracketEditPage',
  components: {
    ScoreboardItem,
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
    },

    match_alternate(matches) {
      const ordered = new Array(matches.length);
      matches.forEach((m, i) => {
        const index = i % 2 === 0 ?
          i:
          matches.length - i;
        ordered[index] = m;
      });
      return ordered;
    }
  },

  beforeMount() {
    this.load_bracket_scores();

    this.listener.on('message', (event) => {
      const score = JSON.parse(event.data);
      if(score['score-id'] in this.scoreboards) {
        this.scoreboards[score['score-id']] = score;
      }
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
  }

  .round-col {
    display: flex;
    flex-flow: column;
    justify-content: space-evenly;
    margin-left: auto;
    margin-right: auto;
  }

  .round-col .scoreboard {
    margin: 5% 0px;
    padding: 0px;
    width: 100%;
  }
</style>
