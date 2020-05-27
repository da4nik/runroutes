import Vue from 'vue'
import Vuex from 'vuex'
import API from '../api'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    points: [],
    ways: []
  },
  mutations: {
    setPoints: (state, points) => {
      state.points = points
    },
    setWays: (state, ways) => {
      state.ways = ways
    }
  },
  actions: {
    loadPoints ({ commit }) {
      API.points.list()
        .then(resp => {
          if (resp.status === 200) {
            commit('setPoints', resp.data)
            return
          }
          commit('setPoints', [])
        })
    },
    createPoint ({ dispatch }, point) {
      API.points.create(point)
        .then(data => {
          if (data.status === 201) {
            dispatch('loadPoints')
          }
        })
    },
    loadWays ({ commit }) {
      API.ways.list()
        .then(resp => {
          if (resp.status === 200) {
            commit('setWays', resp.data)
            return
          }
          commit('setPoints', [])
        })
    },
    createWay ({ dispatch }, way) {
      API.ways.create(way)
        .then(data => {
          if (data.status === 201) {
            dispatch('loadWays')
          }
        })
    }
  },
  modules: {}
})
