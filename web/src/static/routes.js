import Vue from 'vue'
import VueResource from 'vue-resource'

Vue.use(VueResource)

const wsSchema = window.location.protocol === 'http:' ? 'ws:' : 'wss:'
const baseSocket = window.location.hostname === '127.0.0.1' ? `${wsSchema}//127.0.0.1:8000` : `${wsSchema}//${window.location.hostname}`

export default {
  startStream: () => (
    Vue.resource(`/api/start-stream`)
  ),
  socketPath: id => (
    `${baseSocket}/stream/${id}`
  )
}
