import Vuex from 'vuex'
import Vue from 'vue'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    loginState: false,
    uid: 0,
    uname: '',
    passwd:'',
  },
  mutations: {
    setLoginState(state, r) {
      state.loginState = r
    },
    setUid(state, r) {
      state.uid = r
    },
    setUname(state, r) {
      state.uname = r
    },
    setPasswd(state, r) {
      state.passwd = r
    },
  }
})

export default store
