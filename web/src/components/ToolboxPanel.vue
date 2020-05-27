<template>
  <div id="toolbox-panel">
    <nav>
      <div class="link">
        <router-link to="/points">Points</router-link>
      </div>

      <div class="link">
        <router-link to="/ways">Ways</router-link>
      </div>
    </nav>

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

  nav {
    width: 100%;

    .link {
      display: inline-block;
      width: 50%;
      text-align: center;

      a {
        text-decoration: none;

        &.router-link-active {
          font-weight: bold;
        }

        &, &:visited {
          color: cornflowerblue;
        }
      }
    }
  }
}
</style>
