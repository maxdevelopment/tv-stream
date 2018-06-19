import Routes from '~/static/routes'
import router from '~/router'

export default {
  state: {
    id: null,
    createdAt: null,
    ws: null
  },

  getters: {
    stream: state => state
  },

  mutations: {
    CREATE_STREAM(state, data) {
      state.id = data.id
      state.createdAt = data.created_at
    },
    SET_WS(state, data) {
      state.ws = data.ws
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
    }
  }
}
