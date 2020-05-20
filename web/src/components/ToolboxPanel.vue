<template>
  <div id="toolbox-panel">
    <h2>Toolbox panel</h2>
    {{ point }}
    <router-view></router-view>
  </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  name: 'ToolboxPanel',
  data: () => ({
    point: null
  }),
  methods: {
    ...mapActions(['loadPoints']),
    onMapClicked: function (coords) {
      this.point = coords
    }
  },
  created () {
    this.$eventHub.$on('map-clicked', this.onMapClicked)
  },
  beforeDestroy () {
    this.$eventHub.$off('map-clicked')
  },
  mounted () {
    this.loadPoints()
  }
}
</script>

<style lang="scss">
#toolbox-panel {
  position: fixed;
  top: 20%;
  bottom: 20%;
  right: 10px;
  width: 300px;
  z-index: 2;

  background-color: azure;
  opacity: 0.8;

  box-shadow: 0 0 10px rgba(0,0,0,0.2);
}
</style>
