<template>
  <div>
    <p>CreateStream</p>
    <div>
      <button type="button" @click="startStream()">
        Create Stream
      </button>
      <button type="button" @click="connectTv()">
        Connect TV
      </button>
    </div>
  </div>
</template>

<script>
  export default {
    name: "CreateStream",
    methods: {
      startStream() {
        this.$store.dispatch('startStream')
      },
      connectTv() {
        this.$store.dispatch('connectTv')
      }
    }
  }
</script>

<style scoped>

</style>
