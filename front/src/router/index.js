//import Vue from 'vue'
//import Router from 'vue-router'
import bbs from '../components/bbs'
import threadsview from '../components/threadsview'
import usersview from '../components/usersview'

export default new VueRouter({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: bbs,
      meta: {rmode: 'new'},
    },
    {
      path: '/new/:sortmode/:page',
      name: 'new',
      component: bbs,
      meta: {rmode: 'new'},
    },
    {
      path: '/forums/view/:sortmode/:fid/:page',
      name: 'forumsview',
      component: bbs,
      meta: {rmode: 'normal'},
    },
    {
      path: '/threads/userThreads/:sortmode/:uid/:page',
      name: 'userThreads',
      component: bbs,
      meta: {rmode: 'self'},
    },
    {
      path: '/threads/view/:tid',
      name: 'threadsview',
      component: threadsview
    },
    {
      path: '/users/:sortmode/:page',
      name: 'usersview',
      component: usersview
    },
  ]
})
