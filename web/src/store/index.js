import Vue from 'vue'
import Vuex from 'vuex'
import API from '../api'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    points: []
  },
  mutations: {
    setPoints: (state, points) => {
      state.points = points
    }
  },
  actions: {
    loadPoints ({ commit }) {
      API.points.list()
        .then(resp => {
          if (resp.status === 200) {
            commit('setPoints', resp.data)
          }
          commit('setPoints', [])
        })
    },
    createPoint ({ commit, dispatch }, point) {
      API.points.create(point)
        .then(data => {
          if (data.status === 200) {
            dispatch('loadPoints')
          }
          return new Promise(resolve => resolve(data))
        })
    }
  },
  modules: {}
})
