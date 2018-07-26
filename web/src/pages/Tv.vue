<template>
  <div>
    tv
    <video ref="videoRef" src="" autoplay playsinline></video>
  </div>
</template>

<script>
  import Routes from '~/static/routes'

  export default {
    name: "Tv",
    data() {
      return {
        ws: null
      }
    },
    beforeMount() {
      this.ws = new WebSocket(Routes.socketPath('tv'))
    },
    mounted() {
      this.ws.addEventListener('open', ()=> {
        console.log('TV ws connected')
      })
      this.ws.addEventListener('close', ()=> {
        console.log('TV ws disconnected')
      })
      this.ws.addEventListener('error', ()=> {
        console.log('TV ws ERR')
      })
      this.ws.addEventListener('message', (message)=> {
        console.log('TV ws MSG')
        console.log(message)

        let blob = message.data
        console.log('size:', blob.size)

        let newBlob = blob.slice(0, blob.size, 'video/webm')

        const video = this.$refs.videoRef



        let reader = new FileReader()
        reader.addEventListener('loadend', ()=> {
          // reader.result contains the contents of blob as a typed array
          console.log('loadend')
          console.log(reader.result)

        })
        reader.readAsArrayBuffer(newBlob)


        // const video = this.$refs.videoRef
        // video.src= objUrl
        // setTimeout(() => video.play().catch(this.logError), 1000)




        // let videoUrl = window.URL.createObjectURL(message.data.data)
        //
        // video.muted = true
        // video.srcObject = videoUrl
        // setTimeout(() => video.play().catch(this.logError), 1000)


      })
    }
  }
</script>

<style scoped>

</style>
