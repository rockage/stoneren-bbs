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
      component: bbs,
      meta: {rmode: 'new'},
    },
    {
      path: '/new/:page',
      name: 'new',
      component: bbs,
      meta: {rmode: 'new'},
    },
    {
      path: '/forums/view/:fid/:page',
      name: 'forumsview',
      component: bbs,
      meta: {rmode: 'normal'},
    },
    {
      path: '/threads/userThreads/:uid/:page',
      name: 'userThreads',
      component: bbs,
      meta: {rmode: 'self'},
    },
    {
      path: '/threads/view/:tid',
      name: 'threadsview',
      component: threadsview
    },
  ]
})
