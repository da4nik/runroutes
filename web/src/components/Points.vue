<template>
  <div id="points">
    <h2>Points</h2>
    <div id="point-form" v-if="point">
      <p>
        lat: {{ point.lat }}
      </p>
      <p>
        long: {{ point.long }}
      </p>
      <p>
        <label>Asphalt %
          <input type="number" v-model.number="point.asphalt_percent" min="0" max="100"/>
        </label>
      </p>
      <p>
        <label>Preferred %
          <input type="number" v-model.number="point.preferred" min="0" max="100"/>
        </label>
      </p>
      <button v-on:click="create">Create</button>
    </div>

  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
  name: 'Points',
  data () {
    return {
      point: null
    }
  },
  computed: {
    ...mapState(['points'])
  },
  methods: {
    ...mapActions(['createPoint']),
    onMapClicked: function (coords) {
      this.point = {
        lat: '' + coords[0],
        long: '' + coords[1],
        asphalt_percent: 100,
        preferred: 100
      }
    },
    create () {
      this.createPoint(this.point)
    }
  },
  created: function () {
    this.$eventHub.$on('map-clicked', this.onMapClicked)
  },
  beforeDestroy: function () {
    this.$eventHub.$off('map-clicked')
  }
}
</script>

<style lang="scss" scoped>
#point-form {
  text-align: left;
}
</style>
