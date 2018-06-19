import Routes from '~/static/routes'
import router from '~/router'

export default {
  state: {
    id: null,
    createdAt: null,
    ws: null,
    stream: null
  },

  getters: {
    stream: state => state,
    getStream: state => state.stream
  },

  mutations: {
    CREATE_STREAM(state, data) {
      state.id = data.id
      state.createdAt = data.created_at
    },
    SET_WS(state, data) {
      state.ws = data.ws
    },
    SET_STREAM(state, data) {
      state.stream = data
    }
  },

  actions: {
    startStream({commit}) {
      Routes.startStream().get().then(response => {
        commit('CREATE_STREAM', response.body)
        router.push({path: `/stream`})
      }, error => {
        console.log(error)
      })
    },
    setWs({commit}, data) {
      commit('SET_WS', data)
    },
    setStream({commit}, stream) {
      commit('SET_STREAM', stream)
    }
  }
}
