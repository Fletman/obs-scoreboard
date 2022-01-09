<template>
  <div class="scoreboard-focused primary-halo-font">
    <div>{{ scoreId }}</div>
    <div>
      <span v-for="(team, index) in data.teams" :key=index>
        {{ team.name }} ({{ team.score }})
        <span v-if="index+1 != data.teams.length"> vs. </span>
      </span>
    </div>
  </div>
</template>

<script>
import '../styling/fonts.css';

export default {
  name: 'ScoreboardStatic',
  components: {
    
  },

  props: {
      'score-id': {
          type: String,
          required: true
      },
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

  }
}
</script>

<style>  
  .scoreboard-focused {
    width: 100%;
    height: auto;
    font-size: 2.5em;
    color: white;
    text-shadow: 2px 2px black;
  }
</style>
