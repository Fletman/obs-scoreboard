<template>
  <div class="scoreboard-edit">
    <div>{{ scoreId }}</div>
    <div>
      <form v-on:change="update_score">
        <span v-for="(team, index) in data.teams" :key="index">
          <input type="text" readonly v-model="team.name"/>
          <input type="number" :readonly="data.completed" v-model="team.score"/>
        </span>
      </form>
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
      handler: {
        type: Object,
        required: true
      }
  },

  data() {
    return {
        
    }
  },

  methods: {
    update_score() {
      this.handler.set_scoreboard(this.data['score-id'], this.data);
    }  
  },

  created() {
      console.log("Loading ScoreboardEdit");
  },

  computed: {
      
  }
}
</script>

<style>

</style>
