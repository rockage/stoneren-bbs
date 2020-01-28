// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

import VueQuillEditor from 'vue-quill-editor'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'
import VueCookie from 'vue-cookie'

import Post from './components/post.vue'
import Login from "./components/login.vue";

import globalVariable from './global_variable'

Vue.prototype.GLOBAL = globalVariable

import base from './base'//引用
Vue.use(base);//将全局函数当做插件来进行注册

import Vuex from 'vuex'
Vue.use(Vuex)


Vue.prototype.axios = axios
Vue.use(ElementUI);
Vue.use(VueQuillEditor)
Vue.use(VueCookie)

Vue.component('post', Post) //自定义组件，<post> </post>

Vue.config.productionTip = false
Vue.config.devtools = false

const store = new Vuex.Store({
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

new Vue({
  el: "#app",
  store,
  router,
  components: {
    App,
  },
  template: '<App/>',

  methods: {},
  mounted: function () {
    const vm = this
    this.loginCheck(function (r) {
      console.log('store:'+store.state.loginState)

    })
  },
})

Vue.prototype.$login = function (data) { //自定义全局弹窗组件，也可以把这一段放到一个独立的js文件中
  const LoginBox = Vue.extend(Login)
  let instance = new LoginBox({data})
  let LoginEl = instance.$mount().$el
  document.body.appendChild(LoginEl)
  Vue.nextTick(() => {
    instance.dialogVisible = true
  })
}
