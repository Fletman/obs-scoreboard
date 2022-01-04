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

export default {
  name: 'Scoreboard',
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

 .scoreboard-listed {
      width: 25%;
      height: fit-content;
      font-size: 1.5em;
      background-color: rgba(216, 216, 216, 0.875);
  }

  .scoreboard-listed:hover {
      background-color: rgba(216, 216, 216, 1);
      cursor: pointer;
  }
</style>
