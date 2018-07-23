import Routes from '~/static/routes'
import router from '~/router'

export default {
  state: {
    id: null,
    createdAt: null,
    ws: null,
    stream: null,
    mediaRecorder: null
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
    BROADCAST_STREAM_DATA(state, data) {
      state.ws.send(data)
    },
    SET_STREAM(state, data) {
      state.stream = data
    },
    SET_MEDIA_RECORDER(state, data) {
      state.mediaRecorder = data
    }
  },

  actions: {
    connectTv({commit}) {
      router.push({path: `/tv`})
    },
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
    setStreamMedia({commit}, stream) {
      const mediaRecorder = new MediaRecorder(stream);
      mediaRecorder.addEventListener('dataavailable', function (event) {
        console.log('dataavailable')
        commit('BROADCAST_STREAM_DATA', event.data)
      })
      mediaRecorder.addEventListener('start', function (event) {
        console.log('start')
      })
      mediaRecorder.addEventListener('stop', function (event) {
        console.log('stop')
      })
      mediaRecorder.addEventListener('pause', function (event) {
        console.log('pause')
      })
      mediaRecorder.addEventListener('resume', function (event) {
        console.log('resume')
      })
      mediaRecorder.addEventListener('error', function (event) {
        console.log('error')
      })
      mediaRecorder.start(500);



      commit('SET_STREAM', stream)
      commit('SET_MEDIA_RECORDER', mediaRecorder)
    }
  }
}
