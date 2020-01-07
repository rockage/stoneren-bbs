import Vue from 'vue'
import Router from 'vue-router'
import bbs from '../components/bbs'
import threadsview from '../components/threadsview'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'bbs',
      component: bbs,
      meta: {
        keepAlive: false,
      },
    },
    {
      path: '/threads/view',
      name: 'threadsview',
      component: threadsview
    },
    {
      path: '/forums/view/:fid',
      name:'forumsview',
      component: bbs,
      meta: {
        keepAlive: false,
      },
    }
  ]
})
