import Vue from 'vue'
import Router from 'vue-router'
import bbs from '../components/bbs'
import threadsview from '../components/threadsview'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    // {path: '/', name: 'HelloWorld', component: HelloWorld},
    {
      path: '/',
      name: 'bbs',
      query: {
        fid: 0
      },
      component: bbs
    },
    {
      path: '/threads/view/:tid',
      component: threadsview
    },
    {
      path: '/forums/view/:fid',
      component: bbs
    }
  ]
})
