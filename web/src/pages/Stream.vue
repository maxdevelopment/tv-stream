<template>
  <div>
    Stream
  </div>
</template>

<script>
  import Routes from '~/static/routes'
  import router from '~/router'

  export default {
    name: "Stream",
    beforeMount() {
      if (!this.stream.id) {
        router.push({path: `/`})
        return
      }
      this.$store.dispatch('setWs', {ws: new WebSocket(Routes.socketPath(this.stream.id))})
    },
    mounted() {
      console.log(this.stream.ws)
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
