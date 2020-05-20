<template>
  <div id="app">
    <Toolbox-panel></Toolbox-panel>
    <Map></Map>
  </div>
</template>

<script>
import Map from './components/Map'
import ToolboxPanel from './components/ToolboxPanel'
import { mapActions } from 'vuex'

export default {
  components: { Map, ToolboxPanel },
  methods: {
    ...mapActions(['loadPoints'])
  },
  mounted () {
    this.loadPoints()
  }
}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  position: relative;
}
</style>
