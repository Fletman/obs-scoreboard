<template>
  <div class="scoreboard-focused primary-halo-font">
    <div id="team-list">
      <div :style="set_sizes" v-for="(team, index) in data.teams" :key=index>
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
      }
  },

  data() {
    return {
        
    }
  },

  methods: {
      
  },

  created() {
      console.log("Loading ScoreboardStatic");
  },

  computed: {
    set_sizes() {
      const width_pct = this.data.teams.length != 0 ?
        Math.floor(100/this.data.teams.length):
        0;
      return `width: ${width_pct}%; font-size: ${width_pct/20}vw`;
    }
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
    width: 100%;
    height: auto;
    color: white;
    text-shadow: 2px 2px black;
  }
</style>
