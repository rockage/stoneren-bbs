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




import post from './components/post.vue'



Vue.prototype.axios = axios
Vue.use(ElementUI);
Vue.use(VueQuillEditor)
Vue.component('post', post) //自定义模块:post

Vue.config.productionTip = false

new Vue({
  el: "#app",
  router,
  components: {
    App,
  },
  template: '<App/>',

  methods: {

    xxx:function () {
      console.log("main ok")

    }

  }
})

Vue.prototype.xxx = function (){//changeData是函数名
  alert('执行成功');
}
