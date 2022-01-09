<template>
  <div class="scoreboard-listed">
    <span id="featured-marker" v-if="data.featured">💫</span>
    <a id="edit-button" target="_blank" :href="edit_link">
      <img src="../assets/edit_icon.webp" alt="Edit Score"/>
    </a>
    <div id="score-id" class="primary-halo-font">{{ scoreId }}</div>
    <div>
      <table id="scoreboard-teams" class="secondary-halo-font">
        <tr v-for="(team, index) in data.teams" :key=index>
          <td class="team-name">{{ team.name }}</td>
          <td class="team-score">{{ team.score }}</td>
          <td><span v-if="data.completed && is_winner(index)">✔</span></td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
import '../styling/fonts.css';

export default {
  name: 'ScoreboardItem',
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
    is_winner(i) {
      let winner = {
        score: Number.MIN_VALUE,
        index: -1
      };
      this.data.teams.forEach((team, index) => {
        if(team.score > winner.score) {
          winner.score = team.score;
          winner.index = index;
        }
      });
      return i === winner.index;
    }
  },

  created() {
    console.log("Loading ScoreboardItem");
  },

  computed: {
    edit_link() {
      return `${window.location.origin}/${this.data['score-id']}/edit`;
    }
  }
}
</script>

<style>
  #edit-button {
    position: absolute;
    top: 0px;
    right: 0px;
    padding: 1px;
  }

  #edit-button img {
    width: calc(0.75vh + 0.75vw);
    height: calc(0.75vh + 0.75vw);
  }

  #edit-button img:hover {
    background-color: #4a5363;
    border-radius: 2px;
  }

  #featured-marker {
    position: absolute;
    top: 0px;
    left: 0px;
  }

  #score-id {
    font-size: calc(1vh + 1vw);
  }

  #scoreboard-teams {
    width: 100%;
  }
                
 .scoreboard-listed {
      position: relative;
      margin: 10px 2px;
      width: 25%;
      height: fit-content;
      overflow-x: auto;
      background-color: #172030;
      color: white;
      font-size: calc(0.75vh + 0.75vw);
      filter: drop-shadow(1px 1px 5px gray);
  }

  .team-name {
    width: 70%;
    text-align: left;
  }

  .team-score {
    width: 15%;
  }
</style>
