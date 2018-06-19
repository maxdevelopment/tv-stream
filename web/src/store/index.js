import Vuex from 'vuex'
import Stream from '~/store/stream'

export default () => new Vuex.Store({
  modules: {
    Stream
  }
})
