<template>
  <div class="scoreboard-content" :style="set_size">
    <span>{{ scoreId }}: </span>
    <span v-for="(team, index) in data.teams" :key=index>
      {{ team.name }} ({{ team.score }})
      <span v-if="index+1 != data.teams.length"> vs. </span>
    </span>
    <span v-if="data.completed"> | Completed</span>
    <span v-else> | In Progress</span>
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
      set_size() {
          return this.focused ?
            "width:80%; height:90vh;" :
            "width:25%; height:10vh;";
      }
  }
}
</script>

<style>
    .scoreboard-content {
        padding: 2px;
        text-align: center;
        background-color: rgba(216, 216, 216, 0.875);
    }

    .scoreboard-content:hover {
        background-color: rgba(216, 216, 216, 1);
        cursor: pointer;
    }
</style>
