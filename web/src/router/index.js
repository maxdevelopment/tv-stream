import Vue from 'vue'
import Router from 'vue-router'
import CreateStream from '~/pages/CreateStream'
import Stream from '~/pages/Stream'
import NotFoundComponent from '~/pages/NotFoundComponent'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'CreateStream',
      component: CreateStream
    },
    {
      path: '/stream',
      name: 'Stream',
      component: Stream
    },
    {
      path: '*',
      component: NotFoundComponent
    }
  ]
})
