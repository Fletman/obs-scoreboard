<template>
  <div id="main">
    <div id="main-view">
      <NavComponent/>
      <header class="primary-halo-font">
        <h1 class="title text-glow">HALO</h1>
        <h3 class="title text-glow">Tournament of the Chosen</h3>
        <h6 class="title">Sponsored by Flet Inc.™</h6>
      </header>
      <div id="bracket-list" class="secondary-halo-font">
        <div style="margin-bottom:5%; grid-template-columns:10% 75% 15%" class="bracket-list-row">
          <input type="button" class="halo-button-input" value="↻" v-on:click="load_brackets"/>
          <input type="text" id="bracket-search-input" class="secondary-halo-font" placeholder="Search for a bracket" v-model="search_id"/>
          <input type="button" class="halo-button-input secondary-halo-font" value="New"/>
        </div>
        <div id="bracket-list-scrollable">
          <div style="grid-template-columns:85% 15%" class="bracket-list-row" v-for="id in search_ids" :key="id">
            <router-link class="bracket-list-item bracket-list-link" :to="`/brackets/${id}/view`">
              {{ id }}
            </router-link>
            <router-link class="bracket-list-item bracket-list-link" :to="`/brackets/${id}/edit`">
              Edit
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ScoreboardAPI from '../js/api';
import NavComponent from '../components/NavComponent.vue';

export default {
  name: 'BracketListPage',
  components: {
    NavComponent
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
    margin: auto;
    margin-top: 10px;
    width: 50%;
  }

  #bracket-list-scrollable {
    max-height: 50vh;
    overflow: auto;
    background-color: rgba(23, 32, 48, 0.75);
  }

  #bracket-search-input {
    padding: 5px;
    font-size: 1.25em;
    width: auto;
  }

  .bracket-list-row {
    display: grid;
    margin: auto;
    width: 100%;
  }

  .bracket-list-item {
    padding: 10px;
    border: solid rgba(128, 128, 128, 0.5);
    border-width: 1px 0px;
    font-size: 1.125em;
  }

  .bracket-list-link {
    color: white;
    text-decoration: none;
    text-align: center;
    border: solid rgba(128, 128, 128, 0.5);
    border-width: 1px 1px;
  }

  .bracket-list-link:hover {
    background-color: rgba(290, 116, 61, 0.5);
    cursor: pointer;
  }
</style>