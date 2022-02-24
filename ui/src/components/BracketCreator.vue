<template>
  <div>
    <form v-on:submit.prevent="">
      <input id="bracket-name-input" class="secondary-halo-font" type="text" placeholder="Enter bracket name" required v-model="bracket_name"/>
      <div class="bracket-seed-input" v-for="(seed, index) in seeds" :key="index">
        <span>[{{ index+1 }}]</span>
        <input class="bracket-team-input secondary-halo-font" type="text" placeholder="Enter name" v-model="seed.name"/>
        <input type="button" value="X" v-on:click="remove_seed(index)"/>
      </div>
      <input type="button" value="Add Seed" v-on:click="add_seed"/>
      <input v-if="seeds.length > 0" type="button" value="Submit" v-on:click="create_bracket"/>
    </form>
  </div>
</template>

<script>
import ScoreboardAPI from '../js/api';
export default {
  name: 'BracketEditor',

  props: {
    
  },

  data() {
    return {
      handler: new ScoreboardAPI(),

      bracket_name: "",

      seeds: []
    }
  },

  methods: {
    reset_fields() {
      this.bracket_name = "";
      this.seeds = []; 
    },

    add_seed() {
      this.seeds.push({name: ""});
    },

    remove_seed(index) {
      this.seeds.splice(index, 1);
    },

    create_bracket() {
      this.handler.create_bracket(this.bracket_name, this.seeds.map((seed) => seed.name))
        .then((bracket) => {
          this.reset_fields();
          this.$emit('bracket-created', bracket);
        }).catch((err) => console.error(err));
    }
  }
}
</script>

<style>
  @import '../css/main.css';
  @import '../css/fonts.css';

  #bracket-name-input {
    display: block;
    font-size: 1em;
    opacity: 0.75;
  }

  .bracket-seed-input {
    display: grid;
    grid-template-columns: 5% 90% 5%;
  }

  .bracket-team-input {
    font-size: 1em;
    opacity: 0.75;
  }
</style>
