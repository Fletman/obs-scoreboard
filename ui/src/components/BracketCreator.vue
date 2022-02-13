<template>
  <div>
    <form v-on:submit.prevent="create_bracket">
      <input type="text" placeholder="Enter bracket name" v-model="bracket_name"/>
      <div v-for="(seed, index) in seeds" :key="index">
        <span>[{{ index+1 }}]</span>
        <input type="text" placeholder="Enter name" v-model="seed.name"/>
        <input type="button" value="X" v-on:click="remove_seed(index)"/>
      </div>
      <input type="button" value="Add Seed" v-on:click="add_seed"/>
    </form>
  </div>
</template>

<script>
export default {
  name: 'BracketEditor',

  props: {
    visible: {
      type: Boolean,
      required: true,
      default: false
    }
  },

  data() {
    return {
      bracket_name: "",

      seeds: []
    }
  },

  methods: {
    reset_fields() {
      this.seeds = []; 
    },

    add_seed() {
      this.seeds.push({name: ""});
    },

    remove_seed(index) {
      this.seeds.splice(index, 1);
    },

    create_bracket() {
      console.log(this.bracket_name);
      console.log(this.seeds);
    }
  }
}
</script>

<style>
  @import '../css/main.css';
  @import '../css/fonts.css';
</style>
