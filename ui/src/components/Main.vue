<template>
  <div>
      <h1>Scoreboard</h1>
      <div v-for="[id, score] in scores" :key=id>
          {{ id }}: {{ score }}
      </div>
  </div>
</template>

<script>
import ScoreboardAPI from '../js/api';
import ScoreListener from '../js/listener';

export default {
  name: 'Main',
  components: {
    
  },
  data() {
    return {
      listener: new ScoreListener() ,
      handler: new ScoreboardAPI(),
      scoreboards: {}
    }
  },
  created() {
    this.handler.list_scoreboards()
        .then((scoreboards) => {
            console.log(scoreboards)
            this.scoreboards = scoreboards;
        }).catch((err) => {
            console.error(err);
        });

    this.listener.on('message', (event) => {
      const score = JSON.parse(event.data);
      this.scoreboards[score['score-id']] = score.scoreboard;
    });
    this.listener.connect();
  },
  computed: {
      scores() {
          return Object.entries(this.scoreboards);
      }
  }
}
</script>

<style>

</style>
