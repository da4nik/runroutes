<template>
  <div id="map">
    <h2>Map</h2>
    <yandex-map
      :coords="[52.638634, 41.443438]"
      ymap-class='maps-container'
      v-on:click='mapClicked'
      :controls="['geolocationControl', 'rulerControl', 'typeSelector']"
    />
  </div>
</template>

<script>
import { yandexMap, loadYmap } from 'vue-yandex-maps'

export default {
  name: 'Map',
  components: { yandexMap },

  methods: {
    mapClicked: function (e) {
      this.$eventHub.$emit('map-clicked', e.get('coords'))
    }
  },

  async mounted () {
    await loadYmap({ debug: true })
  }
}
</script>

<style lang="scss">
#map {
  .maps-container {
    position: fixed;
    top: 0; bottom: 0;
    left: 0; right: 0;
    width: 100%;
    height: 100%;
    z-index: 1;
  }
}
</style>
