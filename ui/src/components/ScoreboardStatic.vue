<template>
  <div class="scoreboard-focused primary-halo-font">
    <div id="team-list" v-if="layout==='default'">
      <div :style="team_style(index)" v-for="(team, index) in data.teams" :key=index>
        <div>{{ team.name }}</div>
        <div>{{ team.score }}</div>
      </div>
    </div>
    <div id="team-list" v-else-if="layout==='banner'">
      <div :style="team_style(index)" v-for="(team, index) in data.teams" :key=index :set="is_even = index % 2 === 0">
        <div v-if="is_even"></div>
        <div v-if="is_even">{{ team.name }}</div>
        <div v-else>{{ team.score }}</div>
        <div v-if="is_even">{{ team.score }}</div>
        <div v-else>{{ team.name }}</div>
        <div v-if="!is_even"></div>
      </div>
    </div>
    <div v-else-if="layout==='block'">
      <div :style="team_style(index)" v-for="(team, index) in data.teams" :key=index :set="is_even = index % 2 === 0">
        <div>{{ team.name }}</div>
        <div>{{ team.score }}</div>
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
        let width_f = () => {
          const width_pct = this.data.teams.length != 0 ?
            Math.floor(100/this.data.teams.length):
            0;
          return `width: ${width_pct}%; font-size: ${width_pct/20}vw`;
        };
        let border_f = () => 'border: none';
        let display_f = () => 'display: grid';
        let bg_f = () => 'background-color:none';

        switch(this.layout) {
          case 'default':
            break;
          case 'banner':
            display_f = () => 'display: grid; grid-template-columns:30% 40% 30%';
            bg_f = (i) => {
              const gradient = i % 2 === 0 ?
                {
                  direction: 'right,',
                  color_0: 'rgba(255, 255, 255, 0),',
                  color_1: 'rgba(128, 0, 0, 0.875),'.repeat(3),
                  color_2: 'rgba(255, 255, 255, 0)'
                }:
                {
                  direction: 'left,',
                  color_0: 'rgba(255, 255, 255, 0),',
                  color_1: 'rgba(0, 0, 128, 0.875),'.repeat(3),
                  color_2: 'rgba(255, 255, 255, 0)'
                };
              return `background-image: linear-gradient(to ${gradient.direction} ${gradient.color_0} ${gradient.color_1} ${gradient.color_2})`;
            }
            break;
          case 'block':
            width_f = () => `width:${100/(this.data.teams.length * 4)}%; font-size:2em;`;
            border_f = () => 'border: solid 2px rgba(128, 128, 128, 0.5)';
            display_f = (i) => i % 2 === 0 ? 'display:inline-block; float:left; margin-left:1%' : 'display:inline-block; float:right; margin-right:1%';
            bg_f = (i) => i % 2 === 0 ? 'background-color:rgba(128, 0, 0, 0.875)' : 'background-color:rgba(0, 0, 128, 0.875)';
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
