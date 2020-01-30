import Vuex from 'vuex'

export const store = new Vuex.Store({
  state: {
    count: 1,
    loginState: false,
  },
  mutations: {
    increment(state) {
      state.count++
    },
    loginStateChanged(state, r) {
      state.loginState = r
    }
  }
})
