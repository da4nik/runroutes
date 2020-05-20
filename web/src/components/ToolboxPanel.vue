<template>
  <div id="toolbox-panel">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'ToolboxPanel'
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

  padding: 10px;

  background-color: azure;
  opacity: 0.8;

  box-shadow: 0 0 10px rgba(0,0,0,0.2);
}
</style>
