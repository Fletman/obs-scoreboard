<template>
  <div class="scoreboard-focused primary-halo-font">
    <div id="team-list" v-if="['default'].includes(layout)">
      <div :style="team_style(index)" v-for="(team, index) in data.teams" :key=index>
        <div>{{ team.name }}</div>
        <div>{{ team.score }}</div>
      </div>
    </div>
    <div id="team-list" v-else>
      <div :style="team_style(index)" v-for="(team, index) in data.teams" :key=index>
        <div v-if="index % 2 === 0">{{ team.name }}</div>
        <div v-else>{{ team.score }}</div>
        <div v-if="index % 2 === 0">{{ team.score }}</div>
        <div v-else>{{ team.name }}</div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ScoreboardStatic',
  components: {
    
  },

  props: {
      data: {
        type: Object,
        required: true
      },
      layout: {
        type: String,
        default: 'default'
      }
  },

  data() {
    return {
        
    }
  },

  methods: {
      team_style(index) {
        let width_f;
        let border_f;
        let display_f;
        let bg_f;

        width_f = () => {
          const width_pct = this.data.teams.length != 0 ?
            Math.floor(100/this.data.teams.length):
            0;
          return `width: ${width_pct}%; font-size: ${width_pct/20}vw`;
        };

        switch(this.layout) {
          case 'default':
            border_f = () => 'border: none';
            display_f = () => 'display: grid';
            bg_f = () => 'background-color:none';
            break;
          case 'banner':
            border_f = () => 'border: none';
            display_f = (i) => {
              const display_cols = i % 2 === 0 ? '70% 30%;' : '30% 70%;';
              return `display: grid; grid-template-columns: ${display_cols}`;
            };
            bg_f = (i) => {
              const gradient = i % 2 === 0 ?
                { direction: 'right,', color_1: 'rgba(128, 0, 0, 0.875),'.repeat(4), color_2: 'rgba(255, 255, 255, 0)' }:
                { direction: 'left,', color_1: 'rgba(0, 0, 128, 0.875),'.repeat(4), color_2: 'rgba(255, 255, 255, 0)' };
              return `background-image: linear-gradient(to ${gradient.direction} ${gradient.color_1} ${gradient.color_2})`;
            }
            break;
          default:
            throw(`Unknown layout ID: ${this.layout}`);
        }
        return `${width_f()}; ${border_f()}; ${display_f(index)}; ${bg_f(index)};`;
      }
  },

  created() {
      console.log("Loading ScoreboardStatic");
  },

  computed: {
    
  }
}
</script>

<style>
  @import '../css/fonts.css';

  #team-list {
    display: flex;
    flex-flow: row;
    justify-content: space-evenly;
  }

  .scoreboard-focused {
    padding: 0px;
    width: 100%;
    height: auto;
    color: white;
    text-shadow: 2px 2px black;
  }
</style>
