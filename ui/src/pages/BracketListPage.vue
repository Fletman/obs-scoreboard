<template>
  <div id="main">
    <div id="main-view">
      <NavComponent/>
      <header class="primary-halo-font">
        <h1 class="title text-glow">HALO</h1>
        <h3 class="title text-glow">Tournament of the Chosen</h3>
        <h6 class="title">Sponsored by Flet Inc.™</h6>
      </header>
      <div class="bracket-menu-row">
        <input type="button" class="halo-button-input secondary-halo-font" value="Create New Bracket" v-on:click="toggle_creator"/>
      </div>
      <div class="bracket-menu-row">
        <BracketCreator :visible="show_creator"/>
      </div>
      <div style="grid-template-columns:90% 10%" class="bracket-menu-row">
        <input type="text" id="bracket-search-input" class="secondary-halo-font" placeholder="Search for a bracket" v-model="search_id"/>
        <input type="button" class="halo-button-input secondary-halo-font" value="↻" v-on:click="load_brackets"/>
      </div>
      <div id="bracket-list" class="bracket-menu-row">
        <router-link class="link-style" v-for="id in search_ids" :key="id" :to="`/brackets/${id}/edit`">
          <div class="bracket-item secondary-halo-font">
            {{ id }}
          </div>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import ScoreboardAPI from '../js/api';
import NavComponent from '../components/NavComponent.vue';
import BracketCreator from '../components/BracketCreator.vue';

export default {
  name: 'BracketListPage',
  components: {
    NavComponent,
    BracketCreator
  },

  data() {
    return {
      handler: new ScoreboardAPI(),

      search_id: "",

      bracket_ids: [],

      show_creator: false
    }
  },

  methods: {
    toggle_creator() {
      this.show_creator = !this.show_creator;
    },

    load_brackets() {
      this.handler.list_brackets()
        .then((bracket_ids) => this.bracket_ids = bracket_ids)
        .catch((err) => console.error(err));
    }
  },

  beforeMount() {
    this.load_brackets();
  },

  computed: {
    search_ids() {
      const s_id = this.search_id.toLowerCase();
      return this.search_id === "" ?
        this.bracket_ids :
        this.bracket_ids.filter((bracket_id) => bracket_id.toLowerCase().indexOf(s_id) > -1);
    }
  }
}
</script>

<style>
  @import '../css/main.css';
  @import '../css/fonts.css';

  #bracket-list {
    max-height: 50%;
    overflow: auto;
    background-color: rgba(23, 32, 48, 0.75);
  }

  #bracket-search-input {
    padding: 5px;
    font-size: 1.25em;
    width: auto;
  }

  .bracket-item {
    width: auto;
    padding: 10px;
    border: solid rgba(128, 128, 128, 0.5);
    border-width: 1px 0px;
    font-size: 1.125em;
  }

  .bracket-item:hover {
    background-color: rgba(290, 116, 61, 0.5);
    cursor: pointer;
  }

  .bracket-menu-row {
    display: grid;
    margin: auto;
    width: 50%;
  }
</style>