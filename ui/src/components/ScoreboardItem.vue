<template>
  <div class="scoreboard-listed">
    <div>
      <span id="featured-marker" v-if="data.featured">💫</span>
      <router-link id="edit-button" :to="edit_link">
        <img src="../assets/edit_icon.webp" alt="Edit Score"/>
      </router-link>
    </div>
    <router-link id="score-id" class="primary-halo-font" :to="view_link">
      {{ scoreId }}
    </router-link>
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
        score: Number.MIN_SAFE_INTEGER,
        index: -1
      };
      this.data.teams.forEach((team, index) => {
        if(team.score > winner.score) {
          winner.score = team.score;
          winner.index = index;
        }
      });
      return i === winner.index;
    },

    open_editor() {
      this.$router.push(this.edit_link)
    }
  },

  computed: {
    view_link() {
      return `/scores/${this.data['score-id']}/view`;
    },

    edit_link() {
      return `/scores/${this.data['score-id']}/edit`;
    }
  }
}
</script>

<style>
  @import '../css/fonts.css';

  #edit-button {
    padding: 1px;
    clear: right;
    float: right;
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
    float: left;
  }

  #score-id {
    font-size: calc(1vh + 1vw);
    color: white;
    text-decoration: none;
  }

  #score-id:hover {
    text-decoration: underline;
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
