<template>
  <div>
    Stream
    <own-stream></own-stream>
  </div>
</template>

<script>
  import OwnStream from '~/components/OwnStream'
  import Routes from '~/static/routes'
  import router from '~/router'

  export default {
    name: "Stream",
    components: {
      OwnStream
    },
    beforeMount() {
      if (!this.stream.id) {
        router.push({path: `/`})
        return
      }
      this.$store.dispatch('setWs', {ws: new WebSocket(Routes.socketPath(this.stream.id))})
    },
    mounted() {
      console.log(this.stream.ws)
      this.startStream()
    },
    methods: {
      startStream() {

        if (navigator.mediaDevices && navigator.mediaDevices.getUserMedia) {
          navigator.mediaDevices.getUserMedia({video: true, audio: true})
            .then(stream => {


              this.$store.dispatch('setStreamMedia', stream)

            }).catch(this.logError);
        }
      }
    },
    computed: {
      stream() {
        return this.$store.getters.stream
      }
    }
  }
</script>

<style scoped>

</style>
