<template>
  <div :class="set_style">
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
  name: 'ScoreboardEdit',
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
      },
      focused: {
          type: Boolean,
          default: false
      },
      editable: {
        type: Boolean,
        default: false
      }
  },

  data() {
    return {
        
    }
  },

  methods: {
      
  },

  created() {
      console.log("Loading ScoreboardEdit");
      console.log(`Editable: ${this.editable}`);
  },

  computed: {
      set_style() {
        return this.focused ?
          "scoreboard scoreboard-focused" :
          "scoreboard scoreboard-listed"
      }
  }
}
</script>

<style>
  .scoreboard {
    padding: 2px;
    text-align: center;
  }

  .scoreboard-focused {
    width: 100%;
    height: auto;
    font-size: 2.5em;
    color: white;
    text-shadow: 2px 2px black;
  }
</style>
