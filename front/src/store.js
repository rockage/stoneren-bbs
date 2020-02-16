import Vuex from 'vuex'
import Vue from 'vue'
import axios from 'axios'
import createPersistedState from "vuex-persistedstate"

Vue.use(Vuex)

const store = new Vuex.Store({
  plugins: [createPersistedState({
    storage: window.sessionStorage //store持久化
  })],
  state: {
    loginState: false,
    uid: 0,
    uname: '',
    passwd: '',
    fid: 0,
    fname: '',  //当前论坛名称
    fsname: '', //全部论坛名称
    rmode: '', //bbs渲染模式
  },
  getters: {
    loginState: function (state) {
      return state.loginState
    },
    uid: function (state) {
      return state.uid
    },
    uname: function (state) {
      return state.uname
    },
    passwd: function (state) {
      return state.passwd
    },
    fid: function (state) {
      return state.fid
    },
    fname: function (state) {
      let fs
      switch (state.rmode) {
        case "new":
          fs = "最新的帖子"
          break;
        case "normal":
          for (let i = 0; i < state.fsname.length; i++) {
            if (state.fid === state.fsname[i].fid) {
              fs = state.fsname[i].name
            }
          }
          break;
        case "self":
          fs = state.fname + "的主题"
          break;
      }
      return fs
    },
    fsname: function (state) {
      return state.fsname
    },
    rmode: function (state) {
      return state.rmode
    },
  },
  mutations: {
    loginState: function (state, r) {
      state.loginState = r
    },
    uid: function (state, r) {
      state.uid = r
    },
    uname: function (state, r) {
      state.uname = r
    },
    passwd: function (state, r) {
      state.passwd = r
    },
    fid: function (state, r) {
      state.fid = r
    },
    fname: function (state, r) {
      state.fname = r
    },
    fsname: function (state, r) {
      state.fsname = r
    },
    rmode: function (state, r) {
      state.rmode = r
    },
  },
  actions: {
    setFsname: ({ commit }) => { //异步mutation必须使用actions
      return new Promise(function (resolve) {
        axios.get('http://localhost:8081/getForums', {})
          .then((response) => {
            commit('fsname', JSON.parse(response.data))
            resolve('asynced')
          })
      })
    },

  }
})

export default store
