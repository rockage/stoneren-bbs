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
      meta: {renderMode: 'new'},
    },
    {
      path: '/forums/view/:fid',
      name: 'forumsview',
      component: bbs,
      meta: {renderMode: 'normal'},
    },
    {
      path: '/threads/userThreads/:uid',
      name: 'userThreads',
      component: bbs,
      meta: {renderMode: 'self'},
    },
    {
      path: '/threads/view/:tid',
      name: 'threadsview',
      component: threadsview
    },
  ]
})
