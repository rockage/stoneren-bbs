// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
// 引用区：
import Vue from 'vue'
import App from './App'
import axios from 'axios'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import md5 from 'js-md5'
import VueQuillEditor from 'vue-quill-editor'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'
import VueCookies from 'vue-cookies'


Vue.use(VueCookies)
Vue.use(ElementUI);
Vue.use(VueQuillEditor)
Vue.prototype.md5 = md5
Vue.prototype.axios = axios
axios.defaults.withCredentials=true

Vue.config.productionTip = false
Vue.config.devtools = false

//axios.defaults.withCredentials = true

// 自定义区：
import router from './router'
import bbsHeader from './components/bbsHeader.vue'
import Post from './components/post.vue'
import * as Popup from './popup' //自定义全局弹窗组件
import base from './base'//全局函数
import store from './store' //vuex全局变量


import VueCropper from 'vue-cropper'

Vue.use(VueCropper)

let globalVariable = { //vue传统全局变量
  globalThis: '',
}
Vue.prototype.GLOBAL = globalVariable
Vue.use(base)
Vue.component('post', Post) //自定义组件: <post> </post>
Vue.component('bbsHeader', bbsHeader) //自定义组件: <bbsHeader> </bbsHeader>
Vue.prototype.$login = Popup.LoginBox
Vue.prototype.$password = Popup.PasswordBox
Vue.prototype.$profile = Popup.ProfileBox
Vue.prototype.$userprofile = Popup.UserProfileBox

new Vue({
  el: "#app",
  store,
  router,
  components: {
    App,
    Post,
  },
  template: '<App/>',

  methods: {
    autoLogin: function () {
      //   const auto = this.getCookie('autologin')
      //  const uname = this.getCookie('username')
      //  const pass = this.getCookie('password')

      const auto = this.getCookie('autologin')
      const uname = this.getCookie('username')
      const pass = this.getCookie('password')
      console.log('auto = ' + auto)
      console.log('uname = ' + uname)
      console.log('pass = ' + pass)

      if (auto !== 'true' || uname === null || pass === null) {
        store.commit('setLoginState', false)
        return
      }

      this.loginCheck(function (data) { //先判断session能否登录
        if (data !== "") {
          store.commit('setLoginState', true)
          store.commit('setUid', data.uid)
          store.commit('setUname', data.username)
          console.log('in session login:' + data.uid)
          console.log('in session login:' + data.username)
        } else {
          axios.get('http://localhost:8081/login', { //再使用cookie登录
            params: {
              withCredentials: true,
              username: uname,
              password: pass,
            }
          })
            .then((response) => {
              if (response.data !== 'not found') {
                const ret = JSON.parse(response.data)
                store.commit('setLoginState', true)
                store.commit('setUid', ret[0].uid)
                store.commit('setUname', uname)
                console.log('in coockie login')
              }
            })
        }
      })
    }
  },
  mounted: function () {
    this.autoLogin()
    this.GLOBAL.globalThis = this
  },
})



