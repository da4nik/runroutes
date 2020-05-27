<template>
  <div id="map">
    <h2>Map</h2>
    <yandex-map
      :coords="[52.638634, 41.443438]"
      ymap-class='maps-container'
      v-on:click='mapClicked'>
<!--      :controls="['geolocationControl', 'rulerControl', 'typeSelector']">-->

      <ymap-marker
        v-for='point in points'
        :key='point.id'
        :marker-id='point.id'
        marker-type='Placemark'
        :coords='[point.lat, point.long]'
        v-on:click='markerClicked(point)'
      />

    </yandex-map>
  </div>
</template>

<script>
import { yandexMap, loadYmap, ymapMarker } from 'vue-yandex-maps'
import { mapState } from 'vuex'

export default {
  name: 'Map',
  components: { yandexMap, ymapMarker },

  computed: {
    ...mapState(['points'])
  },

  methods: {
    mapClicked: function (e) {
      this.$eventHub.$emit('map-clicked', e.get('coords'))
    },
    markerClicked: function (point) {
      this.$eventHub.$emit('marker-clicked', point)
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
