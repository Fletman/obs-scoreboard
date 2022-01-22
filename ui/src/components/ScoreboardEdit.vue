<template>
  <div>
    <header>
      <h1 class="primary-halo-font title text-glow">{{ score_data['score-id'] }}</h1>
    </header>
    <div>
      <form v-on:change="update_score" v-on:submit.prevent="">
        <table id="edit-menu" class="secondary-halo-font">
          <tr class="score-checkbox-row">
            <td>
              <button style="margin-right:10px;" class="score-button-input secondary-halo-font" v-on:click="return_home">Back </button>
              <button class="score-button-input secondary-halo-font" v-on:click="delete_scoreboard">Delete</button>
            </td>
          </tr>
          <tr class="score-checkbox-row">
            <td colspan="3">
              <input class="score-checkbox-input" type="checkbox" id="completed" v-model="score_data.completed"/>
              <label for="completed">Match Completed</label>
            </td>
          </tr>
          <tr class="score-checkbox-row">
            <td colspan="3">
              <input class="score-checkbox-input" type="checkbox" id="featured" v-model="score_data.featured"/>
              <label for="featured">Featured Match</label>
            </td>
          </tr>
          <tr v-for="(team, index) in score_data.teams" :key="index">
            <td><input style="width:100%;" class="score-text-input secondary-halo-font" type="text" :readonly="score_data.completed" v-model="team.name" required/></td>
            <td>
              <span class="score-button-input" v-on:click="score_increment(index, -1)">−</span>
              <input class="score-text-input secondary-halo-font" type="number" :readonly="score_data.completed" v-model="team.score"/>
              <span class="score-button-input" :disabled="score_data.completed" v-on:click="score_increment(index, 1)">+</span>
            </td>
            <td style="text-align:left;">
              <span class="score-button-input" :disabled="score_data.completed" v-on:click="remove_team(index)">X</span>
            </td>
          </tr>
          <tr>
            <td colspan="3">
              <button class="score-button-input secondary-halo-font" v-on:click="add_team">Add Team</button>
            </td>
          </tr>
        </table>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ScoreboardEdit',
  components: {
    
  },

  props: {
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
        score_data: this.data
    }
  },

  methods: {
    prevent_change() {
      window.alert("Changes cannot be made to a scoreboard marked as completed");
    },

    add_team() {
      if(!this.score_data.completed) {
        this.score_data.teams.push({
          name: "",
          score: 0
        });
      } else {
        this.prevent_change();
      }
    },

    remove_team(index) {
      if(!this.score_data.completed) {
        this.score_data.teams.splice(index, 1);
        this.update_score();
      } else {
        this.prevent_change();
      }
    },

    score_increment(index, val) {
      if(!this.score_data.completed) {
        this.score_data.teams[index].score += val;
        this.update_score();
      } else {
        this.prevent_change();
      }
    },

    update_score() {
      this.handler.set_scoreboard(this.score_data['score-id'], this.score_data);
    },

    delete_scoreboard() {
      if(window.confirm(`Are you sure you want to delete ${this.score_data['score-id']}?`)) {
        this.handler.remove_scoreboard(this.score_data['score-id'])
          .then(() => {
            this.$router.replace('/');
          }).catch((err) => {
            console.error(err);
          });
      }
    },

    return_home() {
      this.$router.push('/');
    }
  },

  created() {
      console.log("Loading ScoreboardEdit");
  },

  computed: {
      
  },

  watch: {
    data(new_score) {
      this.score_data = new_score;
    }
  }
}
</script>

<style>
  @import '../css/main.css';
  @import '../css/fonts.css';

  #edit-menu {
    margin: auto;
  }

  #edit-menu td {
    width: 40%;
    padding: 10px 10px;
  }

  .score-button-input {
    padding: 5px;
    font-size: 1.25em;
    background-color: #172030;
    color: white;
    border: solid 1px gray;
    -webkit-touch-callout: none;
    -webkit-user-select: none;
    -khtml-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
  }

  .score-button-input:hover {
    cursor: pointer;
    background: #394252;
  }
  
  .score-checkbox-row {
    text-align: left;
  }

  .score-checkbox-input {
    padding: 5px;
  }

  .score-text-input {
    padding: 5px;
    font-size: 1.25em;
    background-color: #172030;
    color: white;
    border: solid 1px gray;
    text-align: center;
    -moz-appearance: textfield;
  }
  .score-text-input::-webkit-outer-spin-button,
  .score-text-input::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
  }
</style>
