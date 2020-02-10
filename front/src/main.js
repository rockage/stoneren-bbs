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

Vue.use(ElementUI)
Vue.use(VueQuillEditor)

Vue.prototype.md5 = md5
Vue.prototype.axios = axios
axios.defaults.withCredentials = true

Vue.config.productionTip = false
Vue.config.devtools = false


// 自定义区：
import router from './router'
import bbsHeader from './components/bbsHeader.vue'
import * as Popup from './popup' //自定义全局弹窗组件
import base from './base'//全局函数
import store from './store' //vuex响应式全局变量


import VueCropper from 'vue-cropper'

Vue.use(VueCropper)

let globalVariable = { //vue普通的全局变量
  root: '', //模拟根实例$root
  fid: 0,
  forumsData: []
}
Vue.prototype.GLOBAL = globalVariable
Vue.use(base)
Vue.component('bbsHeader', bbsHeader) //自定义组件: <bbsHeader> </bbsHeader>
Vue.prototype.$login = Popup.LoginBox
Vue.prototype.$password = Popup.PasswordBox
Vue.prototype.$userprofile = Popup.UserProfileBox
Vue.prototype.$post = Popup.PostBox
Vue.prototype.$profile = Popup.ProFileBox


new Vue({
  el: "#app",
  store: store,
  router,
  components: {
    App,
  },
  template: '<App/>',

  methods: {},
  mounted: function () {
    this.GLOBAL.root = this.$root //为所有动态实例提供根实例的访问
    this.getForumsData()
    console.log(this.GLOBAL.forumsData)
  },
})



