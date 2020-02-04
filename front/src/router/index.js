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
      meta: {renderMode: 0},
    },
    {
      path: '/forums/view/:fid',
      name: 'forumsview',
      component: bbs,
      meta: {renderMode: 1},
    },
    {
      path: '/threads/userThreads/:uid',
      name: 'userThreads',
      component: bbs,
      meta: {renderMode: 2},
    },
    {
      path: '/threads/view/:tid',
      name: 'threadsview',
      component: threadsview
    },
  ]
})
