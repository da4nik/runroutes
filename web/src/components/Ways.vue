<template>
  <div id='ways'>
    <h2>Ways</h2>
    <div v-if='from' class='point' :class='{ active: isFrom }'>
      <p class='title'>From point:</p>
      <p>
        lat: {{ from.lat }}
      </p>
      <p>
        long: {{ from.long }}
      </p>
    </div>
    <div v-if='to' class='point' :class='{ active: isTo }'>
      <p class='title'>To point:</p>
      <p>
        lat: {{ to.lat }}
      </p>
      <p>
        long: {{ to.long }}
      </p>
    </div>
    <div>
      Asphalt
      <input type="number" v-model.number='asphalt'>
    </div>
    <div>
      Preferred
      <input type="number" v-model.number='preferred'>
    </div>
    <button v-on:click='create'>Create</button>
  </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  name: 'Ways',
  data () {
    return {
      from: null,
      to: null,
      asphalt: 100,
      preferred: 100,
      current: 'from'
    }
  },
  computed: {
    isFrom () {
      return this.current === 'from'
    },
    isTo () {
      return this.current === 'to'
    }
  },
  methods: {
    ...mapActions(['createWay']),
    onMarkerClicked: function (point) {
      if (this.current === 'from') {
        this.current = 'to'
        this.from = point
        return
      }

      this.current = 'from'
      this.to = point
    },
    create: function () {
      const { from, to, asphalt, preferred } = this
      this.createWay({
        from,
        to,
        distance: 1,
        asphalt_percent: asphalt,
        preferred: preferred
      })
    }
  },
  created: function () {
    this.$eventHub.$on('marker-clicked', this.onMarkerClicked)
  },
  beforeDestroy: function () {
    this.$eventHub.$off('marker-clicked')
  }
}
</script>

<style lang="scss" scoped>
  #ways {
    text-align: left;

    .point {
      &.active {
        .title {
          font-weight: bold;
        }
      }
    }
  }
</style>
