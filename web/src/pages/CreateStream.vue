<template>
  <div>
    <p>CreateStream</p>
    <div>
      <button type="button" @click="startStream()">
        Create Stream
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
      }
    }
  }
</script>

<style scoped>

</style>
