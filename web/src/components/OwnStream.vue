<template>
  <div>
    own stream
    <video id="ownPic" ref="ownStream" autoplay playsinline></video>
  </div>
</template>

<script>
  import {mapGetters} from 'vuex'

  export default {
    name: "OwnStream",
    watch: {
      myStream(val) {
        this.startStream(val)
      }
    },
    methods: {
      startStream(stream) {
        const video = this.$refs.ownStream
        video.muted = true
        video.srcObject = stream
        setTimeout(() => video.play().catch(this.logError), 1000)
      },

      logError(err) {
        console.error(`${err.name}: ${err.message}`)
      }
    },
    computed: {
      ...mapGetters({
        myStream: 'getStream'
      })
    }
  }
</script>

<style scoped>

</style>
