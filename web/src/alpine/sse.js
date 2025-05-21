export default (url) => ({
  data: {},
  source: {},
  updatedAt: '',

  init() {
    this.source = new EventSource(url)
    this.source.onmessage = event => {
      this.updatedAt = new Date().toString().split('(')[0].trim()
      this.data = JSON.parse(event.data) ?? {}
      if (this.data.errorCode) {
        console.log(this.data)
      }
    }
  },

  destroy() {
    this.source.close()
  },
})

